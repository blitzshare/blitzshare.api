package random

import (
	"math/rand"
)

type RndImpl struct {
}

type Rnd interface {
	GenerateRandomWordSequence() *string
}

func NewRnd() Rnd {
	return &RndImpl{}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (*RndImpl) GenerateRandomWordSequence() *string {
	b := make([]rune, 50)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	str := string(b)
	return &str
}
