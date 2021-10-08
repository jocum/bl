/*
	@Description 以左上的点做为原点 	 |-------------|  大概是这个样子 所以整体是向上运动的
																	|							|
																	|							|
																	|							|
																	|							|
	@Author	cwy
*/
package core

import (
	"github.com/shopspring/decimal"
)

// 具体操作结构
type Bl struct {
	W            int
	H            int
	Boxs         []*Box //	箱子
	Rects        Rects  //  装箱的元素
	UseRects     Rects  //  具体实施装箱元素集合，会根据装箱过程中next的变化而变化
	Export       Rects  //  无法入箱的矩形
	Adaptability int    // 适应性 百分比这里用0-100的数字替换便于计算
}

func NewBl(w, h int, rects Rects) *Bl {
	return &Bl{
		W:        w,
		H:        h,
		Rects:    rects,
		UseRects: rects,
		Boxs:     make([]*Box, 0),
		Export:   make(Rects, 0),
	}
}

/*
	@description 提供一个自己的copy
*/
func (bl *Bl) Clone() *Bl {
	newB := &Bl{
		W:        bl.W,
		H:        bl.H,
		Boxs:     make([]*Box, 0),
		Export:   make(Rects, 0),
		Rects:    make(Rects, 0),
		UseRects: make(Rects, 0),
	}
	for _, v := range bl.Rects {
		newB.Rects = append(newB.Rects, NewDefaultRect(v.GetW(), v.GetH(), v.GetId()))
	}
	newB.UseRects = newB.Rects
	return newB
}

/*
	@description 随机
*/
func (bl *Bl) Shuffle() *Bl {
	Shuffle(bl.Rects)
	return bl
}

/*
	@description 排序按面积
*/
func (bl *Bl) Sort() *Bl {
	SortByArea(bl.Rects)
	return bl
}

/*
	@descriptin 装箱
*/
func (bl *Bl) Packing() {
	// 如果 需要装箱的矩形不存在 跳过
	if len(bl.UseRects) <= 0 {
		return
	}
	// 申请一个箱子
	box := NewBox(bl.W, bl.H, bl)
	// 循环矩形装箱
	for _, rect := range bl.UseRects {
		box.GetInto(rect)
	}
	// 如果装入箱子的数量为空表示都无法装箱
	if len(box.Rects) <= 0 {
		return
	}
	// 计算使用率
	box.CountRate()
	// 完成一个装箱
	bl.Boxs = append(bl.Boxs, box)

	// 判断是否还有能装箱却未装箱完成的 递归第二个箱子
	if len(box.Next) != 0 {
		bl.UseRects = box.Next
		bl.Packing()
	}
}

/*
	@description 计算整体的使用率
*/
func (bl *Bl) CountAdaptability() {
	total := 0
	h := 0
	for _, box := range bl.Boxs {
		total += box.UseAera
		h += box.UseH
	}
	userArea := decimal.NewFromInt(int64(total))
	bgArea := decimal.NewFromInt(int64(h * bl.W))
	bl.Adaptability = int(userArea.Div(bgArea).Mul(decimal.NewFromInt(1000000)).IntPart())
}
