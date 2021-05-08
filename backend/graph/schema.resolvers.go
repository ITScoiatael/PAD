package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"
	"graphql-server/prisma-client/db"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, name string, imageURL string) (*model.Category, error) {
	category, err := r.Prisma.Category.CreateOne(
		db.Category.ID.Set(uuid.NewString()),
		db.Category.Name.Set(name),
		db.Category.ImageURL.Set(imageURL),
	).Exec(ctx)
	return &model.Category{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}, err
}

func (r *mutationResolver) EditCategory(ctx context.Context, id string, name *string, imageURL *string) (*model.Category, error) {
	category, err := r.Prisma.Category.FindUnique(
		db.Category.ID.Equals(id),
	).Update(
		db.Category.Name.Set(*name),
		db.Category.ImageURL.Set(*imageURL),
	).Exec(ctx)
	return &model.Category{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}, err
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*model.Category, error) {
	category, err := r.Prisma.Category.FindUnique(
		db.Category.ID.Equals(id),
	).Delete().Exec(ctx)

	return &model.Category{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}, err
}

func (r *mutationResolver) AddProduct(ctx context.Context, categoryID string, name string, description string, imageURL string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditProduct(ctx context.Context, categoryID string, name *string, description *string, imageURL *string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSubProduct(ctx context.Context, productID string, price float64, size string, color string, amount int) (*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditSubProduct(ctx context.Context, productID string, price *float64, size *string, color *string, amount *int) (*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, name string, email string, phone string, address string, region string, ccNumber string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditCustomer(ctx context.Context, id string, name *string, email *string, phone *string, address *string, region *string, ccNumber *string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddOrder(ctx context.Context, customerID string, amount int, createdAt string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditOrder(ctx context.Context, id string, customerID *string, amount *int, createdAt *string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveOrder(ctx context.Context, id string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddOrderedProduct(ctx context.Context, orderID string, subProductID string, amount int) (*model.OrderedProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditOrderedProduct(ctx context.Context, id string, orderID *string, subProductID *string, amount *int) (*model.OrderedProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveOrderedProduct(ctx context.Context, id string) (*model.OrderedProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, login string, password string) (*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditAdmin(ctx context.Context, id string, login *string, password *string) (*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAdmin(ctx context.Context, id string) (*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.Prisma.Category.FindMany().Exec(ctx)

	var categories_ []*model.Category
	for _, catecategory := range categories {
		categories_ = append(categories_, &model.Category{
			ID:       catecategory.ID,
			Name:     catecategory.Name,
			ImageURL: catecategory.ImageURL,
		})
	}
	return categories_, err
}

func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	category, err := r.Prisma.Category.FindUnique(
		db.Category.ID.Equals(id),
	).Exec(ctx)

	return &model.Category{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}, err
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllProducts(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, categoryID string) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SubProducts(ctx context.Context, productID string) ([]*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context, customerID string) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OrderedSubProducts(ctx context.Context, orderID string) ([]*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OrderedSubProduct(ctx context.Context, id string) (*model.SubProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Admins(ctx context.Context) ([]*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Admin(ctx context.Context, id string) (*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
