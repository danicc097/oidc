package main

import (
	"context"
	"log"
	"net/http"

	"github.com/zitadel/oidc/example/server/exampleop"
	"github.com/zitadel/oidc/example/server/storage"
)

func main() {
	ctx := context.Background()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore())

	port := "10001"
	// TODO CHANGE ISSUER as needed (localhost needs :port)
	// issuer := "http://localhost:" + port
	issuer := "https://oidcserver.dev.localhost/oidc" // ssl certificate error
	router := exampleop.SetupServer(ctx, issuer, storage)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	// curl certificate error if runnning locally while bookstack is in container.
	// err := server.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")

	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
