package db

import (
	"context"
	"go-sharp/app"
	"go-sharp/ent"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type IDbHandler interface {
	GetDbClient() *ent.Client
	Migrate()
}

type DbHandler struct {

}

func (d *DbHandler) GetDbClient() (*ent.Client) {
	client, err := ent.Open(app.Config.DbDriver, app.Config.ConnectionString)
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

func  (d *DbHandler) Migrate() {
	client, err := ent.Open(app.Config.DbDriver, app.Config.ConnectionString)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	ctx:=context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}
}