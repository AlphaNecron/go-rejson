package rejson

import (
	"github.com/nitishm/go-rejson/v4/clients"
	"github.com/nitishm/go-rejson/v4/rjs"
	"github.com/redis/go-redis/v9"
)

// RedisClient provides interface for Client handling in the ReJSON Handler
type RedisClient interface {
	SetClientInactive()
	SetGoRedisClient(conn redis.UniversalClient)
}

// SetClientInactive resets the handler and unset any client, set to the handler
func (r *Handler) SetClientInactive() {
	_t := &Handler{clientName: rjs.ClientInactive}
	r.clientName = _t.clientName
	r.implementation = _t.implementation
}

// SetGoRedisClient sets Go-Redis (https://github.com/go-redis/redis) client
func (r *Handler) SetGoRedisClient(conn redis.UniversalClient) {
	r.clientName = "goredis"
	r.implementation = clients.NewGoRedisClient(conn)
}
