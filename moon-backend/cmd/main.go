package main

import (
	"moon/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=capstone password=capstone dbname=capstone sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
