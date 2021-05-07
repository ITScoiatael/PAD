package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, name string, imageURL string) (*model.Category, error) {
	category := &model.Category{
		ID:       uuid.NewString(),
		Name:     name,
		ImageURL: imageURL,
	}
	r.categories = append(r.categories, category)
	return category, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*model.Category, error) {
	for i := len(r.categories) - 1; i >= 0; i-- {
		if r.categories[i].ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return r.categories[i], nil
		}
	}
	return nil, nil
}

func (r *mutationResolver) AddProduct(ctx context.Context, categoryID string, name string, description string, imageURL string) (*model.Product, error) {
	var category *model.Category
	for i := range r.categories {
		if r.categories[i].ID == categoryID {
			category = r.categories[i]
			break
		}
	}

	if category == nil {
		return nil, nil
	}

	product := &model.Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
	}

	category.Products = append(category.Products, product)
	return product, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*model.Product, error) {
	for i := range r.categories {
		for j := len(r.categories[i].Products) - 1; i >= 0; i-- {
			if r.categories[i].Products[j].ID == id {
				r.categories[i].Products = append(r.categories[i].Products[:j], r.categories[i].Products[j+1:]...)
				return r.categories[i].Products[j], nil
			}
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateSubProduct(ctx context.Context, productID string, price float64, size string, color string, amount int) (*model.SubProduct, error) {
	var product *model.Product

	for i := range r.categories {
		for j := range r.categories[i].Products {
			if r.categories[i].Products[j].ID == productID {
				product = r.categories[i].Products[j]
				break
			}
		}
	}

	if product == nil {
		return nil, nil
	}

	subProduct := &model.SubProduct{
		ID:     uuid.NewString(),
		Price:  price,
		Size:   size,
		Color:  color,
		Amount: amount,
	}

	product.SubProducts = append(product.SubProducts, subProduct)
	return subProduct, nil
}

func (r *mutationResolver) DeleteSubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	for i := range r.categories {
		for j := range r.categories[i].Products {
			for g := len(r.categories[i].Products[j].SubProducts) - 1; i >= 0; i-- {
				if r.categories[i].Products[j].SubProducts[g].ID == id {
					r.categories[i].Products[j].SubProducts = append(r.categories[i].Products[j].SubProducts[:g], r.categories[i].Products[j].SubProducts[g+1:]...)
					return r.categories[i].Products[j].SubProducts[g], nil
				}
			}
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, name string, email string, phone string, address string, region string, ccNumber string) (*model.Customer, error) {
	customer := &model.Customer{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		Address:  address,
		Region:   region,
		CcNumber: ccNumber,
	}
	r.customers = append(r.customers, customer)
	return customer, nil
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (*model.Customer, error) {
	for i := len(r.customers) - 1; i >= 0; i-- {
		if r.customers[i].ID == id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
			return r.customers[i], nil
		}
	}
	return nil, nil
}

func (r *mutationResolver) AddOrder(ctx context.Context, customerID string, amount int, createdAt string) (*model.Order, error) {
	var customer *model.Customer
	for i := range r.customers {
		if r.customers[i].ID == customerID {
			customer = r.customers[i]
			break
		}
	}

	if customer == nil {
		return nil, nil
	}

	order := &model.Order{
		ID:        uuid.NewString(),
		Amount:    amount,
		CreatedAt: createdAt,
	}

	customer.Orders = append(customer.Orders, order)
	return order, nil
}

func (r *mutationResolver) RemoveOrder(ctx context.Context, id string) (*model.Order, error) {
	for i := range r.customers {
		for j := len(r.customers[i].Orders) - 1; i >= 0; i-- {
			if r.customers[i].Orders[j].ID == id {
				r.customers[i].Orders = append(r.customers[i].Orders[:j], r.customers[i].Orders[j+1:]...)
				return r.customers[i].Orders[j], nil
			}
		}
	}
	return nil, nil
}

func (r *mutationResolver) AddSubProduct(ctx context.Context, orderID string, price float64, size string, color string, amount int) (*model.SubProduct, error) {
	var order *model.Order

	for i := range r.customers {
		for j := range r.customers[i].Orders {
			if r.customers[i].Orders[j].ID == orderID {
				order = r.customers[i].Orders[j]
				break
			}
		}
	}

	if order == nil {
		return nil, nil
	}

	subProduct := &model.SubProduct{
		ID:     uuid.NewString(),
		Price:  price,
		Size:   size,
		Color:  color,
		Amount: amount,
	}

	order.Products = append(order.Products, subProduct)
	return subProduct, nil
}

func (r *mutationResolver) RemoveSubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	for i := range r.customers {
		for j := range r.customers[i].Orders {
			for g := len(r.customers[i].Orders[j].Products) - 1; i >= 0; i-- {
				if r.customers[i].Orders[j].Products[g].ID == id {
					r.customers[i].Orders[j].Products = append(r.customers[i].Orders[j].Products[:g], r.customers[i].Orders[j].Products[g+1:]...)
					return r.customers[i].Orders[j].Products[g], nil
				}
			}
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, login string, password string) (*model.Admin, error) {
	admin := &model.Admin{
		ID:       uuid.NewString(),
		Login:    login,
		Password: password,
	}
	r.admins = append(r.admins, admin)
	return admin, nil
}

func (r *mutationResolver) DeleteAdmin(ctx context.Context, id string) (*model.Admin, error) {
	for i := len(r.admins) - 1; i >= 0; i-- {
		if r.admins[i].ID == id {
			r.admins = append(r.admins[:i], r.admins[i+1:]...)
			return r.admins[i], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.categories, nil
}

func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			return r.categories[i], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return r.customers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	for i := range r.customers {
		if r.customers[i].ID == id {
			return r.customers[i], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Products(ctx context.Context, categoryID string) ([]*model.Product, error) {
	for i := range r.categories {
		if r.categories[i].ID == categoryID {
			return r.categories[i].Products, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	for i := range r.categories {
		for j := range r.categories[i].Products {
			if r.categories[i].Products[j].ID == id {
				return r.categories[i].Products[j], nil
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) SubProducts(ctx context.Context, productID string) ([]*model.SubProduct, error) {
	for i := range r.categories {
		for j := range r.categories[i].Products {
			if r.categories[i].Products[j].ID == productID {
				return r.categories[i].Products[j].SubProducts, nil
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) SubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	for i := range r.categories {
		for j := range r.categories[i].Products {
			for g := range r.categories[i].Products[j].SubProducts {
				if r.categories[i].Products[j].SubProducts[g].ID == id {
					return r.categories[i].Products[j].SubProducts[g], nil
				}
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) Orders(ctx context.Context, customerID string) ([]*model.Order, error) {
	for i := range r.customers {
		if r.customers[i].ID == customerID {
			return r.customers[i].Orders, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	for i := range r.customers {
		for j := range r.customers[i].Orders {
			if r.customers[i].Orders[j].ID == id {
				return r.customers[i].Orders[j], nil
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) OrderedSubProducts(ctx context.Context, orderID string) ([]*model.SubProduct, error) {
	for i := range r.customers {
		for j := range r.customers[i].Orders {
			if r.customers[i].Orders[j].ID == orderID {
				return r.customers[i].Orders[j].Products, nil
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) OrderedSubProduct(ctx context.Context, ID string) (*model.SubProduct, error) {
	for i := range r.customers {
		for j := range r.customers[i].Orders {
			for g := range r.customers[i].Orders[j].Products {
				if r.customers[i].Orders[j].Products[g].ID == ID {
					return r.customers[i].Orders[j].Products[g], nil
				}
			}
		}
	}
	return nil, nil
}

func (r *queryResolver) Admins(ctx context.Context) ([]*model.Admin, error) {
	return r.admins, nil
}

func (r *queryResolver) Admin(ctx context.Context, id string) (*model.Admin, error) {
	for i := range r.admins {
		if r.admins[i].ID == id {
			return r.admins[i], nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
