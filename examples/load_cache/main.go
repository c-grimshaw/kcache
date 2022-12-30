package main

import (
	"flag"
	"log"
	"time"

	kv "golang/kv-store"
)

func main() {
	filepath := flag.String("filepath", "", "path to video ids")
	flag.Parse()

	start := time.Now()
	cache, err := kv.LoadCacheFromFile(*filepath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loaded cache in %v\n", time.Since(start))

	key := "---19Dh5uWZ"
	start = time.Now()
	log.Printf("Found key=%s in %v, %v\n", key, time.Since(start), cache.In(key))
}
