package core

// 这里实现一个默认矩形
type DefaultRect struct {
	W        int    //	宽
	H        int    //	高
	Point    *Point //	右顶点坐标
	Area     int    // 	面积	平方mm
	IsRotate bool   //	是否90度旋转
}

func NewDefaultRect(w, h int) *DefaultRect {
	return &DefaultRect{
		W:     w,
		H:     h,
		Point: &Point{0, 0},
		Area:  w * h,
	}
}

// 返回高
func (dr *DefaultRect) GetW() int {
	return dr.W
}
func (dr *DefaultRect) GetH() int {
	return dr.H
}

/*
	@description 对矩形进行旋转
						如果矩形已经旋转过了则将isRotate改回false
*/
func (dr *DefaultRect) Rotate() (int, int) {
	// 互换宽高
	dr.W, dr.H = dr.H, dr.W
	// 对是否旋转取反 旋转后的旋转就是回归原图
	dr.IsRotate = !dr.IsRotate
	return dr.W, dr.H
}
func (dr *DefaultRect) GetPoint() *Point {
	return dr.Point
}
func (dr *DefaultRect) SetPoint(x, y int) {
	dr.Point.Set(x, y)
}
func (dr *DefaultRect) GetArea() int {
	return dr.Area
}

/*
	@description 通过坐标和宽高获取上水平线
			左点是 矩阵右顶点的 坐标减对应的 宽高
			右点 就y坐标高度变化
*/
func (dr *DefaultRect) GetUpHorizontal() HLine {
	return HLine{
		Left:  Point{X: dr.Point.X - dr.W, Y: dr.Point.Y - dr.H},
		Right: Point{X: dr.Point.X, Y: dr.Point.Y - dr.H},
	}
}

/*
	@description 通过坐标和宽高获取水平下线
			水平下线y坐标不变为顶点坐标  x坐标 左点为坐标-宽度
*/
func (dr *DefaultRect) GetDownHorizontal() HLine {
	return HLine{
		Left:  Point{X: dr.Point.X - dr.W, Y: dr.Point.Y},
		Right: Point{X: dr.Point.X, Y: dr.Point.Y},
	}
}

/*
	@description 通过坐标和宽高获取垂直左线
			垂直左线 x坐标为 顶点x-宽度  y 上点为  顶点y-高度		下点不变
*/
func (dr *DefaultRect) GetLeftVertical() VLine {
	return VLine{
		Up:   Point{X: dr.Point.X - dr.W, Y: dr.Point.Y - dr.H},
		Down: Point{X: dr.Point.X - dr.W, Y: dr.Point.Y},
	}
}

/*
	@description 通过坐标和宽高获取垂直右线
			垂直右线  x 坐标 不变     y 下点不变  上点为 顶点y-高度
*/
func (dr *DefaultRect) GetRightVertical() VLine {
	return VLine{
		Up:   Point{X: dr.Point.X, Y: dr.Point.Y - dr.H},
		Down: Point{X: dr.Point.X, Y: dr.Point.Y},
	}
}
