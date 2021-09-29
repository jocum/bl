package core

import (
	"encoding/json"
	"testing"
)

func TestHSort(t *testing.T) {
	box := &Box{
		Horizontals: []HLine{
			{Left: Point{X: 0, Y: 30}, Right: Point{X: 30, Y: 30}},
			{Left: Point{X: 100, Y: 100}, Right: Point{X: 130, Y: 100}},
			{Left: Point{X: 0, Y: 50}, Right: Point{X: 30, Y: 50}},
			{Left: Point{X: 200, Y: 30}, Right: Point{X: 230, Y: 30}},
			{Left: Point{X: 400, Y: 90}, Right: Point{X: 430, Y: 90}},
			{Left: Point{X: 3, Y: 10}, Right: Point{X: 300, Y: 10}},
		},
	}
	t.Logf("before %+v \n", box)
	box.HSort()
	t.Logf("after %+v \n", box)
}

/*
	@descirption 模拟一个装箱过程 测试装箱是否

*/

func TestPacking(t *testing.T) {
	// 创建复数的箱子  填充
	// 	{Id: 1, Num: "20090809", Count: 2, Size: "105*87", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/001-4.jpeg"},
	// 	{Id: 2, Num: "23451212", Count: 1, Size: "105*48", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/test-search2.jpg"},
	// 	{Id: 3, Num: "23451213", Count: 1, Size: "40*39", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/test.jpg"},
	// 	{Id: 4, Num: "23451222", Count: 1, Size: "85*105", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/6b06c30f-5031-48bc-9c21-62bc061bbb6c.jpeg"},
	// 	{Id: 5, Num: "23451214", Count: 8, Size: "105*105", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/6f803d08326a04f50067980928bd9a1d.jpeg"},
	// 	{Id: 6, Num: "23451215", Count: 1, Size: "87*105", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/6b06c30f-5031-48bc-9c21-62bc061bbb6c.jpeg"},
	// 	{Id: 7, Num: "23451216", Count: 1, Size: "79*105", Url: "https://air-oss.oss-accelerate.aliyuncs.com/flex/tmp_e36f769a08320f4297ffa37a29879557.jpg"},
	// }
	d := []map[string]int{
		{"num": 2, "w": 105, "h": 87},
		{"num": 1, "w": 105, "h": 48},
		{"num": 1, "w": 40, "h": 39},
		{"num": 1, "w": 85, "h": 105},
		{"num": 8, "w": 105, "h": 105},
		{"num": 1, "w": 87, "h": 105},
		{"num": 1, "w": 79, "h": 105},
	}
	rects := make([]Rect, 0)
	for _, v := range d {
		num := v["num"]
		for i := 0; i < num; i++ {
			tmp := NewDefaultRect(v["w"], v["h"])
			rects = append(rects, tmp)
		}
	}
	bl := NewBl(1650, 2400, rects)
	bl.Packing()
	jsonb, err := json.Marshal(bl)
	if err != nil {
		t.Log("err ", err)
	}
	t.Log("json ", string(jsonb))
}
