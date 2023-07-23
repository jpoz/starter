package env

import (
	"context"
	"fmt"

	"github.com/jpoz/starter/pkg/constants"
	"github.com/jpoz/starter/pkg/query"
)

func QueryCtx(ctx context.Context, obj *query.Query) context.Context {
	return context.WithValue(ctx, constants.QueryContextKey, obj)
}

func Query(ctx context.Context) *query.Query {
	obj, ok := ctx.Value(constants.QueryContextKey).(*query.Query)
	if !ok {
		panic(fmt.Errorf("query not found in context"))
	}

	return obj
}
