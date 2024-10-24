package stormiclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	reconfig "github.com/stormi-li/Reconfig"
	researd "github.com/stormi-li/Researd"
	resync "github.com/stormi-li/Resync"
	ripc "github.com/stormi-li/Ripc"
)

type Client struct {
	redisClient    *redis.Client
	RipcClient     *ripc.Client
	ReconfigClient *reconfig.Client
	ResyncClient   *resync.Client
	ReseardClient  *researd.Client
}

func NewClient(username, password string) *Client {
	addr, password, err := login(username, password)
	if err != nil {
		return nil
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	redisClient.Set(context.Background(), "fff", "ff", 0)
	c := Client{redisClient: redisClient,
		RipcClient:     ripc.NewClient(redisClient),
		ReconfigClient: reconfig.NewClient(redisClient),
		ResyncClient:   resync.NewClient(redisClient),
		ReseardClient:  researd.NewClient(redisClient),
	}
	c.RipcClient.SetNamespace(username)
	c.ReconfigClient.SetNamespace(username)
	c.ReseardClient.SetNamespace(username)
	c.ResyncClient.SetNamespace(username)
	return &c
}

func Register(username, password string) error {
	return register(username, password)
}
