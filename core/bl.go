/*
	@description 以左上的点做为原点 	 |-------------|  大概是这个样子 所以整体是向上运动的
																	|							|
																	|							|
																	|							|
																	|							|
*/
package core

// 具体操作结构
type Bl struct {
	W      int
	H      int
	Boxs   []*Box //	箱子
	Rects  []Rect //  需要装箱的元素
	Export []Rect //  无法入箱的矩形
}

func NewBl(w, h int, rects []Rect) *Bl {
	return &Bl{
		W:      w,
		H:      h,
		Rects:  rects,
		Boxs:   make([]*Box, 0),
		Export: make([]Rect, 0),
	}
}

/*
	@descriptin 装箱
*/
func (bl *Bl) Packing() {
	// 如果 需要装箱的矩形不存在 跳过
	if len(bl.Rects) <= 0 {
		return
	}
	// 申请一个箱子
	box := NewBox(bl.W, bl.H, bl)
	// 循环矩形装箱
	for _, rect := range bl.Rects {
		box.GetInto(rect)
	}
	// 如果装入箱子的数量为空表示都无法装箱
	if len(box.Rects) <= 0 {
		return
	}
	// 完成一个装箱
	bl.Boxs = append(bl.Boxs, box)

	// 判断是否还有能装箱却未装箱完成的 递归第二个箱子
	if len(box.Next) != 0 {
		bl.Rects = box.Next
		bl.Packing()
	}
}
