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
	"github.com/rs/cors"
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

	initCache()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/refresh", refresh)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: client}}))))

	log.Printf("Connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost/")
	if err != nil {
		panic(err)
	}

	cache = conn
}

type credentials struct {
	Password string `json:"password"`
	Login    string `json:"login"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body))).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	admin, err := client.Admin.FindFirst(
		db.Admin.Login.Equals(creds.Login),
	).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !hash.CheckPasswordHash(creds.Password, admin.InnerAdmin.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	_, err = cache.Do("SETEX", sessionToken, "900", creds.Login)
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

func welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		fmt.Println(sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Здарова, %s!", response)))
}

func refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	_, err = cache.Do("SETEX", newSessionToken, "900", fmt.Sprintf("%s", response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = cache.Do("DEL", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(15 * time.Minute),
	})
}
