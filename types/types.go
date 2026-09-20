package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required`
	LastName  string `json:"lastName" validate:"required`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=120`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required`
}

type ProductStore interface {
	CreateProduct(Product) error
	GetProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
}

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       string    `json:"price"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateProductPayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required,url,max=255"`
	Price       string `json:"price" validate:"required,numeric"`
	Quantity    uint   `json:"quantity" validate:"required"`
}
