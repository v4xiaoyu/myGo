package quickDb

import (
	"github.com/gosexy/redis"
	"log"
)

var client *redis.Client

func Init() {
	var err error

	client = redis.New()

	err = client.Connect("localhost", 6379)

	if err != nil {
		log.Fatalf("Connect failed: %s\n", err.Error())
		return
	}

	log.Println("Connected to redis-server.")
}

func Get(key string) string {
	s, err := client.Get(key)
	if err != nil {
		return err.Error()
	}
	return s
}

func Delete(key string) {
	client.Del(key)
}

func Set(key string, v interface{}) {
	client.Set(key, v)
}

func Quit() {
	client.Quit()
}

func Increase(key string) {
	client.Incr(key)
}

func Decrease(key string) {
	client.Decr(key)
}
