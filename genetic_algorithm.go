/*
	@description 遗传算法 通过不停的杂交变异提高符合预期的结构 这里通过这个算法增加矩阵拼版的使用率
	@author	cwy
*/
package ga

import (
	"math"
	"sort"
	"sync"

	"github.com/jocum/bl/core"

	"github.com/shopspring/decimal"
)

// 定义一个 染色体族群结构 为复数的染色体
type ChromosomeGroup []*core.Bl

// 定义 遗传计算基础属性
type GeneticAlgorithm struct {
	Options Options
	// 初始个体
	AdamChromosome *core.Bl
	// 初始个体基因数量
	AdamGeneNum int
	// 染色体族群
	ChromosomeGroup ChromosomeGroup
	// 新一代的染色体族群
	NewChromosomeGroup ChromosomeGroup
	// 历代染色体记录
	History []ChromosomeGroup
	// 需要复制的染色体数量
	CopyNum int64
	// 需要通过交叉生成的染色体数量
	CrossNum int64
	// 需要交换的基因数量
	CrossGeneNum int
	// 变异的基因数量
	MutationGeneNum int
}

// 初始化 遗传对象
func NewGeneticAlgorithm(adam *core.Bl, opts ...Option) *GeneticAlgorithm {
	// 加载配置
	os := newOptions(opts...)
	// 基因长度
	geneNum := len(adam.Rects)
	// 计算需要复制的染色体个数
	copyNum := decimal.NewFromInt(int64(os.ChromosomeNum)).Mul(decimal.NewFromFloat(os.CpRate)).IntPart()
	// 计算需要变异的基因个数
	mutationGeneNum := decimal.NewFromInt(int64(geneNum)).Mul(decimal.NewFromFloat(os.MutationGeneRate)).IntPart()
	ga := &GeneticAlgorithm{
		Options:            os,
		ChromosomeGroup:    make(ChromosomeGroup, 0),
		NewChromosomeGroup: make(ChromosomeGroup, 0),
		History:            make([]ChromosomeGroup, 0),
		CopyNum:            copyNum,
		CrossNum:           int64(os.ChromosomeNum) - copyNum,
		AdamChromosome:     adam,
		AdamGeneNum:        geneNum,
		MutationGeneNum:    int(mutationGeneNum),
	}
	return ga
}

/*
	@description 初始化最初的族群
	@params
		  chromosome   Chromosom 		// 最初的染色体
*/
func (ga *GeneticAlgorithm) ChromosomeGroupInit() {
	// 默认给一个正序染色体
	ga.ChromosomeGroup = append(ga.ChromosomeGroup, ga.AdamChromosome.Clone().Sort())
	// 循环随机创建其他染色体
	for i := 0; i < ga.Options.ChromosomeNum-1; i++ {
		ga.ChromosomeGroup = append(ga.ChromosomeGroup, ga.AdamChromosome.Clone().Shuffle())
	}
	// 计算需要交叉的基因数量
	ga.CrossGeneNum = int(math.Floor(ga.Options.CrossRate*float64(ga.AdamGeneNum) + 0.5))
	// 初始化完成后做适应度计算
	ga.Adaptability()
}

/*
	@description  计算每个染色体的适应度
		计算适应度就是面积使用率 这里对每个染色体做实际的拼版操作获取适应度
*/
func (ga *GeneticAlgorithm) Adaptability() {
	var wg sync.WaitGroup
	wg.Add(len(ga.ChromosomeGroup))
	for _, v := range ga.ChromosomeGroup {
		go func(chromosome *core.Bl) {
			// 拼版前 初始化箱子
			chromosome.Boxs = make([]*core.Box, 0)
			// 实际拼版
			chromosome.Packing()
			// 计算适应度
			chromosome.CountAdaptability()
			wg.Done()
		}(v)
	}
	wg.Wait()
	// 对染色体按适应度排序
	sort.Slice(ga.ChromosomeGroup, func(i, j int) bool {
		return ga.ChromosomeGroup[i].Adaptability > ga.ChromosomeGroup[j].Adaptability
	})
}

/*
	@description 轮盘
		这里随机获取一条染色体，染色体的适应度越高被获取的概率就越大
*/
func (ga *GeneticAlgorithm) Roulette() *core.Bl {
	if len(ga.ChromosomeGroup) == 0 {
		return nil
	}
	for {
		key := core.RandInt(len(ga.ChromosomeGroup))
		chromosome := ga.ChromosomeGroup[key]
		// 获取一个0-99的随机数  染色体的适应度为百分比 这里用0-100替代  如果适应度大于 随机数表示
		randNum := core.RandInt(100)
		if chromosome.Adaptability > randNum {
			return chromosome.Clone()
		}
	}
}

