package core

// 入箱的个体矩形定义
type Rect interface {
	/*
		@description 入箱矩形需要能获取到宽高 返回矩形的宽高数据
	*/
	GetW() (w int)
	GetH() (h int)
	/*
		@description 入箱矩形90度旋转
			对箱体90旋转 互换宽高
	*/
	Rotate() (w int, h int)
	/*
		@description 设置入箱坐标
			设置该矩形目前所在的坐标 即移动矩形
			ps:该点为矩形右下点，整体逻辑使用该点为顶点
	*/
	SetPoint(x, y int)
	GetPoint() *Point
	/*
		@description 获取面积
			获取该矩形的面积
	*/
	GetArea() (area int)
	/*
		@description 获取矩形上水平线
			通过矩形本身的宽高和坐标，获取矩形的上水平线
	*/
	GetUpHorizontal() HLine
	/*
		@desctiption 获取矩形的下水平线
	*/
	GetDownHorizontal() HLine
	/*
		@description 获取矩形的左垂直线
	*/
	GetLeftVertical() VLine
	/*
		@description 获取矩形的右垂直线
	*/
	GetRightVertical() VLine
}
