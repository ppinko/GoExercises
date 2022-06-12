package main

type BinarySearch struct {
	first int
	last  int
}

func SearchInts(list []int, key int) int {
	if len(list) == 0 || key < list[0] || key > list[len(list)-1] {
		return -1
	}

	bs := BinarySearch{0, len(list) - 1}
	for bs.last-bs.first > 1 {
		midVal := (bs.first + bs.last) / 2
		if list[midVal] > key {
			bs.last = midVal
		} else {
			bs.first = midVal
		}
	}
	if list[bs.first] == key {
		return bs.first
	} else if list[bs.last] == key {
		return bs.last
	} else {
		return -1
	}
}

func main() {
	SearchInts([]int{1, 3, 4, 6, 8, 9, 11}, 6)
}
