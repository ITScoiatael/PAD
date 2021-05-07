package graph

import "graphql-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	categories []*model.Category
	admins     []*model.Admin
	customers  []*model.Customer
}
