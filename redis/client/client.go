package client

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	client    *redis.Client
	redisAddr string = os.Getenv("REDIS_ADDR")
)

func connect() {
	if checkConnected() {
		return
	}
	log.Printf("Connecting redis at %s\n", redisAddr)
	client = redis.NewClient(&redis.Options{Addr: redisAddr, Password: "", DB: 0})

	if checkConnected() {
		log.Printf("Connected to redis at %s ", redisAddr)
	} else {
		log.Printf("Could not connect to redis at %s ", redisAddr)
	}

}

func checkConnected() bool {
	if client == nil {
		return false
	}

	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("redis ping to %s failed due to %v\n", redisAddr, err.Error())
		return false
	}
	return true
}

func init() {
	connect()
}

//Set sets value to a key, overwrites if exists.
func Set(key string, value string) error {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("Could not set %s to %s due to %v\n", value, key, err.Error())
		return err
	}

	log.Printf("Set %s to %s\n", value, key)
	return nil
}

//Get gets value to a key, return nil if not found.
func Get(key string) (string, error) {
	value := client.Get(key)
	err := value.Err()
	if err != nil {
		log.Printf("Could not get value for %s due to %v\n", key, err.Error())
		return "", err
	}

	log.Printf("Got value %s for %s\n", value.Val(), key)
	return value.Val(), nil
}

//Add appends value to a key, .
func Add(key string, value string) error {
	return nil
}
