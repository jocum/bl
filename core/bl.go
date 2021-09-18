/*
	@Description 尝试bl算法完成装箱
	@Author	cwy
*/

package core

// 点
type Point struct {
	X int
	Y int
}

// 线
type Line struct {
	Left  Point
	Right Point
}

// 箱子主体
type Box struct {
	vertexes []Point // 顶点合集  每个物品装箱后保留顶点在这个位置方便后续判断
}

func NewBox() *Box {
	return &Box{
		vertexes: make([]Point, 0),
	}
}

/*
	@description 判断两条水平线是否会相交如果会返回最大竖直移动距离
							 因为我们的矩形永远是从右下角往上所以右点可以不判断
							 --__ 例如这样 其实只需要判断 动线的左点 和不动线的右点是否重叠即可
	@params
		move 		Line 		需要移动的水平线
		static  Line    静止的水平线
	@return
		isIntersection 	是否相交
		y 			int 		可移动的最大距离 如果会相交hd为正数 如果不相交说明可以越过可能为负数
*/
func HorizontalLineIntersection(left, right Point) (y int, isIntersection bool) {
	// 返回两点的最大y坐标距离
	y = left.Y - right.Y
	// 判断是否会相交 如果移动线的左点x坐标小于 静止线的右点x坐标表示会水平相交
	if left.X < right.X {
		isIntersection = true
	}
	return
}

/*
	@description 判断两条垂直线是否相交，如果相交返回最大的水平移动距离
							 因为我们的矩形永远是从右下角往上所以右点可以不判断
							 --__ 判断动线的左点 是否和不动线的右点是否重叠即可
	@params
		move 	Line 		需要移动的垂直线
		static Line 	静止的垂直线
	@return
		isIntersection 	是否相交
		x 						最大的水平移动距离
*/
func VerticalLineIntersection(left, right Point) (x int, isIntersection bool) {
	// 返回两点最大x坐标距离
	x = left.X - right.X
	// 判断是否会相交 如果移动线的左点y 小于静止线的右点y坐标表示会垂直相交
	if left.Y < right.Y {
		isIntersection = true
	}
	return
}

/*
	@description  计算矩阵向上移动的最大距离
*/
