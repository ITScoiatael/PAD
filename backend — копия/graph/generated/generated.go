package generated

import (
	"bytes"
	"context"
	"errors"
	"graphql-server/prisma/db"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Admin struct {
		ID       func(childComplexity int) int
		Login    func(childComplexity int) int
		Password func(childComplexity int) int
	}

	Category struct {
		ID       func(childComplexity int) int
		ImageURL func(childComplexity int) int
		Name     func(childComplexity int) int
		Products func(childComplexity int) int
	}

	Customer struct {
		Address  func(childComplexity int) int
		CcNumber func(childComplexity int) int
		Email    func(childComplexity int) int
		ID       func(childComplexity int) int
		Name     func(childComplexity int) int
		Orders   func(childComplexity int) int
		Phone    func(childComplexity int) int
		Region   func(childComplexity int) int
	}

	Mutation struct {
		AddOrder             func(childComplexity int, customerID string, amount int, createdAt string) int
		AddOrderedProduct    func(childComplexity int, orderID string, subProductID string, amount int) int
		AddProduct           func(childComplexity int, categoryID string, name string, description string, imageURL string) int
		CreateAdmin          func(childComplexity int, login string, password string) int
		CreateCategory       func(childComplexity int, name string, imageURL string) int
		CreateCustomer       func(childComplexity int, name string, email string, phone string, address string, region string, ccNumber string) int
		CreateSubProduct     func(childComplexity int, productID string, price float64, size string, color string, amount int) int
		DeleteAdmin          func(childComplexity int, id string) int
		DeleteCategory       func(childComplexity int, id string) int
		DeleteCustomer       func(childComplexity int, id string) int
		DeleteProduct        func(childComplexity int, id string) int
		DeleteSubProduct     func(childComplexity int, id string) int
		EditAdmin            func(childComplexity int, id string, login *string, password *string) int
		EditCategory         func(childComplexity int, id string, name *string, imageURL *string) int
		EditCustomer         func(childComplexity int, id string, name *string, email *string, phone *string, address *string, region *string, ccNumber *string) int
		EditOrder            func(childComplexity int, id string, customerID *string, amount *int, createdAt *string) int
		EditOrderedProduct   func(childComplexity int, id string, orderID *string, subProductID *string, amount *int) int
		EditProduct          func(childComplexity int, id string, name *string, description *string, imageURL *string) int
		EditSubProduct       func(childComplexity int, id string, price *float64, size *string, color *string, amount *int) int
		RemoveOrder          func(childComplexity int, id string) int
		RemoveOrderedProduct func(childComplexity int, id string) int
	}

	Order struct {
		Amount          func(childComplexity int) int
		CreatedAt       func(childComplexity int) int
		ID              func(childComplexity int) int
		OrderedProducts func(childComplexity int) int
	}

	OrderedProduct struct {
		ID         func(childComplexity int) int
		SubProduct func(childComplexity int) int
	}

	Product struct {
		Description func(childComplexity int) int
		ID          func(childComplexity int) int
		ImageURL    func(childComplexity int) int
		Name        func(childComplexity int) int
		SubProducts func(childComplexity int) int
	}

	Query struct {
		Admin           func(childComplexity int, id string) int
		Admins          func(childComplexity int) int
		AllProducts     func(childComplexity int) int
		Categories      func(childComplexity int) int
		Category        func(childComplexity int, id string) int
		Customer        func(childComplexity int, id string) int
		Customers       func(childComplexity int) int
		Order           func(childComplexity int, id string) int
		OrderedProduct  func(childComplexity int, id string) int
		OrderedProducts func(childComplexity int, orderID string) int
		Orders          func(childComplexity int, customerID string) int
		Product         func(childComplexity int, id string) int
		Products        func(childComplexity int, categoryID string) int
		SubProduct      func(childComplexity int, id string) int
		SubProducts     func(childComplexity int, productID string) int
	}

	SubProduct struct {
		Amount func(childComplexity int) int
		Color  func(childComplexity int) int
		ID     func(childComplexity int) int
		Price  func(childComplexity int) int
		Size   func(childComplexity int) int
	}
}

