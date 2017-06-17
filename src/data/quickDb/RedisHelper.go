package quickDb

import (
	"log"
	"github.com/gosexy/redis"
)

func init() {
	var client *redis.Client
	var err error

	client = redis.New()

	err = client.Connect("localhost", "8082")

	if err != nil {
		log.Fatalf("Connect failed: %s\n", err.Error())
		return
	}

	log.Println("Connected to redis-server.")

	// DEL hello
	log.Printf("DEL hello\n")
	client.Del("hello")

	// SET hello 1
	log.Printf("SET hello 1\n")
	client.Set("hello", 1)

	// INCR hello
	log.Printf("INCR hello\n")
	client.Incr("hello")

	// GET hello
	log.Printf("GET hello\n")
	s, err := client.Get("hello")

	if err != nil {
		log.Fatalf("Could not GET: %s\n", err.Error())
		return
	}

	log.Printf("> hello = %s\n", s)

	client.Quit()
}
