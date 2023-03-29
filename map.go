package main
import (
	"fmt"
	"sort"
)

func main() {
    // var map1 map[string]int
    // map2 := make(map[string]float32, 100)
	// map3 := make(map[int][]int)
	scene := make(map[int]int)
	scene[3] = 66
	scene[6] = 4
	scene[2] = 960
	delete(scene, 3)
	for k, v := range scene {
		fmt.Println(k, v)
	}	



	mapInfo := map[string]int32{
		"roy":18,
		"kitty":16,
		"hugo":21,
		"tina":35,
		"jason":23,
	}
	type peroson struct {
		Name string
		Age int32
	}
	var lstPerson []peroson
	for k, v := range mapInfo {
		lstPerson = append(lstPerson, peroson {k, v})
	}
	sort.Slice(lstPerson, func(i, j int) bool {
		return lstPerson[i].Age < lstPerson[j].Age  // 升序
	})
	fmt.Println(lstPerson)
}