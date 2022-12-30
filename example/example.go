package main

import (
	"log"
	"time"

	kv "golang/kv-store"
)

const filepath = "../video_ids_sorted.txt"

func main() {
	start := time.Now()
	cache, err := kv.LoadCacheFromFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loaded cache in %v\n", time.Since(start))

	start = time.Now()
	in := cache.In("---19Dh5uWZ")
	log.Printf("Found key in %v, %v\n", time.Since(start), in)
}
