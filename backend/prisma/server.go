package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"graphql-server/graph"
	"graphql-server/graph/generated"
	"graphql-server/hash"
	"graphql-server/prisma/db"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

const port = "8080"

var cache redis.Conn
var client *db.PrismaClient

func main() {
	client = db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/signin", signin)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: client}})))

	log.Printf("Connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}

	cache = conn
}

type Credentials struct {
	Password string `json:"password"`
	Login    string `json:"login"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	body, err := ioutil.ReadAll(r.Body)
	parsedString := ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println(parsedString)

	err = json.NewDecoder(parsedString).Decode(&creds)
	if err != nil {
		fmt.Println(err)
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	admin, err := client.Admin.FindFirst(
		db.Admin.Login.Equals(creds.Login),
	).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expectedPassword := admin.InnerAdmin.Password

	if hash.CheckPasswordHash(creds.Password, expectedPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	_, err = cache.Do("SETEX", sessionToken, "15", creds.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(15 * time.Minute),
	})
}
