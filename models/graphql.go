package models

import "github.com/graphql-go/graphql"

// Object Type

var HeroType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Hero",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"friends": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"characteristics": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: StatusEnumType,
			},
		},
	},
)

var SuperHeroType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Superhero",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: StatusEnumType,
			},
		},
	},
)

// Enum

var enumValues = graphql.EnumValueConfigMap{
	"ACTIVE": &graphql.EnumValueConfig{
		Value: "ACTIVE",
	},
	"INACTIVE": &graphql.EnumValueConfig{
		Value: "INACTIVE",
	},
}

var StatusEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Status",
	Description: "Status of a resource",
	Values:      enumValues,
})

// Union

var CharacterUnionType = graphql.NewUnion(graphql.UnionConfig{
	Name:  "Character",
	Types: []*graphql.Object{HeroType, SuperHeroType},
	ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
		switch p.Value.(type) {
		case *Hero:
			return HeroType
		case *SuperHero:
			return SuperHeroType
		default:
			return nil
		}
	},
})
