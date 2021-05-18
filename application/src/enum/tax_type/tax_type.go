package tax_type

type TaxType struct {
	Key string
	Value string
}

var INCLUDE = TaxType{"1", "税込"}
var EXCLUDE = TaxType{ "2","税抜"}

func (e TaxType) String() string {
	if e.Value == "" {
		return ""
	}
	return e.Value
}

func List() []TaxType {
	list := []TaxType{}

	list = append(list, INCLUDE)
	list = append(list, EXCLUDE)

	return list
}