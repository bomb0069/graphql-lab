package graphql

import (
	// "strconv"
	"graphql-api/pkg/data/models"
	"graphql-api/pkg/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

/*
Contact Types
*/
var ContactGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Contact",
	Fields: graphql.Fields{
		"contact_id": &graphql.Field{Type: graphql.Int},
		"name":       &graphql.Field{Type: graphql.String},
		"first_name": &graphql.Field{Type: graphql.String},
		"last_name":  &graphql.Field{Type: graphql.String},
		"gender_id":  &graphql.Field{Type: graphql.Int},
		"dob":        &graphql.Field{Type: graphql.DateTime},
		"email":      &graphql.Field{Type: graphql.String},
		"phone":      &graphql.Field{Type: graphql.String},
		"address":    &graphql.Field{Type: graphql.String},
		"photo_path": &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.DateTime},
		"created_by": &graphql.Field{Type: graphql.String},
		// Add field here
	},
})

var ContactPaginationGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ContactPagination",
	Fields: graphql.Fields{
		"contacts":   &graphql.Field{Type: graphql.NewList(ContactGraphQLType)},
		// Add field here
	},
})

type ContactQueries struct {
	Gets          func(string) ([]*models.ContactModel, error)         `json:"gets"`
}

// Define the ContactQueries type
var ContactQueriesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ContactQueries",
	Fields: graphql.Fields{
		"gets": &graphql.Field{
			Type: graphql.NewList(ContactGraphQLType),
			Args: SearhTextQueryArgument,
			Resolve:// auth.AuthorizeResolverClean("contacts.gets", monitoring.TraceResolver( cache.GetCacheResolver(resolvers.GetContactResolve))),
			resolvers.GetContactResolve,
		},
		"getContactById": &graphql.Field{
			Type: graphql.NewList(ContactGraphQLType),
			Args: SearhTextQueryArgument,
			Resolve:// auth.AuthorizeResolverClean("contacts.gets", monitoring.TraceResolver( cache.GetCacheResolver(resolvers.GetContactResolve))),
			resolvers.GetContactResolveById,
		},
	},
})

