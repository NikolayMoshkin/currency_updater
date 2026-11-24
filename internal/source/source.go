package source

type Source interface {
	ParseCurrency() *Message
}

const (
	FRANKFURTER string = "frankfurter"
)
