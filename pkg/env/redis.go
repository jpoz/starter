package env

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/jpoz/starter/pkg/constants"
)

func RedisCtx(ctx context.Context, r *redis.Client) context.Context {
	return context.WithValue(ctx, constants.RedisContextKey, r)
}

func Redis(ctx context.Context) *redis.Client {
	r, ok := ctx.Value(constants.RedisContextKey).(*redis.Client)
	if !ok {
		panic(fmt.Errorf("redis not found in context"))
	}

	return r
}
