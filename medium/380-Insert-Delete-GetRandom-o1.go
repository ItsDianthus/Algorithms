package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	setMap map[int]int
	setLst []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{setMap: make(map[int]int), setLst: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.setMap[val]
	if !exists {
		this.setLst = append(this.setLst, val)
		this.setMap[val] = len(this.setLst) - 1
	}
	return !exists
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.setMap[val]
	if !exists {
		return false
	}
	delete(this.setMap, val)

	if len(this.setLst) > 1 && idx != len(this.setLst)-1 {
		this.setLst[idx] = this.setLst[len(this.setLst)-1]
		this.setMap[this.setLst[idx]] = idx
	}
	this.setLst = this.setLst[:len(this.setLst)-1]
	return true
}

// АХАХ ну смешно ведь реально работает! Правда перевес в сторону первого добавленного значения.
// Но в целом довольно рандомно. Только распределение не равномерное ыыы
func (this *RandomizedSet) GetRandom1() int {
	for v, _ := range this.setMap {
		return v
	}
	return 0
}

func (this *RandomizedSet) GetRandom() int {
	return this.setLst[rand.Intn(len(this.setLst))]
}

func main() {

	setA := Constructor()

	fmt.Println(setA.Insert(1))
	fmt.Println(setA.Remove(1))
	fmt.Println(setA.Insert(1))
	fmt.Println(setA.Insert(1))

}
