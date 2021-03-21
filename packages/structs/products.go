package structs

type Product []map[int]string

func NewProduct(products []string) *Product{
	prod := Product{}
	for index,val := range products{
		prod = append(prod, map[int]string{index:val})
	}
	return &prod
}

