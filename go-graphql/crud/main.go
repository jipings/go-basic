package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Info string `json:"info,omitempty"`
	Price float64 `json:"price"`
}

var products = []Product{
	{
		ID:    1,
		Name:  "Chicha Morada",
		Info:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
		Price: 7.99,
	},
	{
		ID:    2,
		Name:  "Chicha de jora",
		Info:  "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
		Price: 5.95,
	},
	{
		ID:    3,
		Name:  "Pisco",
		Info:  "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
		Price: 9.95,
	},
}

var productType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Product",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"info": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:"Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type: productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) ( interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// find product
						for _, product := range products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) product list
			http://localhost:8888/product?query={list{id, name, info, price}}
			*/
			"list": &graphql.Field{
					Type: graphql.NewList(productType),
					Description: "Get product list",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return products, nil
					},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new product item
			http:// localhost:8888/product?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		*/
		"create": &graphql.Field{
			Type: productType,
			Description: "Create new Product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				product := Product{
					ID:    int64(rand.Intn(100000)),
					Name:  p.Args["name"].(string),
					Info:  p.Args["info"].(string),
					Price: p.Args["price"].(float64),
				}
				products = append(products, product)
				return product, nil
			},
		},
		"update": &graphql.Field{
			Type: productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String) },
				"name": &graphql.ArgumentConfig{Type:graphql.String},
				"info": &graphql.ArgumentConfig{Type: graphql.String},
				"price": &graphql.ArgumentConfig{Type:graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) ( interface{},  error) {
				id, _ := p.Args["id"].(int)
				name, nameOk := p.Args["name"].(string)
				info, infoOk := p.Args["info"].(string)
				price, priceOk := p.Args["price"].(float64)
				product := Product{}

				for i, p := range products {
					if int64(id) == p.ID {
						if nameOk {
							products[i].Name = name
						}
						if infoOk {
							products[i].Info = info
						}
						if priceOk {
							products[i].Price = price
						}
						product = products[i]
						break
					}
				}
				return product, nil
			},
		},
		/* Delete product by id
		http://localhost:8888/product?query=mutation+_{delete(id:1){id, name, info, price}}
		*/
		"delete": &graphql.Field{
			Type: productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				product := Product{}
				for i, p := range products {
					if int64(id) == p.ID {
						product = products[i]
						// Remove from product list
						products = append(products[:i], products[i+1:]...)
					}
				}
				return product, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
			Mutation: mutationType,
		})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("server is running on port 8888")
	http.ListenAndServe(":8888", nil)
}