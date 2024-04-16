package ctx

import (
	"context"
	"time"
)

func Ctx() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return ctx
}
