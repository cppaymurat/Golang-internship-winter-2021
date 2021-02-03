package main

import (
	"fmt"
)

func main() {
	all_goods := make([][][]int32, 0)
	all_goods = append(all_goods, [][]int32{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}})
	all_goods = append(all_goods, [][]int32{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}})
	all_goods = append(all_goods, [][]int32{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}})
	all_goods = append(all_goods, [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}})
	all_goods = append(all_goods, [][]int32{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}})
	all_goods = append(all_goods, [][]int32{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}})
	all_goods = append(all_goods, [][]int32{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}})
	all_goods = append(all_goods, [][]int32{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}})
	fmt.Print(all_goods[0])
}
