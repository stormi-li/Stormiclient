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
	c := Client{redisClient: redisClient,
		RipcClient:     ripc.NewClient(redisClient, username),
		ReconfigClient: reconfig.NewClient(redisClient, username),
		ResyncClient:   resync.NewClient(redisClient, username),
		ReseardClient:  researd.NewClient(redisClient, username),
	}
	return &c
}

func Register(username, password string) error {
	return register(username, password)
}
