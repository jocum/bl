package core

/*
	@Description  箱子
			矩阵装箱 主要逻辑这里完成
	@Author cwy
*/

import (
	"sort"

	"github.com/shopspring/decimal"
)

// 箱子主体
type Box struct {
	Width       int     //	箱子宽度
	Height      int     //	箱子高度
	Rects       []Rect  //	已入箱
	Next        []Rect  `json:"-"` //  实际能入箱但目前装不下
	horizontals []HLine // 	水平线合集？合集  每个物品装箱后保留顶点在这个位置方便后续判断
	verticals   []VLine // 	垂直线集合？ 每个物品装箱完成后将垂直线坐标保留在这个位置
	UseH        int     //  使用了多少高度
	UseAera     int     //  使用的面积总和
	Rate        float64 //  面积使用率
	parent      *Bl     `json:"-"` //	父级地址
}

// 初始化箱子
func NewBox(w, h int, bl *Bl) *Box {
	return &Box{
		Width:       w,
		Height:      h,
		parent:      bl,
		horizontals: make([]HLine, 0),
		verticals:   make([]VLine, 0),
		Rects:       make([]Rect, 0),
	}
}

// 水平线排序 按y坐标
func (b *Box) HSort() {
	sort.Slice(b.horizontals, func(i, j int) bool {
		return b.horizontals[i].Right.Y > b.horizontals[j].Right.Y
	})
}

// 垂直线排序 按x坐标
func (b *Box) VSort() {
	sort.Slice(b.verticals, func(i, j int) bool {
		return b.verticals[i].Down.X > b.verticals[j].Down.X
	})
}

/*
	@description 计算面积使用率
*/
func (b *Box) CountRate() {
	userArea := decimal.NewFromInt(int64(b.UseAera))
	bgArea := decimal.NewFromInt(int64(b.UseH * b.Width))
	b.Rate, _ = userArea.Div(bgArea).Float64()
}

/*
	@description 判断两条水平线是否会相交如果会返回最大竖直移动距离
							 因为我们的矩形永远是从右下角往上所以右点可以不判断
							 --__ 例如这样 其实只需要判断 动线的左点 和不动线的右点是否重叠即可
							 __-- 例如这种情况 判断左点小于静线右点 判断动线右点和静线左点即可
							其他情况皆为相交
	@params
		move 		Line 		需要移动的水平线
		static  Line    静止的水平线
	@return
		isIntersection 	是否相交
		y 			int 		可移动的最大距离 如果会相交hd为正数 如果不相交说明可以越过可能为负数
*/
func horizontalLineIntersection(move, static HLine) (y int, isIntersection bool) {
	// 默认两线会相交
	isIntersection = true
	// 返回两点的最大y坐标距离
	y = move.Left.Y - static.Right.Y
	// 判断是否会相交 如果移动线的左点x坐标大于 静止线的右点x坐标表示不会水平相交
	if move.Left.X >= static.Right.X {
		isIntersection = false
		// 如果左点 大于 静止线右点 再判断右点和 静止线左点  如果小于则不相交
	} else if move.Right.X <= static.Left.X {
		isIntersection = false
	}
	return
}

/*
	@description 判断两条垂直线是否相交，如果相交返回最大的水平移动距离
							 因为我们的矩形永远是从右下角往上所以右点可以不判断
							 竖线不好表示啊 大概和水平线差不多 判断动线的左点 是否和不动线的右点是否重叠即可
	@params
		move 	Line 		需要移动的垂直线
		static Line 	静止的垂直线
	@return
		isIntersection 	是否相交
		x 						最大的水平移动距离
*/
func verticalLineIntersection(move, static VLine) (x int, isIntersection bool) {
	// 默认相交
	isIntersection = true
	// 返回两点最大x坐标距离
	x = move.Down.X - static.Down.X
	// 判断是否会相交 如果移动线的左点y 小于静止线的右点y坐标表示会垂直相交
	if move.Up.Y >= static.Down.Y {
		isIntersection = false
	} else if move.Down.Y <= static.Up.Y {
		isIntersection = false
	}
	return
}

/*
	@description  计算矩阵向上移动的最大距离
								逻辑为水平线按y坐标排序  循环比较
								1.如果相交则返回 距离
								2.如果全部不相交 返回最大高度-箱子高度的距离
*/
func (b *Box) getMoveY(rect Rect) int {
	// 需要返回的最大可移动距离
	yy := 0
	// 是否相交
	intersect := false
	// 对水平线按y坐标从高到低排序
	b.HSort()
	// 循环水平线
	for _, horizontal := range b.horizontals {
		if horizontal.Left.Y > (rect.GetPoint().Y - rect.GetH()) {
			continue
		}
		y, isIntersect := horizontalLineIntersection(rect.GetUpHorizontal(), horizontal)
		if isIntersect {
			yy = y
			intersect = isIntersect
			break
		}
	}
	if !intersect {
		yy = rect.GetPoint().Y - rect.GetH()
	}
	if yy < 0 {
		yy = 0
	}
	return yy
}

