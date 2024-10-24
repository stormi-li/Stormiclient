package main

import (
	"fmt"
	"time"

	stormiclient "github.com/stormi-li/Stormiclient"
)

func main() {
	client := stormiclient.NewClient("administor", "123456654321")
	go func() {
		res := client.RipcClient.Wait("c1", 2*time.Second)
		fmt.Println(res)
	}()
	time.Sleep(100 * time.Millisecond)
	client.RipcClient.Notify("c1", "hh")
	time.Sleep(1 * time.Second)
}
