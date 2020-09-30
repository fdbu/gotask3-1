package card

import (
	"image"
	"time"
)

type Transaction struct {
	Id     int64
	Sum    float64
	Date   time.Unix
	Mcc    string
	Status string
}
type Card struct {
	id       int64
	issuer   string
	balance  int64
	currency string
	number   string
	icon     string
}
