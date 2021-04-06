package types

type Money int64

type Category string

type Currency string

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type PAN string

type Payment struct {
	ID int
	Amount	Money
	Category Category
	Status Status
}

type Card struct {
	ID	int
	PAN PAN
	Balance	Money
	MinBalance	Money
	Currency	Currency
	Color string 
	Name	string
	Active	bool
}