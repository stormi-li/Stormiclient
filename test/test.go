package main

import (
	"fmt"

	stormiclient "github.com/stormi-li/Stormiclient"
)

func main() {
	c := stormiclient.NewClient("administor", "123456654321")
	cfg := c.ReconfigClient.GetConfig("redis")
	fmt.Println(cfg)
}
