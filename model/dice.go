package model

import "math/rand"

type Dice struct {
	minValue int64
	maxValue int64
	currentValue int64
}

func InitDice(maxValue int64) *Dice {
	return &Dice{
		minValue:     0,
		maxValue:     maxValue,
		currentValue: 2,
	}
}

func(dice *Dice) Roll() int64 {
	currentValue := rand.Int63n(dice.maxValue)+1
	dice.currentValue = currentValue
	return currentValue
}