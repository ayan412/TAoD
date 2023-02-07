package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) method() {
	fmt.Println(p.X, " - call Point method")

}

func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 13,
			X: 15,
		},
	} //фиг-ые скобки означают что мапа проинициализирована

	otherPointsMap := make(map[int]Point)

	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	// fmt.Println(pointsMap)
	// fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{1, 2}
	// fmt.Println(otherPointsMap)
	// fmt.Println(otherPointsMap[1])

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {1, 2},
		}
	}

	// fmt.Println(oneMorePointsMap["a"].X)
	// fmt.Println(oneMorePointsMap["a"].Y)
	// oneMorePointsMap["a"].method()

	/*
		if oneMorePointsMap == nil {
			fmt.Println("oneMorePointsMap is nill")
		} else {
			oneMorePointsMap["a"] = Point{1,2}
		}
	*/

	key := "d"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s does exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s doesn't exist in map\n", key)
		fmt.Println(value)
		fmt.Println(Point{})
	}

}
