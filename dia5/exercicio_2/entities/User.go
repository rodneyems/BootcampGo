package entities

type product struct {
	Name string
	Price float32
	Quantity int
}

type User struct{
	Name string
	LastName string
	Email string
	Products []product
}

func NewProduct (name string, price float32) product {
	p := product{
		name,
		price,
		0,
	}
	return p
}
func AddProduct(user *User, product *product, quantity int) {
	product.Quantity = quantity
	user.Products = append(user.Products, *product)
}
func RemoveProducts(user *User){
	user.Products = []product{}
}
