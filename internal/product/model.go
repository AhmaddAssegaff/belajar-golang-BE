package product

// Product response model
type Product struct {
	ID          string
	Name        string
	Description string
	Price       int64
	Stock       int
	ImageURL    string
}
