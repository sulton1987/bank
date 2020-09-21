package types

type Money int64
type Cathegory string

type Payment struct {
	ID        int
	Amount    Money
	Cathegory Cathegory
}
