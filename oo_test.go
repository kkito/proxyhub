package main

import (
	"fmt"
	"testing"
)

/**
* buy
 */

// ==== interface
type IFood interface {
	buy(value int)
	eat()
	grade(value int)
	getPrice() int
	getFeel() string
	getGrade() int
}

// ============ base struct =========
type BaseFood struct {
	payPrice int
	grading  int
	feel     string
}

func (base *BaseFood) buy(pay int) {
	base.payPrice = pay
}

func (base *BaseFood) getPrice() int {
	return base.payPrice
}

func (base *BaseFood) getGrade() int {
	return base.grading
}

func (base *BaseFood) getFeel() string {
	return base.feel
}

func (base *BaseFood) grade(value int) {
	base.grading = value
}

func (base *BaseFood) eat() string {
	return "nomal eat"
}

// logical struct Noodle
type Noodle struct {
	BaseFood
}

func (noodle *Noodle) eat() {
	noodle.feel = "eat noodle"
}

// logical struct Rice
type Rice struct {
	BaseFood
}

func (rice *Rice) eat() {
	rice.feel = "eat rice"
}

// a wrapper
type FoodWrapper struct {
	food IFood
}

func (wrapper *FoodWrapper) goEat(price int, grading int) {
	wrapper.food.buy(price)
	wrapper.food.eat()
	wrapper.food.grade(grading)
}

func (wrapper *FoodWrapper) showResult() string {
	return fmt.Sprintf("pay %d, eat %s, grade %d",
		wrapper.food.getPrice(),
		wrapper.food.getFeel(),
		wrapper.food.getGrade(),
	)
}

// ========
func TestNoodle(t *testing.T) {
	noodle := Noodle{}
	wrapper := FoodWrapper{
		food: &noodle,
	}
	wrapper.goEat(10, 1)
	t.Log(wrapper.showResult())

	rice := Rice{}
	wrapper.food = &rice
	wrapper.goEat(5, 2)
	t.Log(wrapper.showResult())
}
