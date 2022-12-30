package kv

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

const (
	// Define the population of characters to generate IDs from
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numerals = "0123456789"
	special  = "-_"

	charset = alphabet + numerals + special
)

// createRandomID creates a unique ID of length k.
func createRandomID(k int) []byte {
	id, n := make([]byte, k), len(charset)
	for i := range id {
		id[i] = charset[rand.Intn(n)]
	}
	return id
}

// createRandomData generates numRows of k-length IDs
// writing the results to a file.
func createRandomData(filename string, numRows, k int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i < numRows; i++ {
		w.Write(append(createRandomID(k), '\n'))
	}
}
