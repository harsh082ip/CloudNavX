package redis_db

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/harsh082ip/CloudNavX/internal/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

// RedisStore implements the sessions.Store interface for session management using Redis.
type RedisStore struct {
	client    *redis.Client
	keyPrefix string
	maxAge    time.Duration
}

// InitializeRedis initializes the Redis client using the URI from the config.
func InitializeRedis() {
	redisURI := config.AppConfig.RedisURI
	if redisURI == "" {
		log.Fatal("REDIS_URI cannot be empty!")
	}

	options, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Fatal("Error parsing REDIS_URI:", err)
	}

	RedisClient = redis.NewClient(options)

	// Test Redis connection
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}
	log.Println("Connected to Redis successfully!")
}

// NewRedisStore creates a new RedisStore instance with default options.
func NewRedisStore() *RedisStore {
	return &RedisStore{
		client:    RedisClient,
		keyPrefix: "session_", // Prefix for Redis keys to avoid collisions
		maxAge:    time.Hour,  // Default TTL of 1 hour
	}
}

// Get fetches a session by its name from Redis.
func (r *RedisStore) Get(req *http.Request, name string) (*sessions.Session, error) {
	key := r.keyPrefix + name

	// Retrieve session data from Redis
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		// If no session exists, return an empty session
		return sessions.NewSession(r, name), nil
	} else if err != nil {
		return nil, err
	}

	// Deserialize session data
	session := sessions.NewSession(r, name)
	buffer := bytes.NewBufferString(data)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(&session.Values); err != nil {
		return nil, err
	}

	return session, nil
}

// New creates and initializes a new session.
func (r *RedisStore) New(req *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(r, name)
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: int(r.maxAge.Seconds()),
	}
	return session, nil
}

// Save persists session data in Redis.
func (r *RedisStore) Save(req *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	if session.IsNew {
		session.IsNew = false
	}

	key := r.keyPrefix + session.Name()
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(session.Values); err != nil {
		return err
	}

	// Store serialized session data in Redis with TTL
	err := r.client.Set(ctx, key, buffer.String(), r.maxAge).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes session data from Redis.
func (r *RedisStore) Delete(req *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	key := r.keyPrefix + session.Name()

	// Remove the session data from Redis
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
