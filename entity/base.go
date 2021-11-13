package entity

import (
	"fmt"
	"time"
)

// type Car struct {
// 	// ตัวแปรนี้ สามารถเรียกได้จากภายนอก
// 	Name string
// 	// ตัวแปรนี้ต้องเรียนจากภายใน
// 	color string
// }

// func (c *Car) getName() string {
// 	mazda := entity.Car{}

// }

type MinerBase struct {
	Name          string
	Value         int
	Miners        map[string]int
	UpdateChannel chan string // ใช้งาน channel
}

// func (m *Minebase) GetName() string {
// 	golds := entity.Minebase()
// }

func NewMinerBase(name string, value int) MinerBase {
	return MinerBase{
		Name:          name,
		Value:         value,
		Miners:        make(map[string]int),
		UpdateChannel: make(chan string), // initial
	}
}

// setter
// func (m *MinerBase) SetName(name string) {
// 	m.Name = name
// }

// getter
func (m *MinerBase) GetName() string {
	return m.Name
}
func (m *MinerBase) GetValue() int {
	return m.Value
}

func (m *MinerBase) GetMiner() map[string]int {
	return m.Miners
}

func (m *MinerBase) Spawn(carname string) {
	m.Miners[carname] = 0
	// sentmining
	go m.SentMining(carname)
}

func (m *MinerBase) SentMining(carname string) {
	for {
		// m.Miners[carname] = m.Miners[carname] + 100
		// fmt.Println(carname, "recive gold:", m.Miners[carname])
		// fmt.Printf("%s: %d \n", carname, m.Miners[carname])
		time.Sleep(10 * time.Second)
		m.UpdateChannel <- carname
	}
}

func (m *MinerBase) ListeningMining() {
	for {
		select {
		case carname := <-m.UpdateChannel:
			m.Miners[carname] = m.Miners[carname] + 100
			fmt.Println(carname, "recive gold:", m.Miners[carname])
		default:
			fmt.Println("no one update")
			time.Sleep(1 * time.Second)
		}
	}
}
