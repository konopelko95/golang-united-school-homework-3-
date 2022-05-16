package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	map1 := map[int]string{2: "a", 0: "b", 1: "c"}

	keys := make([]int, 0, len(map1))
	for k := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	map2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}

	keys2 := make([]int, 0, len(map2))
	for k2 := range map2 {
		keys2 = append(keys2, k2)
	}

	sort.Ints(keys2)

	return
}
