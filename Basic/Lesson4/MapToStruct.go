package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)
type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("call Point method")

}

func main() {
	pointsMap := map[string]int{
		"xx": 123,
		"yy": 456,
	}
	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)

	for k,v := range pointsMap {
		fmt.Println(k,v)
	}
}