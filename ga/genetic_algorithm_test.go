package ga

import (
	"fmt"
	"testing"
)

func Cross() {
	a := []string{"c", "b", "a", "d"}
	b := []string{"b", "a", "d", "c"}
	// c := []string{"c", "b", "b", "d"}
	// d := []string{"a", "d", "b", "c"}
	tmpa := map[int]string{2: "a", 3: "d"}
	tmpb := map[int]string{2: "d", 3: "c"}
	fmt.Println(tmpa, tmpb)
	for mk, mv := range tmpb {
		for k, v := range a {
			if v == mv {
				a[mk], a[k] = a[k], a[mk]
			}
		}
	}
	fmt.Println(a, b)
}
func TestCross(t *testing.T) {
	Cross()
}
