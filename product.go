package paystack

import "fmt"

type ProductService service

type Product struct {
	Name              string      `json:"name,omitempty"`
	Description       string      `json:"description,omitempty"`
	Currency          string      `json:"currency,omitempty"`
	Price             int         `json:"price,omitempty"`
	Quantity          int         `json:"quantity,omitempty"`
	IsShippable       bool        `json:"is_shippable,omitempty"`
	Unlimited         bool        `json:"unlimited,omitempty"`
	Integration       int         `json:"integration,omitempty"`
	Domain            string      `json:"domain,omitempty"`
	Metadata          interface{} `json:"metadata,omitempty"`
	Slug              string      `json:"slug,omitempty"`
	ProductCode       string      `json:"product_code,omitempty"`
	QuantitySold      int         `json:"quantity_sold,omitempty"`
	Type              string      `json:"type,omitempty"`
	ShippingFields    interface{} `json:"shipping_fields,omitempty"`
	Active            bool        `json:"active,omitempty"`
	InStock           bool        `json:"in_stock,omitempty"`
	Minimum_orderable int         `json:"minimum_orderable,omitempty"`
	MaximumOrderable  int         `json:"maximum_orderable,omitempty"`
	LowStockAlert     bool        `json:"low_stock_alert,omitempty"`
	Id                int         `json:"id,omitempty"`
	CreatedAt         string      `json:"createdAt,omitempty"`
	UpdatedAt         string      `json:"updatedAt,omitempty"`
}

type ProductRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       string `json:"price,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Unlimited   bool   `json:"unlimited,omitempty"` // Optional
	Quantity    int    `json:"quantity,omitempty"`  // Optional
}

// ProductList is a list object for Products.
type ProductList struct {
	Meta   ListMeta
	Values []Product `json:"data"`
}

// Create a product on your integration
// For more details see https://paystack.com/docs/api/product/#create
func (s *ProductService) Create(request *ProductRequest) (*Product, error) {
	u := "/product"
	product := &Product{}
	err := s.client.Call("POST", u, request, product)
	return product, err
}

// List returns a list of Products.
// For more details see https://paystack.com/docs/api/product/#list
func (s *ProductService) List() (*ProductList, error) {
	return s.ListN(10, 1)
}

// ListN returns a list of Products
// For more details see https://paystack.com/docs/api/product/#list
func (s *ProductService) ListN(count, offset int) (*ProductList, error) {
	u := paginateURL("/product", count, offset)
	products := &ProductList{}
	err := s.client.Call("GET", u, nil, products)
	return products, err
}

// Get details of Product with the specified id
// For more details see https://paystack.com/docs/api/product/#fetch
func (s *ProductService) Get(id int) (*Product, error) {
	url := fmt.Sprintf("/product/%d", id)
	product := &Product{}
	err := s.client.Call("GET", url, nil, product)
	return product, err
}

// Update details of a Product on your integration
// For more details see https://paystack.com/docs/api/product/#update
func (s *ProductService) Update(id int, request *ProductRequest) (*Product, error) {
	url := fmt.Sprintf("product/%d", id)
	product := &Product{}
	err := s.client.Call("PUT", url, request, product)
	return product, err
}
