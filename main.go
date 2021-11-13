package main

import (
	"fmt"
	// module/folder
	"workshop/entity"
)

func main() {
	// ประกาศตัวแปร 2 แบบ
	// golang --> strong type

	// var name string

	// ประกาศแบบ assign ค่า
	// name2 := ""
	// number := 0

	// fmt.Println(name2, number)
	// // log มีผลต่อ concurrent
	// log.Println(number)

	// array
	// number2 := make([]int, 5)
	// number2 = append(number2, 100)
	// fmt.Println(number2)
	// slide
	// number3 := []int{}

	// mapping
	// mapping := map[string]int{}
	// mapInMap := map[string]int[string]int{}

	// mapping["1"] = 0
	// mineBases := entity.MinerBase{}
	// fmt.Println(mineBases.GetName())
	// fmt.Println(mineBases.GetValue())

	mineBase := entity.NewMinerBase("ฐานแร่", 0)
	fmt.Println(mineBase.GetName())
	fmt.Println(mineBase.GetValue())

	mineBase.Spawn("ไทรถแห่-1")
	mineBase.Spawn("ไทยรถแห่-2")
	fmt.Println(mineBase.GetMiner())

	// mineBase.SentMining("car1")

	mineBase.ListeningMining()
}

// สร้าง method เพื่อสร้าง รถแร่
// map key string & value int
