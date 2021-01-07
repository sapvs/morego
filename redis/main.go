package main

import "github.com/sudosapan/redis/client"

func main() {
	client.Set("key", "value")

	client.Get("key")
}
