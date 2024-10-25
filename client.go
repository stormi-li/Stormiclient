package stormiclient

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	reconfig "github.com/stormi-li/Reconfig"
	researd "github.com/stormi-li/Researd"
	resync "github.com/stormi-li/Resync"
	ripc "github.com/stormi-li/Ripc"
)

type Client struct {
	redisClient    *redis.Client
	ripcClient     *ripc.Client
	reconfigClient *reconfig.Client
	resyncClient   *resync.Client
	researdClient  *researd.Client
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
		ripcClient:     ripc.NewClient(redisClient),
		reconfigClient: reconfig.NewClient(redisClient),
		resyncClient:   resync.NewClient(redisClient),
		researdClient:  researd.NewClient(redisClient),
	}
	c.ripcClient.SetNamespace(username)
	c.reconfigClient.SetNamespace(username)
	c.researdClient.SetNamespace(username)
	c.resyncClient.SetNamespace(username)
	return &c
}

func (c *Client) Notify(channel string, msg string) {
	c.ripcClient.Notify(channel, msg)
}

func (c *Client) Wait(channel string, ttl time.Duration) string {
	return c.ripcClient.Wait(channel, ttl)
}

func (c *Client) NewListener(channel string) *ripc.Listener {
	return c.ripcClient.NewListener(channel)
}

func (c *Client) NewLock(lockName string) *resync.Lock {
	return c.resyncClient.NewLock(lockName)
}

func (c *Client) GetConfig(name string) *reconfig.ConfigInfo {
	return c.reconfigClient.GetConfig(name)
}

func (c *Client) GetConfigNames() []string {
	return c.reconfigClient.GetConfigNames()
}

func (c *Client) ListenConfig(name string, handler func(config *reconfig.ConfigInfo)) {
	c.reconfigClient.Connect(name, handler)
}

func (c *Client) NewConfig(name string, addr string) *reconfig.Config {
	return c.reconfigClient.NewConfig(name, addr)
}

func (c *Client) Register(name, addr string, weight int) {
	c.researdClient.Register(name, addr, weight)
}

func (c *Client) Discover(name string) {
	c.researdClient.Discover(name)
}

func Register(username, password string) error {
	return register(username, password)
}
