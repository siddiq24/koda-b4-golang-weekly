package utils

import ()

import (
	"fmt"
	"math/rand"
	"time"
)

func NewOrderCode() string {
	now := time.Now()
	datePart := now.Format("060102")
	randomNum := rand.Intn(900) + 99
	return fmt.Sprintf("POS-%s-%d", datePart, randomNum)
}
