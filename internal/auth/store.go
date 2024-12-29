package auth

// import (
// 	"github.com/gorilla/sessions"
// 	redis_db "github.com/harsh082ip/CloudNavX/internal/db/redis"
// 	// "github.com/markbates/sessions"
// )

// // RedisStoreWrapper is a wrapper around the RedisStore to implement sessions.Store.
// type RedisStoreWrapper struct {
// 	redisStore *redis_db.RedisStore
// }

// // NewRedisStoreWrapper creates a new instance of RedisStoreWrapper.
// func NewRedisStoreWrapper() *RedisStoreWrapper {
// 	return &RedisStoreWrapper{
// 		redisStore: redis_db.NewRedisStore(),
// 	}
// }

// // Get retrieves session data from Redis.
// func (r *RedisStoreWrapper) Get(sid string) (string, error) {
// 	return r.redisStore.Get(sid)
// }

// // Set stores session data in Redis.
// func (r *RedisStoreWrapper) Set(sid, value string) error {
// 	return r.redisStore.Set(sid, value)
// }

// // Delete removes session data from Redis.
// func (r *RedisStoreWrapper) Delete(sid string) error {
// 	return r.redisStore.Delete(sid)
// }

// // Example: Using sessions package
// func (r *RedisStoreWrapper) GetSession(sid string) (*sessions.Session, error) {
// 	// Logic to retrieve session from Redis and return a sessions.Session object
// 	// You could create a session wrapper if necessary
// }
