package main

import (
	"developer-notes/config"
	"developer-notes/graph"
	"developer-notes/graph/model"
	"developer-notes/routes"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8081"

func main() {
	// Initialize database connection
	config.Connect()

	// Create a new resolver with the event channel
	resolver := &graph.Resolver{
		NoteEvents: make(chan *model.Note, 100),
	}

	// Create GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Configure WebSocket transport
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Create Gin router
	router := gin.New()

	// Add REST API routes under /api prefix
	api := router.Group("/api")
	routes.NotesRoute(api)

	// Add GraphQL routes
	router.GET("/graphql", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
	})
	router.POST("/graphql/query", gin.WrapH(srv))
	router.GET("/graphql/query", gin.WrapH(srv)) // Add GET handler for WebSocket

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("GraphQL playground available at http://localhost:%s/graphql", port)
	log.Printf("REST API available at http://localhost:%s/api/notes", port)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
