package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"graphql/models"
	"graphql/repos"
	"log"
	"net/http"
)

func GraphQl(c *gin.Context) {
	var req models.ReqQuery

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	queryFields := graphql.Fields{
		"hero": &graphql.Field{
			Type:        models.HeroType,
			Description: "get hero by Id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"status": &graphql.ArgumentConfig{
					Type: models.StatusEnumType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if id, ok := p.Args["id"].(int); ok {
					return repos.GetHeroById(id), nil
				} else if status, ok := p.Args["status"].(string); ok {
					return repos.GetHeroByStatus(status), nil
				}
				return nil, nil
			},
		},
		"heroes": &graphql.Field{
			Type: graphql.NewList(models.HeroType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return repos.GetHeroes(), nil
			},
		},
		"characters": &graphql.Field{
			Type: models.CharacterUnionType,
			Args: graphql.FieldConfigArgument{
				"character_type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if characterType, ok := p.Args["character_type"].(string); ok {
					switch characterType {
					case "hero":
						return repos.GetHeroById(1), nil
					case "superhero":
						return &models.SuperHero{
							Id:     1,
							Name:   "superhero",
							Status: "ACTIVE",
						}, nil
					default:
						return nil, nil
					}
				}
				return nil, nil
			},
		},
	}

	mutationFields := graphql.Fields{
		"updateHero": &graphql.Field{
			Type:        models.HeroType,
			Description: "update hero name by Id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if id, ok := p.Args["id"].(int); ok {
					return repos.UpdateHeroNameById(id, "IM"), nil
				}
				return nil, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutationFields,
	}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("failed to create schema")
	}
	params := graphql.Params{Schema: schema, RequestString: req.Query, OperationName: req.OperationName, VariableValues: req.Variables}

	result := graphql.Do(params)

	c.JSON(http.StatusOK, result)
}
