package main

import (
	"fmt"

	stormiclient "github.com/stormi-li/Stormiclient"
)

func main() {
	stormiclient.Register("your username", "yourpassword")
	c := stormiclient.NewClient("your username", "yourpassword")
	cfg := c.ReconfigClient.NewConfig("redis", "localhost:6379")
	cfg.Upload()
	cfg = c.ReconfigClient.GetConfig("redis")
	fmt.Println(cfg)
}
