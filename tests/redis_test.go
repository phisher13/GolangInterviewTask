package main

import (
	"context"
	"fmt"
	"testing"
	"urls/pkg/redis"

	"github.com/stretchr/testify/assert"
)


func TestRedisConnection(t *testing.T) {
	ctx := context.Background()
	rdb := redis.InitRedisConnect()
	err := rdb.Set(ctx, "ping", "pong", 0).Err()

	if err != nil {
		t.Errorf("Couldn't connect to the redis. Error: %s", err)
	}
	result, err := rdb.Get(ctx, "ping").Result()

	if err != nil {
		t.Errorf("Couldn't connect to the redis. Error: %s", err)
	}

	t.Logf("Set key: 'ping', value: 'pong'. Expect: 'pong', Result: %s\n", 
				result)

	assert.Equal(t, "pong", result,
		fmt.Sprintf("Incorrect value. Expect: pong, got %s",
		result))
}