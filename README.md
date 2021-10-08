#  遗传算法+势场算法 装箱


## 单独势场算法 装箱
  - 默认实现了一个矩形 所以可以用core.NewDefaultRect()来获取矩阵。如果需要自己特定的矩阵或者有附加功能可以在自己实现一个。
  - 具体定义在core/rect_type.go 文件中
```
  import (
    "github.com/jocum/bl/core"
  )

  func Packing() string {
    rects := make([]core.Rect, 0)
    rects = append(rects,
      core.NewDefaultRect(105, 87, 1),
      core.NewDefaultRect(87, 87, 2),
      core.NewDefaultRect(57, 105, 3),
      core.NewDefaultRect(66, 57, 4),
    )
    // 尝试排序 
    sort.Slice(rects, func(i, j int) bool {
     	return rects[i].GetArea() > rects[j].GetArea()
    })
    // 传入需要装箱的矩形和箱子的大小
    bl := core.NewBl(1650, 2400, rects)
    // 打包装箱
    bl.Packing()
    // 计算综合面积使用率
    bl.CountAdaptability()
    jsonb, err := json.Marshal(gabl)
    if err != nil {
      return err.Error()
    }
    return string(jsonb)
  }
  
```

## 遗传算法优化势场算法

```
  func Ga() string {
    // 初始化遗传算法
    gaStruct := ga.NewGeneticAlgorithm(
      // bl算法结构体
      bl,
      // 需要遗传迭代的次数
      ga.IteratorNum(50),
      // 族群数量
      ga.ChromosomeNum(20),
      // 族群每次迭代复制比例
      ga.CpRate(0.2),
      // 个体基因突变几率
      ga.MutationRate(0.2),
      // 个体基因突变后需要执行突变的基因比例
      ga.MutationGeneRate(0.1),
    )
    // 迭代
    gabl := gaStruct.Iterator()
    jsonb, err := json.Marshal(gabl)
    if err != nil {
      return err.Error()
    }
    return string(jsonb)
  }
```