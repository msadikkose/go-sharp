package tests

import (
	"context"
	"go-sharp/app"
	"go-sharp/ent"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const testDbDriver = "sqlite3"
const testDbConnectionString = "file:ent?mode=memory&cache=shared&_fk=1"

type TestDbHandler struct {

}

func (d *TestDbHandler) GetDbClient() (*ent.Client) {
	client, err := ent.Open(testDbDriver, testDbConnectionString)
	if err != nil {
		log.Fatalf(
			"failed opening connection to %v: %v",
			app.Config.DbDriver,
			err,
		)
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	return client
}

func  (d *TestDbHandler) Migrate() {
	client, err := ent.Open(testDbDriver, testDbConnectionString)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	ctx:=context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	client.User.
		Create().
		SetName("Ali").
		SetAge(30).
		SetIsDeleted(false).
		Save(ctx)

}
