package main

import (
	"log"
	"os"

	handler "github.com/99designs/gqlgen/handler"
	"github.com/Akshaytermin/gqltest/graph"
	"github.com/Akshaytermin/gqltest/graph/generated"
	"github.com/Akshaytermin/gqltest/graph/model"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)

	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	h := handler.Playground("GraphQL", "/query")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)

	}
}

func main() {

	//Migrate Db
	db := model.FetchConnection()
	db.AutoMigrate(&model.Product{}, &model.Ingredient{})
	db.Close()

	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
