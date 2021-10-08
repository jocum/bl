package core

import (
	"math/rand"
	"sort"
	"time"
)

/*
	@description 将slice 随机化
*/
func Shuffle(slice []Rect) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

/*
	@description 按面积排序
*/
func SortByArea(rects []Rect) {
	sort.Slice(rects, func(i, j int) bool {
		return rects[i].GetArea() > rects[j].GetArea()
	})
}

/*
	@description  获取一个随机整数
*/
func RandInt(len int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len)
}
