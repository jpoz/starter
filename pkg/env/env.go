package env

import (
	"context"

	"github.com/jpoz/starter/pkg/query"
	"github.com/redis/go-redis/v9"
)

func Attach(ctx context.Context,
	query *query.Query,
	redis *redis.Client,
) context.Context {
	ctx = QueryCtx(ctx, query)
	ctx = RedisCtx(ctx, redis)

	return ctx
}