/*
	@description 截取需要交换的基因
*/
func (ga *GeneticAlgorithm) CutGene(genes core.Rects) map[int]core.Rect {
	genePart := make(map[int]core.Rect)
	// 基因交换就是切片的位置替换由于每个基因都是独一无二的无法直接的交换 只对其位置做操作
	// 这里先获取切片需要获取数量之前的随机数，通过该随机数为起点窃取Ga.CrossGeneNum个基因
	randNum := core.RandInt(ga.AdamGeneNum - ga.CrossGeneNum)
	for i := randNum; i < randNum+ga.CrossGeneNum; i++ {
		genePart[i] = genes[i]
	}
	return genePart
}

/*
	@description 交叉基因
*/
func crossGene(genes core.Rects, genePart map[int]core.Rect) {
	// 循环基因片段
	for pk, pv := range genePart {
		// 循环整个基因
		for k, v := range genes {
			// 匹配到相同的基因
			if v.GetId() == pv.GetId() {
				// 交换位置
				genes[pk], genes[k] = genes[k], genes[pk]
			}
		}
	}
}

/*
	@description 交叉染色体获取新个体
		势场法拼版的最优解 实际上就是矩形的入场顺序 所以这里的基因交叉只交换顺序
*/
func (ga *GeneticAlgorithm) Cross() {
	// 循环需要交叉生成的染色体数量
	var wg sync.WaitGroup
	wg.Add(int(ga.CrossNum))
	for i := 0; i < int(ga.CrossNum); i++ {
		go func() {
			// 轮盘随机获取父 母 染色体
			father := ga.Roulette()
			mother := ga.Roulette()
			// 截取父母基因片段
			// fatherGenePart := ga.CutGene(father.Rects)
			motherGenePart := ga.CutGene(mother.Rects)
			// 交换基因
			crossGene(father.Rects, motherGenePart)
			// crossGene(mother.Rects, fatherGenePart)
			// 在交叉过程中 概率发生变异
			randNum := core.RandInt(100)
			if int(ga.Options.MutationRate*100) > randNum {
				ga.Mutation(father)
				// ga.Mutation(mother)
			}
			// 加入新族群
			ga.NewChromosomeGroup = append(ga.NewChromosomeGroup, father)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
	@description 变异 对每个独立基因进行变异处理
		在势场拼版的方式中基因的变异为对个体基因的旋转
*/
func (ga *GeneticAlgorithm) Mutation(chromosome *core.Bl) {
	// 根据基因变异比例 循环变异基因
	for i := 0; i < ga.MutationGeneNum; i++ {
		// 获取随机数
		randNum := core.RandInt(ga.AdamGeneNum)
		// 对当前索引的基因进行变异
		chromosome.Rects[randNum].Rotate()
	}
}

/*
	@description 复制  按设置的比例复制一定比例的染色体
		这里复制适应度前几的染色体，所以复制之前要求已经计算适应度并排序
*/
func (ga *GeneticAlgorithm) Copy() {
	var wg sync.WaitGroup
	wg.Add(int(ga.CopyNum))
	for i := 0; i < int(ga.CopyNum); i++ {
		go func(index int) {
			// 复制的染色体
			chromosome := ga.ChromosomeGroup[index]
			// 加入新的染色体族群
			ga.NewChromosomeGroup = append(ga.NewChromosomeGroup, chromosome.Clone())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
	@description  迭代
		根据迭代次数、完成条件迭代
*/
func (ga *GeneticAlgorithm) Iterator() *core.Bl {
	// 先初始化获取生成族群
	ga.ChromosomeGroupInit()
	// 迭代
	for i := 0; i < ga.Options.IteratorNum; i++ {
		// 先清空 新族群
		ga.NewChromosomeGroup = make(ChromosomeGroup, 0)
		// 复制
		ga.Copy()
		// 交叉
		ga.Cross()
		// 用新的族群替换旧族群
		ga.ChromosomeGroup = ga.NewChromosomeGroup
		// 计算适应度
		ga.Adaptability()
	}

	// 迭代完成后 返回适应度最高的 个体
	return ga.ChromosomeGroup[0]
}
