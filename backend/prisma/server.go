package main

import (
	"graphql-server/graph"
	"graphql-server/graph/generated"
	"graphql-server/middleware"
	"graphql-server/prisma/db"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const port = "8080"

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router := chi.NewRouter()

	router.Use(middleware.Middleware(*client))
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: client}})))
	log.Printf("Connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
