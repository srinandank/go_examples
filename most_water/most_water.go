package main

import "fmt"

func most_water(pillars []int) (int, int, int, int) {
	first, last := 0, len(pillars)-1
	max_water_volume := 0
	max_first_pillar := first
	max_second_pillar := last
	pillar_gap := last - first
	for first != last {
		vol := (last - first) * min(pillars[first], pillars[last])
		if vol > max_water_volume {
			max_water_volume = vol
			max_first_pillar = first
			max_second_pillar = last
			pillar_gap = last - first
		}
		if pillars[first] <= pillars[last] {
			first++
		} else {
			last--
		}
	}
	return max_first_pillar, max_second_pillar, pillar_gap, max_water_volume
}

func main() {
	pillars := [...]int{1, 2, 8, 3, 4, 6, 5, 5, 3}
	first_pillar, second_pillar, pillar_gap, water_area := most_water(pillars[:])
	fmt.Printf("Most water is found between pillars of heights %d and %d with distance %d between them: %d\n", pillars[first_pillar], pillars[second_pillar], pillar_gap, water_area)
}
