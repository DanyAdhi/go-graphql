package graph

import (
	"github.com/danyadhi/go-graphql/graph/model"
	"github.com/danyadhi/go-graphql/internal/orders"
	"github.com/danyadhi/go-graphql/internal/pets"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos         []*model.Todo
	PetsService   pets.Service
	OrdersService orders.Service
}
