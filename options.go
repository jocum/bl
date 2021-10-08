/*
	@Description 定义一些基础属性 和默认值

*/
package ga

type Options struct {
	/*
		@description  生成的染色体数量
		初始化族群的数量
	*/
	ChromosomeNum int
	/*
		@description  染色体复制比例
			每次迭代遗传都会保留复制上一代优秀的染色体。
			根据这个比例保留染色体
	*/
	CpRate float64 //	染色体复制比例
	/*
		@description  族群繁衍多少代
			遗传算法是个求最优解的过程是没有终点的 所以这里设置参数控制次数
	*/
	IteratorNum int // 遗传多少代 即迭代次数

	/*
		@descirption 基因交叉率
	*/
	CrossRate float64

	/*
		@description 变异率
		每次遗传都有一定概率发生变异,该参数决定变异概率
	*/
	MutationRate float64

	/*
		@description 变异基因比例
	*/
	MutationGeneRate float64
	/*
		@description  适应度目标值
			遗传算法的结束条件之一 不推荐使用
	*/
	// AdaptabilityTarget float64 // 适应度目标
}

var (
	_defaultChromosomeNum            = 20
	_defaultCpRate           float64 = 0.2
	_defaultIteratorNum              = 50
	_defaultCrossRate        float64 = 0.2
	_defaultMutationRate     float64 = 0.1
	_defaultMutationGeneRate float64 = 0.1
)

// 定义一个参数赋值规范
// type Option interface {
// apply(*Options)
// }

// 用函数实现一个参数赋值 这个函数调用自身
type Option func(*Options)

/*
	@description  设置族群数量
*/
func ChromosomeNum(num int) Option {
	return func(os *Options) {
		os.ChromosomeNum = num
	}
}

/*
	@description 设置 复制比例
*/
func CpRate(rate float64) Option {
	return func(os *Options) {
		os.CpRate = rate
	}
}

/*
	@descriptin 设置配置族群迭代数
*/
func IteratorNum(num int) Option {
	return func(os *Options) {
		os.IteratorNum = num
	}
}

/*
	@description 设置交叉概率
*/
func CrossRate(rate float64) Option {
	return func(os *Options) {
		os.CrossRate = rate
	}
}

/*
	@description 变异概率
*/
func MutationRate(rate float64) Option {
	return func(os *Options) {
		os.MutationRate = rate
	}
}

/*
	@description 参与变异的基因数量
*/
func MutationGeneRate(rate float64) Option {
	return func(os *Options) {
		os.MutationGeneRate = rate
	}
}

// 初始化&配置options
func newOptions(opts ...Option) Options {
	os := Options{
		IteratorNum:      _defaultIteratorNum,
		ChromosomeNum:    _defaultChromosomeNum,
		CrossRate:        _defaultCrossRate,
		CpRate:           _defaultCpRate,
		MutationRate:     _defaultMutationRate,
		MutationGeneRate: _defaultMutationGeneRate,
	}
	for _, opt := range opts {
		opt(&os)
	}
	return os
}