type MutationResolver interface {
	CreateCategory(ctx context.Context, name string, imageURL string) (*db.CategoryModel, error)
	EditCategory(ctx context.Context, id string, name *string, imageURL *string) (*db.CategoryModel, error)
	DeleteCategory(ctx context.Context, id string) (*db.CategoryModel, error)
	AddProduct(ctx context.Context, categoryID string, name string, description string, imageURL string) (*db.ProductModel, error)
	EditProduct(ctx context.Context, id string, name *string, description *string, imageURL *string) (*db.ProductModel, error)
	DeleteProduct(ctx context.Context, id string) (*db.ProductModel, error)
	CreateSubProduct(ctx context.Context, productID string, price float64, size string, color string, amount int) (*db.SubProductModel, error)
	EditSubProduct(ctx context.Context, id string, price *float64, size *string, color *string, amount *int) (*db.SubProductModel, error)
	DeleteSubProduct(ctx context.Context, id string) (*db.SubProductModel, error)
	CreateCustomer(ctx context.Context, name string, email string, phone string, address string, region string, ccNumber string) (*db.CustomerModel, error)
	EditCustomer(ctx context.Context, id string, name *string, email *string, phone *string, address *string, region *string, ccNumber *string) (*db.CustomerModel, error)
	DeleteCustomer(ctx context.Context, id string) (*db.CustomerModel, error)
	AddOrder(ctx context.Context, customerID string, amount int, createdAt string) (*db.OrderModel, error)
	EditOrder(ctx context.Context, id string, customerID *string, amount *int, createdAt *string) (*db.OrderModel, error)
	RemoveOrder(ctx context.Context, id string) (*db.OrderModel, error)
	AddOrderedProduct(ctx context.Context, orderID string, subProductID string, amount int) (*db.OrderedProductModel, error)
	EditOrderedProduct(ctx context.Context, id string, orderID *string, subProductID *string, amount *int) (*db.OrderedProductModel, error)
	RemoveOrderedProduct(ctx context.Context, id string) (*db.OrderedProductModel, error)
	CreateAdmin(ctx context.Context, login string, password string) (*db.AdminModel, error)
	EditAdmin(ctx context.Context, id string, login *string, password *string) (*db.AdminModel, error)
	DeleteAdmin(ctx context.Context, id string) (*db.AdminModel, error)
}
type QueryResolver interface {
	Categories(ctx context.Context) ([]db.CategoryModel, error)
	Category(ctx context.Context, id string) (*db.CategoryModel, error)
	Customers(ctx context.Context) ([]db.CustomerModel, error)
	Customer(ctx context.Context, id string) (*db.CustomerModel, error)
	AllProducts(ctx context.Context) ([]db.ProductModel, error)
	Products(ctx context.Context, categoryID string) ([]db.ProductModel, error)
	Product(ctx context.Context, id string) (*db.ProductModel, error)
	SubProducts(ctx context.Context, productID string) ([]db.SubProductModel, error)
	SubProduct(ctx context.Context, id string) (*db.SubProductModel, error)
	Orders(ctx context.Context, customerID string) ([]db.OrderModel, error)
	Order(ctx context.Context, id string) (*db.OrderModel, error)
	OrderedProducts(ctx context.Context, orderID string) ([]db.OrderedProductModel, error)
	OrderedProduct(ctx context.Context, id string) (*db.OrderedProductModel, error)
	Admins(ctx context.Context) ([]db.AdminModel, error)
	Admin(ctx context.Context, id string) (*db.AdminModel, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Admin.id":
		if e.complexity.Admin.ID == nil {
			break
		}

		return e.complexity.Admin.ID(childComplexity), true

	case "Admin.login":
		if e.complexity.Admin.Login == nil {
			break
		}

		return e.complexity.Admin.Login(childComplexity), true

	case "Admin.password":
		if e.complexity.Admin.Password == nil {
			break
		}

		return e.complexity.Admin.Password(childComplexity), true

	case "Category.id":
		if e.complexity.Category.ID == nil {
			break
		}

		return e.complexity.Category.ID(childComplexity), true

	case "Category.image_url":
		if e.complexity.Category.ImageURL == nil {
			break
		}

		return e.complexity.Category.ImageURL(childComplexity), true

	case "Category.name":
		if e.complexity.Category.Name == nil {
			break
		}

		return e.complexity.Category.Name(childComplexity), true

	case "Category.products":
		if e.complexity.Category.Products == nil {
			break
		}

		return e.complexity.Category.Products(childComplexity), true

	case "Customer.address":
		if e.complexity.Customer.Address == nil {
			break
		}

		return e.complexity.Customer.Address(childComplexity), true

	case "Customer.cc_number":
		if e.complexity.Customer.CcNumber == nil {
			break
		}

		return e.complexity.Customer.CcNumber(childComplexity), true

	case "Customer.email":
		if e.complexity.Customer.Email == nil {
			break
		}

		return e.complexity.Customer.Email(childComplexity), true

	case "Customer.id":
		if e.complexity.Customer.ID == nil {
			break
		}

		return e.complexity.Customer.ID(childComplexity), true

	case "Customer.name":
		if e.complexity.Customer.Name == nil {
			break
		}

		return e.complexity.Customer.Name(childComplexity), true

	case "Customer.orders":
		if e.complexity.Customer.Orders == nil {
			break
		}

		return e.complexity.Customer.Orders(childComplexity), true

	case "Customer.phone":
		if e.complexity.Customer.Phone == nil {
			break
		}

		return e.complexity.Customer.Phone(childComplexity), true

	case "Customer.region":
		if e.complexity.Customer.Region == nil {
			break
		}

		return e.complexity.Customer.Region(childComplexity), true

	case "Mutation.addOrder":
		if e.complexity.Mutation.AddOrder == nil {
			break
		}

		args, err := ec.field_Mutation_addOrder_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.AddOrder(childComplexity, args["customer_id"].(string), args["amount"].(int), args["created_at"].(string)), true

	case "Mutation.addOrderedProduct":
		if e.complexity.Mutation.AddOrderedProduct == nil {
			break
		}

		args, err := ec.field_Mutation_addOrderedProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.AddOrderedProduct(childComplexity, args["order_id"].(string), args["sub_product_id"].(string), args["amount"].(int)), true

	case "Mutation.addProduct":
		if e.complexity.Mutation.AddProduct == nil {
			break
		}

		args, err := ec.field_Mutation_addProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.AddProduct(childComplexity, args["category_id"].(string), args["name"].(string), args["description"].(string), args["image_url"].(string)), true

	case "Mutation.createAdmin":
		if e.complexity.Mutation.CreateAdmin == nil {
			break
		}

		args, err := ec.field_Mutation_createAdmin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateAdmin(childComplexity, args["login"].(string), args["password"].(string)), true

	case "Mutation.createCategory":
		if e.complexity.Mutation.CreateCategory == nil {
			break
		}

		args, err := ec.field_Mutation_createCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateCategory(childComplexity, args["name"].(string), args["image_url"].(string)), true

	case "Mutation.createCustomer":
		if e.complexity.Mutation.CreateCustomer == nil {
			break
		}

		args, err := ec.field_Mutation_createCustomer_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateCustomer(childComplexity, args["name"].(string), args["email"].(string), args["phone"].(string), args["address"].(string), args["region"].(string), args["cc_number"].(string)), true

	case "Mutation.createSubProduct":
		if e.complexity.Mutation.CreateSubProduct == nil {
			break
		}

		args, err := ec.field_Mutation_createSubProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateSubProduct(childComplexity, args["product_id"].(string), args["price"].(float64), args["size"].(string), args["color"].(string), args["amount"].(int)), true

	case "Mutation.deleteAdmin":
		if e.complexity.Mutation.DeleteAdmin == nil {
			break
		}

		args, err := ec.field_Mutation_deleteAdmin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteAdmin(childComplexity, args["id"].(string)), true

	case "Mutation.deleteCategory":
		if e.complexity.Mutation.DeleteCategory == nil {
			break
		}

		args, err := ec.field_Mutation_deleteCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteCategory(childComplexity, args["id"].(string)), true

	case "Mutation.deleteCustomer":
		if e.complexity.Mutation.DeleteCustomer == nil {
			break
		}

		args, err := ec.field_Mutation_deleteCustomer_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteCustomer(childComplexity, args["id"].(string)), true

	case "Mutation.deleteProduct":
		if e.complexity.Mutation.DeleteProduct == nil {
			break
		}

		args, err := ec.field_Mutation_deleteProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteProduct(childComplexity, args["id"].(string)), true

	case "Mutation.deleteSubProduct":
		if e.complexity.Mutation.DeleteSubProduct == nil {
			break
		}

		args, err := ec.field_Mutation_deleteSubProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteSubProduct(childComplexity, args["id"].(string)), true

	case "Mutation.editAdmin":
		if e.complexity.Mutation.EditAdmin == nil {
			break
		}

		args, err := ec.field_Mutation_editAdmin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditAdmin(childComplexity, args["id"].(string), args["login"].(*string), args["password"].(*string)), true

	case "Mutation.editCategory":
		if e.complexity.Mutation.EditCategory == nil {
			break
		}

		args, err := ec.field_Mutation_editCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditCategory(childComplexity, args["id"].(string), args["name"].(*string), args["image_url"].(*string)), true

	case "Mutation.editCustomer":
		if e.complexity.Mutation.EditCustomer == nil {
			break
		}

		args, err := ec.field_Mutation_editCustomer_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditCustomer(childComplexity, args["id"].(string), args["name"].(*string), args["email"].(*string), args["phone"].(*string), args["address"].(*string), args["region"].(*string), args["cc_number"].(*string)), true

	case "Mutation.editOrder":
		if e.complexity.Mutation.EditOrder == nil {
			break
		}

		args, err := ec.field_Mutation_editOrder_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditOrder(childComplexity, args["id"].(string), args["customer_id"].(*string), args["amount"].(*int), args["created_at"].(*string)), true

	case "Mutation.editOrderedProduct":
		if e.complexity.Mutation.EditOrderedProduct == nil {
			break
		}

		args, err := ec.field_Mutation_editOrderedProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditOrderedProduct(childComplexity, args["id"].(string), args["order_id"].(*string), args["sub_product_id"].(*string), args["amount"].(*int)), true

	case "Mutation.editProduct":
		if e.complexity.Mutation.EditProduct == nil {
			break
		}

		args, err := ec.field_Mutation_editProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditProduct(childComplexity, args["id"].(string), args["name"].(*string), args["description"].(*string), args["image_url"].(*string)), true

	case "Mutation.editSubProduct":
		if e.complexity.Mutation.EditSubProduct == nil {
			break
		}

		args, err := ec.field_Mutation_editSubProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditSubProduct(childComplexity, args["id"].(string), args["price"].(*float64), args["size"].(*string), args["color"].(*string), args["amount"].(*int)), true

	case "Mutation.removeOrder":
		if e.complexity.Mutation.RemoveOrder == nil {
			break
		}

		args, err := ec.field_Mutation_removeOrder_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveOrder(childComplexity, args["id"].(string)), true

	case "Mutation.removeOrderedProduct":
		if e.complexity.Mutation.RemoveOrderedProduct == nil {
			break
		}

		args, err := ec.field_Mutation_removeOrderedProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveOrderedProduct(childComplexity, args["id"].(string)), true

	case "Order.amount":
		if e.complexity.Order.Amount == nil {
			break
		}

		return e.complexity.Order.Amount(childComplexity), true

	case "Order.created_at":
		if e.complexity.Order.CreatedAt == nil {
			break
		}

		return e.complexity.Order.CreatedAt(childComplexity), true

	case "Order.id":
		if e.complexity.Order.ID == nil {
			break
		}

		return e.complexity.Order.ID(childComplexity), true

	case "Order.ordered_products":
		if e.complexity.Order.OrderedProducts == nil {
			break
		}

		return e.complexity.Order.OrderedProducts(childComplexity), true

	case "Ordered_Product.id":
		if e.complexity.OrderedProduct.ID == nil {
			break
		}

		return e.complexity.OrderedProduct.ID(childComplexity), true

	case "Ordered_Product.sub_product":
		if e.complexity.OrderedProduct.SubProduct == nil {
			break
		}

		return e.complexity.OrderedProduct.SubProduct(childComplexity), true

	case "Product.description":
		if e.complexity.Product.Description == nil {
			break
		}

		return e.complexity.Product.Description(childComplexity), true

	case "Product.id":
		if e.complexity.Product.ID == nil {
			break
		}

		return e.complexity.Product.ID(childComplexity), true

	case "Product.image_url":
		if e.complexity.Product.ImageURL == nil {
			break
		}

		return e.complexity.Product.ImageURL(childComplexity), true

	case "Product.name":
		if e.complexity.Product.Name == nil {
			break
		}

		return e.complexity.Product.Name(childComplexity), true

	case "Product.sub_products":
		if e.complexity.Product.SubProducts == nil {
			break
		}

		return e.complexity.Product.SubProducts(childComplexity), true

	case "Query.Admin":
		if e.complexity.Query.Admin == nil {
			break
		}

		args, err := ec.field_Query_Admin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Admin(childComplexity, args["id"].(string)), true

	case "Query.Admins":
		if e.complexity.Query.Admins == nil {
			break
		}

		return e.complexity.Query.Admins(childComplexity), true

	case "Query.AllProducts":
		if e.complexity.Query.AllProducts == nil {
			break
		}

		return e.complexity.Query.AllProducts(childComplexity), true

	case "Query.Categories":
		if e.complexity.Query.Categories == nil {
			break
		}

		return e.complexity.Query.Categories(childComplexity), true

	case "Query.Category":
		if e.complexity.Query.Category == nil {
			break
		}

		args, err := ec.field_Query_Category_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Category(childComplexity, args["id"].(string)), true

	case "Query.Customer":
		if e.complexity.Query.Customer == nil {
			break
		}

		args, err := ec.field_Query_Customer_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Customer(childComplexity, args["id"].(string)), true

	case "Query.Customers":
		if e.complexity.Query.Customers == nil {
			break
		}

		return e.complexity.Query.Customers(childComplexity), true

	case "Query.Order":
		if e.complexity.Query.Order == nil {
			break
		}

		args, err := ec.field_Query_Order_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Order(childComplexity, args["id"].(string)), true

	case "Query.OrderedProduct":
		if e.complexity.Query.OrderedProduct == nil {
			break
		}

		args, err := ec.field_Query_OrderedProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.OrderedProduct(childComplexity, args["id"].(string)), true

	case "Query.OrderedProducts":
		if e.complexity.Query.OrderedProducts == nil {
			break
		}

		args, err := ec.field_Query_OrderedProducts_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.OrderedProducts(childComplexity, args["order_id"].(string)), true

	case "Query.Orders":
		if e.complexity.Query.Orders == nil {
			break
		}

		args, err := ec.field_Query_Orders_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Orders(childComplexity, args["customer_id"].(string)), true

	case "Query.Product":
		if e.complexity.Query.Product == nil {
			break
		}

		args, err := ec.field_Query_Product_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Product(childComplexity, args["id"].(string)), true

	case "Query.Products":
		if e.complexity.Query.Products == nil {
			break
		}

		args, err := ec.field_Query_Products_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Products(childComplexity, args["category_id"].(string)), true

	case "Query.SubProduct":
		if e.complexity.Query.SubProduct == nil {
			break
		}

		args, err := ec.field_Query_SubProduct_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.SubProduct(childComplexity, args["id"].(string)), true

	case "Query.SubProducts":
		if e.complexity.Query.SubProducts == nil {
			break
		}

		args, err := ec.field_Query_SubProducts_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.SubProducts(childComplexity, args["product_id"].(string)), true

	case "Sub_Product.amount":
		if e.complexity.SubProduct.Amount == nil {
			break
		}

		return e.complexity.SubProduct.Amount(childComplexity), true

	case "Sub_Product.color":
		if e.complexity.SubProduct.Color == nil {
			break
		}

		return e.complexity.SubProduct.Color(childComplexity), true

	case "Sub_Product.id":
		if e.complexity.SubProduct.ID == nil {
			break
		}

		return e.complexity.SubProduct.ID(childComplexity), true

	case "Sub_Product.price":
		if e.complexity.SubProduct.Price == nil {
			break
		}

		return e.complexity.SubProduct.Price(childComplexity), true

	case "Sub_Product.size":
		if e.complexity.SubProduct.Size == nil {
			break
		}

		return e.complexity.SubProduct.Size(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "graph/schema.graphqls", Input: `type Admin {
  id: ID!
  login: String!
  password: String!
}

type Category {
  id: ID!
  name: String!
  image_url: String!
  products: [Product!]
}

type Product {
  id: ID!
  name: String!
  description: String!
  image_url: String!
  sub_products: [Sub_Product!]
}

type Sub_Product {
  id: ID!
  price: Float!
  size: String!
  color: String!
  amount: Int!
}

type Ordered_Product {
  id: ID!
  sub_product: Sub_Product!
}

type Customer {
  id: ID!
  name: String!
  email: String!
  phone: String!
  address: String!
  region: String!
  cc_number: String!
  orders: [Order!]
}

type Order {
  id: ID!
  amount: Int!
  created_at: Date!
  ordered_products: [Ordered_Product!]!
}

scalar Date


type Query {
  Categories: [Category]
  Category(id: ID!): Category

  Customers: [Customer]
  Customer(id: ID!): Customer

  AllProducts: [Product]
  Products(category_id: ID!): [Product]
  Product(id: ID!): Product

  SubProducts(product_id: ID!): [Sub_Product]
  SubProduct(id: ID!): Sub_Product

  Orders(customer_id: ID!): [Order]
  Order(id: ID!): Order

  OrderedProducts(order_id: ID!): [Ordered_Product]
  OrderedProduct(id: ID!): Ordered_Product

  Admins: [Admin]
  Admin(id: ID!): Admin
}

type Mutation {
  createCategory (
    name: String!
    image_url: String!
  ): Category
  editCategory(
    id: ID!
    name: String
    image_url: String
  ): Category
  deleteCategory(id: ID!): Category

  addProduct (
    category_id: ID!
    name: String!
    description: String!
    image_url: String!
  ) : Product
  editProduct (
    id: ID!
    name: String
    description: String
    image_url: String
  ) : Product
  deleteProduct (id: ID!): Product

  createSubProduct (
    product_id: ID!
    price: Float!
    size: String!
    color: String!
    amount: Int!
  ) : Sub_Product
  editSubProduct (
    id: ID!
    price: Float
    size: String
    color: String
    amount: Int
  ) : Sub_Product
  deleteSubProduct(id: ID!): Sub_Product

  createCustomer(
    name: String!
    email: String!
    phone: String!
    address: String!
    region: String!
    cc_number: String!
  ): Customer
  editCustomer(
    id: ID!
    name: String
    email: String
    phone: String
    address: String
    region: String
    cc_number: String
  ): Customer
  deleteCustomer(id: ID!): Customer

  addOrder(
    customer_id: ID!
    amount: Int!
    created_at: Date!
  ): Order
  editOrder(
    id: ID!
    customer_id: ID
    amount: Int
    created_at: Date
  ): Order
  removeOrder(id: ID!): Order
  
  addOrderedProduct(
    order_id: ID!
    sub_product_id: ID!
    amount: Int!
  ): Ordered_Product
  editOrderedProduct(
    id: ID!
    order_id: ID
    sub_product_id: ID
    amount: Int
  ): Ordered_Product
  removeOrderedProduct(id: ID!): Ordered_Product

  createAdmin(
    login: String!
    password: String!
  ) : Admin
  editAdmin(
    id: ID!
    login: String
    password: String
  ) : Admin
  deleteAdmin(id: ID!): Admin
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_addOrder_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["customer_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("customer_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["customer_id"] = arg0
	var arg1 int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg1, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["created_at"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("created_at"))
		arg2, err = ec.unmarshalNDate2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["created_at"] = arg2
	return args, nil
}

func (ec *executionContext) field_Mutation_addOrderedProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["order_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("order_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["order_id"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["sub_product_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sub_product_id"))
		arg1, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["sub_product_id"] = arg1
	var arg2 int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg2, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg2
	return args, nil
}

func (ec *executionContext) field_Mutation_addProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["category_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("category_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["category_id"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["description"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("description"))
		arg2, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["description"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["image_url"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("image_url"))
		arg3, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["image_url"] = arg3
	return args, nil
}

func (ec *executionContext) field_Mutation_createAdmin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["login"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("login"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["login"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["password"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["password"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_createCategory_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["image_url"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("image_url"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["image_url"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_createCustomer_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["phone"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("phone"))
		arg2, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["phone"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["address"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("address"))
		arg3, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["address"] = arg3
	var arg4 string
	if tmp, ok := rawArgs["region"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("region"))
		arg4, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["region"] = arg4
	var arg5 string
	if tmp, ok := rawArgs["cc_number"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("cc_number"))
		arg5, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["cc_number"] = arg5
	return args, nil
}

func (ec *executionContext) field_Mutation_createSubProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["product_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("product_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["product_id"] = arg0
	var arg1 float64
	if tmp, ok := rawArgs["price"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("price"))
		arg1, err = ec.unmarshalNFloat2float64(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["price"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["size"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("size"))
		arg2, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["size"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["color"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("color"))
		arg3, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["color"] = arg3
	var arg4 int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg4, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg4
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteAdmin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteCategory_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteCustomer_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteSubProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_editAdmin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["login"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("login"))
		arg1, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["login"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["password"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["password"] = arg2
	return args, nil
}

func (ec *executionContext) field_Mutation_editCategory_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg1, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["image_url"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("image_url"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["image_url"] = arg2
	return args, nil
}

func (ec *executionContext) field_Mutation_editCustomer_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg1, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg2
	var arg3 *string
	if tmp, ok := rawArgs["phone"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("phone"))
		arg3, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["phone"] = arg3
	var arg4 *string
	if tmp, ok := rawArgs["address"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("address"))
		arg4, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["address"] = arg4
	var arg5 *string
	if tmp, ok := rawArgs["region"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("region"))
		arg5, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["region"] = arg5
	var arg6 *string
	if tmp, ok := rawArgs["cc_number"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("cc_number"))
		arg6, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["cc_number"] = arg6
	return args, nil
}

func (ec *executionContext) field_Mutation_editOrder_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["customer_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("customer_id"))
		arg1, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["customer_id"] = arg1
	var arg2 *int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg2, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg2
	var arg3 *string
	if tmp, ok := rawArgs["created_at"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("created_at"))
		arg3, err = ec.unmarshalODate2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["created_at"] = arg3
	return args, nil
}

func (ec *executionContext) field_Mutation_editOrderedProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["order_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("order_id"))
		arg1, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["order_id"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["sub_product_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sub_product_id"))
		arg2, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["sub_product_id"] = arg2
	var arg3 *int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg3, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg3
	return args, nil
}

func (ec *executionContext) field_Mutation_editProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg1, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["description"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("description"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["description"] = arg2
	var arg3 *string
	if tmp, ok := rawArgs["image_url"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("image_url"))
		arg3, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["image_url"] = arg3
	return args, nil
}

func (ec *executionContext) field_Mutation_editSubProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 *float64
	if tmp, ok := rawArgs["price"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("price"))
		arg1, err = ec.unmarshalOFloat2ᚖfloat64(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["price"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["size"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("size"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["size"] = arg2
	var arg3 *string
	if tmp, ok := rawArgs["color"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("color"))
		arg3, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["color"] = arg3
	var arg4 *int
	if tmp, ok := rawArgs["amount"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("amount"))
		arg4, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["amount"] = arg4
	return args, nil
}

func (ec *executionContext) field_Mutation_removeOrder_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_removeOrderedProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Admin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Category_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Customer_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Order_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_OrderedProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_OrderedProducts_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["order_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("order_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["order_id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Orders_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["customer_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("customer_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["customer_id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Product_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_Products_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["category_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("category_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["category_id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_SubProduct_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_SubProducts_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["product_id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("product_id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["product_id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("includeDeprecated"))
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("includeDeprecated"))
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Admin_id(ctx context.Context, field graphql.CollectedField, obj *db.AdminModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Admin",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Admin_login(ctx context.Context, field graphql.CollectedField, obj *db.AdminModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Admin",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Login, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Admin_password(ctx context.Context, field graphql.CollectedField, obj *db.AdminModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Admin",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Password, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Category_id(ctx context.Context, field graphql.CollectedField, obj *db.CategoryModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Category",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Category_name(ctx context.Context, field graphql.CollectedField, obj *db.CategoryModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Category",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Category_image_url(ctx context.Context, field graphql.CollectedField, obj *db.CategoryModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Category",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ImageURL, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Category_products(ctx context.Context, field graphql.CollectedField, obj *db.CategoryModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Category",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Product, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐProductᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_id(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_name(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_email(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Email, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_phone(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Phone, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_address(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Address, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_region(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Region, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_cc_number(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CcNumber, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Customer_orders(ctx context.Context, field graphql.CollectedField, obj *db.CustomerModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Customer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Order, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_createCategory(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createCategory_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateCategory(rctx, args["name"].(string), args["image_url"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CategoryModel)
	fc.Result = res
	return ec.marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editCategory(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editCategory_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditCategory(rctx, args["id"].(string), args["name"].(*string), args["image_url"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CategoryModel)
	fc.Result = res
	return ec.marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteCategory(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteCategory_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteCategory(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CategoryModel)
	fc.Result = res
	return ec.marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_addProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_addProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().AddProduct(rctx, args["category_id"].(string), args["name"].(string), args["description"].(string), args["image_url"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditProduct(rctx, args["id"].(string), args["name"].(*string), args["description"].(*string), args["image_url"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteProduct(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_createSubProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createSubProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateSubProduct(rctx, args["product_id"].(string), args["price"].(float64), args["size"].(string), args["color"].(string), args["amount"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editSubProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editSubProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditSubProduct(rctx, args["id"].(string), args["price"].(*float64), args["size"].(*string), args["color"].(*string), args["amount"].(*int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteSubProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteSubProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteSubProduct(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_createCustomer(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createCustomer_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateCustomer(rctx, args["name"].(string), args["email"].(string), args["phone"].(string), args["address"].(string), args["region"].(string), args["cc_number"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CustomerModel)
	fc.Result = res
	return ec.marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editCustomer(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editCustomer_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditCustomer(rctx, args["id"].(string), args["name"].(*string), args["email"].(*string), args["phone"].(*string), args["address"].(*string), args["region"].(*string), args["cc_number"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CustomerModel)
	fc.Result = res
	return ec.marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteCustomer(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteCustomer_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteCustomer(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CustomerModel)
	fc.Result = res
	return ec.marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_addOrder(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_addOrder_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().AddOrder(rctx, args["customer_id"].(string), args["amount"].(int), args["created_at"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editOrder(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editOrder_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditOrder(rctx, args["id"].(string), args["customer_id"].(*string), args["amount"].(*int), args["created_at"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_removeOrder(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_removeOrder_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().RemoveOrder(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_addOrderedProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_addOrderedProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().AddOrderedProduct(rctx, args["order_id"].(string), args["sub_product_id"].(string), args["amount"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderedProductModel)
	fc.Result = res
	return ec.marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editOrderedProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editOrderedProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditOrderedProduct(rctx, args["id"].(string), args["order_id"].(*string), args["sub_product_id"].(*string), args["amount"].(*int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderedProductModel)
	fc.Result = res
	return ec.marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_removeOrderedProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_removeOrderedProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().RemoveOrderedProduct(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderedProductModel)
	fc.Result = res
	return ec.marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_createAdmin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createAdmin_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateAdmin(rctx, args["login"].(string), args["password"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.AdminModel)
	fc.Result = res
	return ec.marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_editAdmin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_editAdmin_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().EditAdmin(rctx, args["id"].(string), args["login"].(*string), args["password"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.AdminModel)
	fc.Result = res
	return ec.marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteAdmin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteAdmin_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteAdmin(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.AdminModel)
	fc.Result = res
	return ec.marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, field.Selections, res)
}

func (ec *executionContext) _Order_id(ctx context.Context, field graphql.CollectedField, obj *db.OrderModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Order",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Order_amount(ctx context.Context, field graphql.CollectedField, obj *db.OrderModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Order",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Amount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Order_created_at(ctx context.Context, field graphql.CollectedField, obj *db.OrderModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Order",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNDate2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Order_ordered_products(ctx context.Context, field graphql.CollectedField, obj *db.OrderModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Order",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OrderedProduct, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]db.OrderedProductModel)
	fc.Result = res
	return ec.marshalNOrdered_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProductᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) _Ordered_Product_id(ctx context.Context, field graphql.CollectedField, obj *db.OrderedProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Ordered_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Ordered_Product_sub_product(ctx context.Context, field graphql.CollectedField, obj *db.OrderedProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Ordered_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubProduct, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*db.SubProductModel)
	fc.Result = res
	return ec.marshalNSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Product_id(ctx context.Context, field graphql.CollectedField, obj *db.ProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Product_name(ctx context.Context, field graphql.CollectedField, obj *db.ProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Product_description(ctx context.Context, field graphql.CollectedField, obj *db.ProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Product_image_url(ctx context.Context, field graphql.CollectedField, obj *db.ProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ImageURL, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Product_sub_products(ctx context.Context, field graphql.CollectedField, obj *db.ProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubProduct, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProductᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Categories(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Categories(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.CategoryModel)
	fc.Result = res
	return ec.marshalOCategory2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Category(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Category_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Category(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CategoryModel)
	fc.Result = res
	return ec.marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Customers(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Customers(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.CustomerModel)
	fc.Result = res
	return ec.marshalOCustomer2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Customer(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Customer_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Customer(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.CustomerModel)
	fc.Result = res
	return ec.marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_AllProducts(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().AllProducts(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Products(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Products_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Products(rctx, args["category_id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Product(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Product_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Product(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.ProductModel)
	fc.Result = res
	return ec.marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_SubProducts(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_SubProducts_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().SubProducts(rctx, args["product_id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_SubProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_SubProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().SubProduct(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.SubProductModel)
	fc.Result = res
	return ec.marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Orders(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Orders_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Orders(rctx, args["customer_id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Order(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Order_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Order(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderModel)
	fc.Result = res
	return ec.marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_OrderedProducts(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_OrderedProducts_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().OrderedProducts(rctx, args["order_id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.OrderedProductModel)
	fc.Result = res
	return ec.marshalOOrdered_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_OrderedProduct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_OrderedProduct_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().OrderedProduct(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.OrderedProductModel)
	fc.Result = res
	return ec.marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Admins(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Admins(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]db.AdminModel)
	fc.Result = res
	return ec.marshalOAdmin2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_Admin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_Admin_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Admin(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*db.AdminModel)
	fc.Result = res
	return ec.marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _Sub_Product_id(ctx context.Context, field graphql.CollectedField, obj *db.SubProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Sub_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Sub_Product_price(ctx context.Context, field graphql.CollectedField, obj *db.SubProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Sub_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Price, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Sub_Product_size(ctx context.Context, field graphql.CollectedField, obj *db.SubProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Sub_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Size, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Sub_Product_color(ctx context.Context, field graphql.CollectedField, obj *db.SubProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Sub_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Color, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Sub_Product_amount(ctx context.Context, field graphql.CollectedField, obj *db.SubProductModel) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Sub_Product",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Amount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalN__DirectiveLocation2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	fc.Result = res
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirectiveᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	fc.Result = res
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	fc.Result = res
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var adminImplementors = []string{"Admin"}

func (ec *executionContext) _Admin(ctx context.Context, sel ast.SelectionSet, obj *db.AdminModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, adminImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Admin")
		case "id":
			out.Values[i] = ec._Admin_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "login":
			out.Values[i] = ec._Admin_login(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "password":
			out.Values[i] = ec._Admin_password(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var categoryImplementors = []string{"Category"}

func (ec *executionContext) _Category(ctx context.Context, sel ast.SelectionSet, obj *db.CategoryModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, categoryImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Category")
		case "id":
			out.Values[i] = ec._Category_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec._Category_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "image_url":
			out.Values[i] = ec._Category_image_url(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "products":
			out.Values[i] = ec._Category_products(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var customerImplementors = []string{"Customer"}

func (ec *executionContext) _Customer(ctx context.Context, sel ast.SelectionSet, obj *db.CustomerModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, customerImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Customer")
		case "id":
			out.Values[i] = ec._Customer_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec._Customer_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "email":
			out.Values[i] = ec._Customer_email(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "phone":
			out.Values[i] = ec._Customer_phone(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "address":
			out.Values[i] = ec._Customer_address(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "region":
			out.Values[i] = ec._Customer_region(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "cc_number":
			out.Values[i] = ec._Customer_cc_number(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "orders":
			out.Values[i] = ec._Customer_orders(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createCategory":
			out.Values[i] = ec._Mutation_createCategory(ctx, field)
		case "editCategory":
			out.Values[i] = ec._Mutation_editCategory(ctx, field)
		case "deleteCategory":
			out.Values[i] = ec._Mutation_deleteCategory(ctx, field)
		case "addProduct":
			out.Values[i] = ec._Mutation_addProduct(ctx, field)
		case "editProduct":
			out.Values[i] = ec._Mutation_editProduct(ctx, field)
		case "deleteProduct":
			out.Values[i] = ec._Mutation_deleteProduct(ctx, field)
		case "createSubProduct":
			out.Values[i] = ec._Mutation_createSubProduct(ctx, field)
		case "editSubProduct":
			out.Values[i] = ec._Mutation_editSubProduct(ctx, field)
		case "deleteSubProduct":
			out.Values[i] = ec._Mutation_deleteSubProduct(ctx, field)
		case "createCustomer":
			out.Values[i] = ec._Mutation_createCustomer(ctx, field)
		case "editCustomer":
			out.Values[i] = ec._Mutation_editCustomer(ctx, field)
		case "deleteCustomer":
			out.Values[i] = ec._Mutation_deleteCustomer(ctx, field)
		case "addOrder":
			out.Values[i] = ec._Mutation_addOrder(ctx, field)
		case "editOrder":
			out.Values[i] = ec._Mutation_editOrder(ctx, field)
		case "removeOrder":
			out.Values[i] = ec._Mutation_removeOrder(ctx, field)
		case "addOrderedProduct":
			out.Values[i] = ec._Mutation_addOrderedProduct(ctx, field)
		case "editOrderedProduct":
			out.Values[i] = ec._Mutation_editOrderedProduct(ctx, field)
		case "removeOrderedProduct":
			out.Values[i] = ec._Mutation_removeOrderedProduct(ctx, field)
		case "createAdmin":
			out.Values[i] = ec._Mutation_createAdmin(ctx, field)
		case "editAdmin":
			out.Values[i] = ec._Mutation_editAdmin(ctx, field)
		case "deleteAdmin":
			out.Values[i] = ec._Mutation_deleteAdmin(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var orderImplementors = []string{"Order"}

func (ec *executionContext) _Order(ctx context.Context, sel ast.SelectionSet, obj *db.OrderModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, orderImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Order")
		case "id":
			out.Values[i] = ec._Order_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "amount":
			out.Values[i] = ec._Order_amount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "created_at":
			out.Values[i] = ec._Order_created_at(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "ordered_products":
			out.Values[i] = ec._Order_ordered_products(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var ordered_ProductImplementors = []string{"Ordered_Product"}

func (ec *executionContext) _Ordered_Product(ctx context.Context, sel ast.SelectionSet, obj *db.OrderedProductModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, ordered_ProductImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Ordered_Product")
		case "id":
			out.Values[i] = ec._Ordered_Product_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "sub_product":
			out.Values[i] = ec._Ordered_Product_sub_product(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var productImplementors = []string{"Product"}

func (ec *executionContext) _Product(ctx context.Context, sel ast.SelectionSet, obj *db.ProductModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, productImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Product")
		case "id":
			out.Values[i] = ec._Product_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec._Product_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec._Product_description(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "image_url":
			out.Values[i] = ec._Product_image_url(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "sub_products":
			out.Values[i] = ec._Product_sub_products(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "Categories":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Categories(ctx, field)
				return res
			})
		case "Category":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Category(ctx, field)
				return res
			})
		case "Customers":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Customers(ctx, field)
				return res
			})
		case "Customer":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Customer(ctx, field)
				return res
			})
		case "AllProducts":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_AllProducts(ctx, field)
				return res
			})
		case "Products":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Products(ctx, field)
				return res
			})
		case "Product":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Product(ctx, field)
				return res
			})
		case "SubProducts":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_SubProducts(ctx, field)
				return res
			})
		case "SubProduct":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_SubProduct(ctx, field)
				return res
			})
		case "Orders":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Orders(ctx, field)
				return res
			})
		case "Order":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Order(ctx, field)
				return res
			})
		case "OrderedProducts":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_OrderedProducts(ctx, field)
				return res
			})
		case "OrderedProduct":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_OrderedProduct(ctx, field)
				return res
			})
		case "Admins":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Admins(ctx, field)
				return res
			})
		case "Admin":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_Admin(ctx, field)
				return res
			})
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var sub_ProductImplementors = []string{"Sub_Product"}

func (ec *executionContext) _Sub_Product(ctx context.Context, sel ast.SelectionSet, obj *db.SubProductModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, sub_ProductImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Sub_Product")
		case "id":
			out.Values[i] = ec._Sub_Product_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "price":
			out.Values[i] = ec._Sub_Product_price(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "size":
			out.Values[i] = ec._Sub_Product_size(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "color":
			out.Values[i] = ec._Sub_Product_color(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "amount":
			out.Values[i] = ec._Sub_Product_amount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	res, err := graphql.UnmarshalBoolean(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNDate2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDate2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNFloat2float64(ctx context.Context, v interface{}) (float64, error) {
	res, err := graphql.UnmarshalFloat(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNFloat2float64(ctx context.Context, sel ast.SelectionSet, v float64) graphql.Marshaler {
	res := graphql.MarshalFloat(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalID(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalID(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNInt2int(ctx context.Context, v interface{}) (int, error) {
	res, err := graphql.UnmarshalInt(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalInt(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx context.Context, sel ast.SelectionSet, v *db.OrderModel) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Order(ctx, sel, v)
}

func (ec *executionContext) marshalNOrdered_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProductᚄ(ctx context.Context, sel ast.SelectionSet, v []db.OrderedProductModel) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalNOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx context.Context, sel ast.SelectionSet, v *db.OrderedProductModel) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Ordered_Product(ctx, sel, v)
}

func (ec *executionContext) marshalNProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v *db.ProductModel) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Product(ctx, sel, v)
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx context.Context, sel ast.SelectionSet, v *db.SubProductModel) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Sub_Product(ctx, sel, v)
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirectiveᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstringᚄ(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstringᚄ(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalOAdmin2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx context.Context, sel ast.SelectionSet, v []db.AdminModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOAdmin2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐAdmin(ctx context.Context, sel ast.SelectionSet, v *db.AdminModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Admin(ctx, sel, v)
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	res, err := graphql.UnmarshalBoolean(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) unmarshalOBoolean2ᚖbool(ctx context.Context, v interface{}) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalBoolean(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOBoolean2ᚖbool(ctx context.Context, sel ast.SelectionSet, v *bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalBoolean(*v)
}

func (ec *executionContext) marshalOCategory2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx context.Context, sel ast.SelectionSet, v []db.CategoryModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOCategory2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCategory(ctx context.Context, sel ast.SelectionSet, v *db.CategoryModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Category(ctx, sel, v)
}

func (ec *executionContext) marshalOCustomer2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx context.Context, sel ast.SelectionSet, v []db.CustomerModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOCustomer2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐCustomer(ctx context.Context, sel ast.SelectionSet, v *db.CustomerModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Customer(ctx, sel, v)
}

func (ec *executionContext) unmarshalODate2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalString(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalODate2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*v)
}

func (ec *executionContext) unmarshalOFloat2ᚖfloat64(ctx context.Context, v interface{}) (*float64, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalFloat(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOFloat2ᚖfloat64(ctx context.Context, sel ast.SelectionSet, v *float64) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalFloat(*v)
}

func (ec *executionContext) unmarshalOID2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalID(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOID2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalID(*v)
}

func (ec *executionContext) unmarshalOInt2ᚖint(ctx context.Context, v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalInt(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOInt2ᚖint(ctx context.Context, sel ast.SelectionSet, v *int) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalInt(*v)
}

func (ec *executionContext) marshalOOrder2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx context.Context, sel ast.SelectionSet, v []db.OrderModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOOrder2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderᚄ(ctx context.Context, sel ast.SelectionSet, v []db.OrderModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOOrder2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrder(ctx context.Context, sel ast.SelectionSet, v *db.OrderModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Order(ctx, sel, v)
}

func (ec *executionContext) marshalOOrdered_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx context.Context, sel ast.SelectionSet, v []db.OrderedProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOOrdered_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐOrderedProduct(ctx context.Context, sel ast.SelectionSet, v *db.OrderedProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Ordered_Product(ctx, sel, v)
}

func (ec *executionContext) marshalOProduct2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v []db.ProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOProduct2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐProductᚄ(ctx context.Context, sel ast.SelectionSet, v []db.ProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOProduct2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v *db.ProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Product(ctx, sel, v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalString(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*v)
}

func (ec *executionContext) marshalOSub_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx context.Context, sel ast.SelectionSet, v []db.SubProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOSub_Product2ᚕᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProductᚄ(ctx context.Context, sel ast.SelectionSet, v []db.SubProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx, sel, &v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOSub_Product2ᚖgraphqlᚑserverᚋgraphᚋmodelᚐSubProduct(ctx context.Context, sel ast.SelectionSet, v *db.SubProductModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Sub_Product(ctx, sel, v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
