package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/iamgoangle/go-graphql-example/cmd/gql/graph"
	"github.com/iamgoangle/go-graphql-example/cmd/gql/graph/generated"
)

var (
	APP_PORT = "8080"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func healthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": http.StatusText(http.StatusOK),
		})
	}
}

func main() {
	r := gin.Default()
	r.GET("/health", healthHandler())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(":" + APP_PORT)
}