/*
	@description  计算矩阵向左移动的最大距离
								垂直线按x坐标排序 循环比较
								1.如果相交返回 距离
								2.如果不相交 返回最大宽度-箱子宽度的距离
*/
func (b *Box) getMoveX(rect Rect) int {
	// 需要返回的 最大向左可移动距离
	xx := 0
	// 是否相交
	intersect := false
	// 对垂直线按x从大道小排序
	b.VSort()
	// 循环垂直线对比是否会相交
	for _, vertical := range b.verticals {
		// 不考虑再后面的线 只考虑前面的
		if vertical.Up.X > (rect.GetPoint().X - rect.GetW()) {
			continue
		}
		x, isIntersect := verticalLineIntersection(rect.GetLeftVertical(), vertical)
		if isIntersect {
			xx = x
			intersect = isIntersect
			break
		}
	}
	if !intersect {
		xx = rect.GetPoint().X - rect.GetW()
	}
	if xx < 0 {
		xx = 0
	}
	return xx
}

/*
	@description 移动矩形
				根据计算出的可移动距离移动矩形
*/
func move(rect Rect, x, y int) {
	rect.GetPoint().X -= x
	rect.GetPoint().Y -= y
}

/*
	@description  判断矩形是否具备进入箱体的客观事实
								就是判断矩形是否超出箱体范围导致哪怕空箱体也无法塞入矩形
								1.情况一 箱子的宽度同时小于矩形的宽高无法入箱
								2.情况二 箱子的高度同时小于矩形的宽高也无法入箱
*/
func (b *Box) Exprot(rect Rect) bool {
	if b.Width < rect.GetW() && b.Width < rect.GetH() {
		return true
	}
	if b.Height < rect.GetW() && b.Height < rect.GetH() {
		return true
	}
	return false
}

/*
	@description 矩形进入箱体初期判断在矩形不动的情况下是否还空间容纳矩形
*/
func (b *Box) Check(rect Rect) bool {
	// 先判断 宽高是否超标
	if b.Width < rect.GetW() {
		return false
	}
	if b.Height < rect.GetH() {
		return false
	}
	for _, v := range b.horizontals {
		x, isIntersect := horizontalLineIntersection(rect.GetUpHorizontal(), v)
		if x < 0 && isIntersect {
			return false
		}
	}
	for _, v := range b.verticals {
		y, isIntersect := verticalLineIntersection(rect.GetLeftVertical(), v)
		if y < 0 && isIntersect {
			return false
		}
	}
	return true
}

/*
	@description 矩形入箱不停移动直到x y 都不可移动为止
		矩形 先往上移动再往左移动 重复这个动作直到无法再移动为止
*/
func (b *Box) GetInto(rect Rect) {
	// 先判断是否能进入到箱体
	if b.Exprot(rect) {
		b.parent.Export = append(b.parent.Export, rect)
		return
	}
	// 矩阵入箱 默认坐标为箱子的右顶点
	rect.SetPoint(b.Width, b.Height)
	// 判断当前的箱子是否还能放入这个矩形
	if !b.Check(rect) {
		// 尝试旋转矩形
		rect.Rotate()
		// 依然无空间 放入下一个箱子
		if !b.Check(rect) {
			b.Next = append(b.Next, rect)
			return
		}
	}
	// 移动装箱
	for {
		// 先向上移动
		var moveX, moveY int = 0, 0
		moveY = b.getMoveY(rect)
		move(rect, 0, moveY)
		// 然后向左移动
		moveX = b.getMoveX(rect)
		move(rect, moveX, 0)
		if moveX == 0 && moveY == 0 {
			break
		}
	}
	// 记录最高的y坐标
	if b.UseH < rect.GetPoint().Y {
		b.UseH = rect.GetPoint().Y
	}
	// 入库完成 添加入已入库数据组
	b.Rects = append(b.Rects, rect)
	// 记录矩阵总面积
	b.UseAera += rect.GetArea()
	b.horizontals = append(b.horizontals, rect.GetDownHorizontal())
	b.verticals = append(b.verticals, rect.GetRightVertical())
}
