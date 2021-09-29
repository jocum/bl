package main

import (
	"Bl/core"
	"encoding/json"
	"net/http"
	"sort"
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
		{"num": 35, "w": 40, "h": 39},
		{"num": 11, "w": 85, "h": 105},
		{"num": 18, "w": 105, "h": 105},
		{"num": 10, "w": 87, "h": 105},
		{"num": 11, "w": 79, "h": 105},
	}
	rects := make([]core.Rect, 0)
	for _, v := range d {
		num := v["num"]
		for i := 0; i < num; i++ {
			tmp := core.NewDefaultRect(v["w"], v["h"])
			rects = append(rects, tmp)
		}
	}
	sort.Slice(rects, func(i, j int) bool {
		return rects[i].GetArea() > rects[j].GetArea()
	})
	bl := core.NewBl(1650, 2400, rects)
	bl.Packing()
	jsonb, err := json.Marshal(bl)
	if err != nil {
		return err.Error()
	}
	return string(jsonb)
}

func main() {
	http.HandleFunc("/", PointHander)
	http.ListenAndServe("localhost:9999", nil)
}
