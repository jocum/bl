package main

import (
	"Bl/core"
	"Bl/ga"
	"encoding/json"
	"net/http"
)

func PointHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(Packing()))
}

func Packing() string {
	d := []map[string]int{
		{"num": 23, "w": 105, "h": 87},
		{"num": 13, "w": 105, "h": 48},
		{"num": 130, "w": 40, "h": 39},
		{"num": 11, "w": 85, "h": 105},
		{"num": 18, "w": 105, "h": 105},
		{"num": 10, "w": 87, "h": 105},
		{"num": 11, "w": 79, "h": 105},
		// {"num": 2, "w": 1700, "h": 1050},
	}
	rects := make([]core.Rect, 0)
	j := 1
	for _, v := range d {
		num := v["num"]
		for i := 0; i < num; i++ {
			tmp := core.NewDefaultRect(v["w"], v["h"], j)
			rects = append(rects, tmp)
			j++
		}
	}
	// 尝试排序 但好像并不能提高面积使用率
	// sort.Slice(rects, func(i, j int) bool {
	// 	return rects[i].GetArea() > rects[j].GetArea()
	// })
	// core.Shuffle(rects)
	bl := core.NewBl(1650, 2400, rects)
	// 初始化遗传算法
	gaStruct := ga.NewGeneticAlgorithm(bl,
		ga.IteratorNum(50),
		ga.ChromosomeNum(20),
		ga.CpRate(0.2),
		ga.MutationRate(0.2),
		ga.MutationGeneRate(0.01),
	)
	gabl := gaStruct.Iterator()
	// fmt.Println("gaBl", gaBl)
	// bl.Packing()
	// bl.CountAdaptability()
	jsonb, err := json.Marshal(gabl)
	if err != nil {
		return err.Error()
	}
	return string(jsonb)
}

func main() {
	http.HandleFunc("/", PointHander)
	http.ListenAndServe(":9999", nil)
}
