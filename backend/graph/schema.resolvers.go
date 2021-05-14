package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-server/graph/generated"
	"graphql-server/hash"
	"graphql-server/prisma/db"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, name string, imageURL string) (*db.CategoryModel, error) {
	return r.Prisma.Category.CreateOne(
		db.Category.ID.Set(uuid.NewString()),
		db.Category.Name.Set(name),
		db.Category.ImageURL.Set(imageURL),
	).Exec(ctx)
}

func (r *mutationResolver) EditCategory(ctx context.Context, id string, name *string, imageURL *string) (*db.CategoryModel, error) {
	return r.Prisma.Category.FindUnique(
		db.Category.ID.Equals(id),
	).With(
		db.Category.Products.Fetch(),
	).Update(
		db.Category.Name.Set(*name),
		db.Category.ImageURL.Set(*imageURL),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*db.CategoryModel, error) {
	return r.Prisma.Category.FindUnique(db.Category.ID.Equals(id)).With(db.Category.Products.Fetch()).Delete().Exec(ctx)
}

func (r *mutationResolver) AddProduct(ctx context.Context, categoryID string, name string, description string, fabric string, manufacturer string) (*db.ProductModel, error) {
	_, err := r.Prisma.Category.FindUnique(
		db.Category.ID.Equals(categoryID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Product.CreateOne(
		db.Product.ID.Set(uuid.NewString()),
		db.Product.Name.Set(name),
		db.Product.Description.Set(description),
		db.Product.Manufacturer.Set(manufacturer),
		db.Product.Fabric.Set(fabric),
		db.Product.CategoryID.Set(categoryID),
	).Exec(ctx)
}

func (r *mutationResolver) EditProduct(ctx context.Context, id string, name *string, description *string) (*db.ProductModel, error) {
	return r.Prisma.Product.FindUnique(
		db.Product.ID.Equals(id),
	).With(
		db.Product.Images.Fetch(),
	).Update(
		db.Product.Name.Set(*name),
		db.Product.Description.Set(*description),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*db.ProductModel, error) {
	return r.Prisma.Product.FindUnique(db.Product.ID.Equals(id)).With(db.Product.Images.Fetch(), db.Product.SubProducts.Fetch()).Delete().Exec(ctx)
}

func (r *mutationResolver) AddImage(ctx context.Context, productID string, url string) (*db.ImageModel, error) {
	_, err := r.Prisma.Product.FindUnique(
		db.Product.ID.Equals(productID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Image.CreateOne(
		db.Image.ID.Set(uuid.NewString()),
		db.Image.URL.Set(url),
		db.Image.ProductID.Set(productID),
	).Exec(ctx)
}

func (r *mutationResolver) EditImage(ctx context.Context, id string, url *string) (*db.ImageModel, error) {
	return r.Prisma.Image.FindUnique(
		db.Image.ID.Equals(id),
	).Update(
		db.Image.URL.Set(*url),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteImage(ctx context.Context, id string) (*db.ImageModel, error) {
	return r.Prisma.Image.FindUnique(db.Image.ID.Equals(id)).Delete().Exec(ctx)
}

func (r *mutationResolver) CreateSubProduct(ctx context.Context, productID string, price float64, size string, color string, amount int) (*db.SubProductModel, error) {
	_, err := r.Prisma.Product.FindUnique(
		db.Product.ID.Equals(productID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.SubProduct.CreateOne(
		db.SubProduct.ID.Set(uuid.NewString()),
		db.SubProduct.Price.Set(price),
		db.SubProduct.Size.Set(size),
		db.SubProduct.Color.Set(color),
		db.SubProduct.Amount.Set(amount),
		db.SubProduct.ProductID.Set(productID),
	).Exec(ctx)
}

func (r *mutationResolver) EditSubProduct(ctx context.Context, id string, price *float64, size *string, color *string, amount *int) (*db.SubProductModel, error) {
	return r.Prisma.SubProduct.FindUnique(
		db.SubProduct.ID.Equals(id),
	).Update(
		db.SubProduct.Price.Set(*price),
		db.SubProduct.Size.Set(*size),
		db.SubProduct.Color.Set(*color),
		db.SubProduct.Amount.Set(*amount),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteSubProduct(ctx context.Context, id string) (*db.SubProductModel, error) {
	return r.Prisma.SubProduct.FindUnique(db.SubProduct.ID.Equals(id)).Delete().Exec(ctx)
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, name string, email string, phone string, address string, region string, ccNumber string) (*db.CustomerModel, error) {
	return r.Prisma.Customer.CreateOne(
		db.Customer.ID.Set(uuid.NewString()),
		db.Customer.Name.Set(name),
		db.Customer.Email.Set(email),
		db.Customer.Phone.Set(phone),
		db.Customer.Address.Set(address),
		db.Customer.Region.Set(region),
		db.Customer.CcNumber.Set(ccNumber),
	).Exec(ctx)
}

func (r *mutationResolver) EditCustomer(ctx context.Context, id string, name *string, email *string, phone *string, address *string, region *string, ccNumber *string) (*db.CustomerModel, error) {
	return r.Prisma.Customer.FindUnique(
		db.Customer.ID.Equals(id),
	).With(
		db.Customer.Orders.Fetch(),
	).Update(
		db.Customer.Name.Set(*name),
		db.Customer.Email.Set(*email),
		db.Customer.Phone.Set(*phone),
		db.Customer.Address.Set(*address),
		db.Customer.Region.Set(*region),
		db.Customer.CcNumber.Set(*ccNumber),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (*db.CustomerModel, error) {
	return r.Prisma.Customer.FindUnique(
		db.Customer.ID.Equals(id),
	).With(
		db.Customer.Orders.Fetch(),
	).Delete().Exec(ctx)
}

func (r *mutationResolver) AddOrder(ctx context.Context, customerID string, amount int, createdAt string) (*db.OrderModel, error) {
	_, err := r.Prisma.Customer.FindUnique(
		db.Customer.ID.Equals(customerID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Order.CreateOne(
		db.Order.ID.Set(uuid.NewString()),
		db.Order.Amount.Set(amount),
		db.Order.CreatedAt.Set(createdAt),
		db.Order.CustomerID.Set(customerID),
	).Exec(ctx)
}

func (r *mutationResolver) EditOrder(ctx context.Context, id string, customerID *string, amount *int, createdAt *string) (*db.OrderModel, error) {
	return r.Prisma.Order.FindUnique(
		db.Order.ID.Equals(id),
	).With(
		db.Order.OrderedProducts.Fetch(),
	).Update(
		db.Order.Amount.Set(*amount),
		db.Order.CreatedAt.Set(*createdAt),
		db.Order.CustomerID.Set(*customerID),
	).Exec(ctx)
}

func (r *mutationResolver) RemoveOrder(ctx context.Context, id string) (*db.OrderModel, error) {
	return r.Prisma.Order.FindUnique(db.Order.ID.Equals(id)).With(db.Order.OrderedProducts.Fetch()).Delete().Exec(ctx)
}

func (r *mutationResolver) AddOrderedProduct(ctx context.Context, orderID string, subProductID string, amount int) (*db.OrderedProductModel, error) {
	_, err := r.Prisma.Order.FindUnique(db.Order.ID.Equals(orderID)).With(db.Order.OrderedProducts.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.OrderedProduct.CreateOne(
		db.OrderedProduct.ID.Set(uuid.NewString()),
		db.OrderedProduct.Amount.Set(amount),
		db.OrderedProduct.SubProductID.Set(subProductID),
	).Exec(ctx)
}

func (r *mutationResolver) EditOrderedProduct(ctx context.Context, id string, orderID *string, subProductID *string, amount *int) (*db.OrderedProductModel, error) {
	return r.Prisma.OrderedProduct.FindUnique(
		db.OrderedProduct.ID.Equals(id),
	).Update(
		db.OrderedProduct.Amount.Set(*amount),
		db.OrderedProduct.SubProductID.Set(*subProductID),
	).Exec(ctx)
}

func (r *mutationResolver) RemoveOrderedProduct(ctx context.Context, id string) (*db.OrderedProductModel, error) {
	return r.Prisma.OrderedProduct.FindUnique(db.OrderedProduct.ID.Equals(id)).Delete().Exec(ctx)
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, login string, password string) (*db.AdminModel, error) {
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return r.Prisma.Admin.CreateOne(
		db.Admin.ID.Set(uuid.NewString()),
		db.Admin.Login.Set(login),
		db.Admin.Password.Set(hashedPassword),
	).Exec(ctx)
}

func (r *mutationResolver) EditAdmin(ctx context.Context, id string, login *string, password *string) (*db.AdminModel, error) {
	hashedPassword, err := hash.HashPassword(*password)
	if err != nil {
		return nil, err
	}
	return r.Prisma.Admin.FindUnique(
		db.Admin.ID.Equals(id),
	).Update(
		db.Admin.Login.Set(*login),
		db.Admin.Password.Set(hashedPassword),
	).Exec(ctx)
}

func (r *mutationResolver) DeleteAdmin(ctx context.Context, id string) (*db.AdminModel, error) {
	return r.Prisma.Admin.FindUnique(db.Admin.ID.Equals(id)).Delete().Exec(ctx)
}

func (r *queryResolver) Categories(ctx context.Context) ([]*db.CategoryModel, error) {
	categories, err := r.Prisma.Category.FindMany().With(db.Category.Products.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var categoriesPointers []*db.CategoryModel
	for i := range categories {
		categoriesPointers = append(categoriesPointers, &categories[i])
	}
	return categoriesPointers, nil
}

func (r *queryResolver) Category(ctx context.Context, id string) (*db.CategoryModel, error) {
	return r.Prisma.Category.FindUnique(db.Category.ID.Equals(id)).With(db.Category.Products.Fetch()).Exec(ctx)
}

func (r *queryResolver) Customers(ctx context.Context) ([]*db.CustomerModel, error) {
	customers, err := r.Prisma.Customer.FindMany().With(db.Customer.Orders.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var customersPointers []*db.CustomerModel
	for i := range customers {
		customersPointers = append(customersPointers, &customers[i])
	}
	return customersPointers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*db.CustomerModel, error) {
	return r.Prisma.Customer.FindUnique(
		db.Customer.ID.Equals(id),
	).Exec(ctx)
}

func (r *queryResolver) AllProducts(ctx context.Context) ([]*db.ProductModel, error) {
	products, err := r.Prisma.Product.FindMany().With(db.Product.SubProducts.Fetch(), db.Product.Images.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var productsPointers []*db.ProductModel
	for i := range products {
		productsPointers = append(productsPointers, &products[i])
	}
	return productsPointers, nil
}

func (r *queryResolver) Products(ctx context.Context, categoryID string) ([]*db.ProductModel, error) {
	products, err := r.Prisma.Product.FindMany(db.Product.CategoryID.Equals(categoryID)).With(db.Product.SubProducts.Fetch(), db.Product.Images.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var productsPointers []*db.ProductModel
	for i := range products {
		productsPointers = append(productsPointers, &products[i])
	}
	return productsPointers, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*db.ProductModel, error) {
	return r.Prisma.Product.FindUnique(db.Product.ID.Equals(id)).With(db.Product.SubProducts.Fetch(), db.Product.Images.Fetch()).Exec(ctx)
}

func (r *queryResolver) Images(ctx context.Context, productID string) ([]*db.ImageModel, error) {
	images, err := r.Prisma.Image.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	var imagesPointers []*db.ImageModel
	for i := range images {
		imagesPointers = append(imagesPointers, &images[i])
	}
	return imagesPointers, nil
}

func (r *queryResolver) Image(ctx context.Context, id string) (*db.ImageModel, error) {
	return r.Prisma.Image.FindUnique(
		db.Image.ID.Equals(id),
	).Exec(ctx)
}

func (r *queryResolver) SubProducts(ctx context.Context, productID string) ([]*db.SubProductModel, error) {
	subProducts, err := r.Prisma.SubProduct.FindMany(
		db.SubProduct.ProductID.Equals(productID),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var subProductsPointers []*db.SubProductModel
	for i := range subProducts {
		subProductsPointers = append(subProductsPointers, &subProducts[i])
	}
	return subProductsPointers, nil
}

func (r *queryResolver) SubProduct(ctx context.Context, id string) (*db.SubProductModel, error) {
	return r.Prisma.SubProduct.FindUnique(
		db.SubProduct.ID.Equals(id),
	).Exec(ctx)
}

func (r *queryResolver) Orders(ctx context.Context, customerID string) ([]*db.OrderModel, error) {
	orders, err := r.Prisma.Order.FindMany(db.Order.CustomerID.Equals(customerID)).With(db.Order.OrderedProducts.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var ordersPointers []*db.OrderModel
	for i := range orders {
		ordersPointers = append(ordersPointers, &orders[i])
	}
	return ordersPointers, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*db.OrderModel, error) {
	return r.Prisma.Order.FindUnique(
		db.Order.ID.Equals(id),
	).Exec(ctx)
}

func (r *queryResolver) OrderedProducts(ctx context.Context, orderID string) ([]*db.OrderedProductModel, error) {
	orderedProducts, err := r.Prisma.OrderedProduct.FindMany(db.OrderedProduct.OrderID.Equals(orderID)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var orderedProductsPointers []*db.OrderedProductModel
	for i := range orderedProducts {
		orderedProductsPointers = append(orderedProductsPointers, &orderedProducts[i])
	}
	return orderedProductsPointers, nil
}

func (r *queryResolver) OrderedProduct(ctx context.Context, id string) (*db.OrderedProductModel, error) {
	return r.Prisma.OrderedProduct.FindUnique(
		db.OrderedProduct.ID.Equals(id),
	).Exec(ctx)
}

func (r *queryResolver) Admins(ctx context.Context) ([]*db.AdminModel, error) {
	admins, err := r.Prisma.Admin.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	var adminsPointers []*db.AdminModel
	for i := range admins {
		adminsPointers = append(adminsPointers, &admins[i])
	}
	return adminsPointers, nil
}

func (r *queryResolver) Admin(ctx context.Context, id string) (*db.AdminModel, error) {
	return r.Prisma.Admin.FindUnique(
		db.Admin.ID.Equals(id),
	).Exec(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
