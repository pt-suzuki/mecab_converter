package tax_rate

type TaxRate struct {
	Key string
	Value string
}

var EIGHT = TaxRate{"8", "8%"}
var TEN = TaxRate{ "10","10%"}

func (e TaxRate) String() string {
	if e.Value == "" {
		return ""
	}
	return e.Value
}

func List() []TaxRate {
	list := []TaxRate{}

	list = append(list, EIGHT)
	list = append(list, TEN)

	return list
}