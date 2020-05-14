package main

import (
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

// Account struct
type Account struct {
	Name    string
	Balance uint64
}

func main() {

	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://wallet-e8234.firebaseio.com",
	}
	opt := option.WithCredentialsFile("config/wallet-e8234-firebase-adminsdk-3tbxy-c03ee23a4c.json")

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}

	writeAcc := Account{
		Name:    "Akagi",
		Balance: 100,
	}

	if err := client.NewRef("accounts/akagi").Set(ctx, &writeAcc); err != nil {
		log.Panic(err)
	}

	var readAcc Account
	if err := client.NewRef("accounts/akagi").Get(ctx, &readAcc); err != nil {
		log.Panic(err)
	}
	log.Printf("%s has a balance of %d\n", readAcc.Name, readAcc.Balance)
}
