package repository

import (
	"context"
	"go-sharp/ent"
)

type Base struct {
	Transaction *ent.Tx
	Ctx         context.Context
}
