package kv

import (
	"bufio"
	"io"
	"os"
	"sort"
)

const (
	// KeyLen is the length of IDs in the cache
	KeyLen = 11
	// PfxLen (Prefix Length) is the length of the prefix key
	PfxLen = 3
	// ValLen is the remainder of the key length
	ValLen = KeyLen - PfxLen
)

// KCache represents a set membership object.
type KCache map[[PfxLen]byte][][ValLen]byte

// In returns true if the given key is present inside the cache,
// false otherwise.
func (k KCache) In(key string) bool {
	prefix, value := getPrefixValue(key)

	prefixCache, ok := k[prefix]
	if !ok {
		return false
	}

	i := sort.Search(len(prefixCache), func(i int) bool {
		return string(prefixCache[i][:]) >= string(value[:])
	})
	return i < len(prefixCache) && prefixCache[i] == value
}

// Add inserts a key into the KCache, disallowing duplicate entries.
func (k KCache) Add(key string) {
	prefix, value := getPrefixValue(key)
	prefixCache := k[prefix]

	// Find insertion index in the sorted array
	i := sort.Search(len(prefixCache), func(i int) bool {
		return string(prefixCache[i][:]) >= string(value[:])
	})

	// Ignore duplicate entries
	if i < len(prefixCache) && prefixCache[i] == value {
		return
	}
	k[prefix] = insertAt(prefixCache, i, value)
}

// LoadCache loads a KCache with keys from a readable interface.
func LoadCache(r io.Reader) (KCache, error) {
	// Scan IDs line-by-line, divided into prefix/value
	scanner, cache := bufio.NewScanner(r), make(KCache)
	for scanner.Scan() {
		prefix, value := getPrefixValue(scanner.Text())
		cache[prefix] = append(cache[prefix], value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Sort prefix caches to maintain binary search invariance
	for _, prefixCache := range cache {
		sort.Slice(prefixCache, func(i, j int) bool {
			return string(prefixCache[i][:]) < string(prefixCache[j][:])
		})
	}
	return cache, nil
}

// LoadCacheFromFile is a wrapper for loading keys from a given file.
func LoadCacheFromFile(filename string) (KCache, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return LoadCache(f)
}

// getPrefixValue splits a key into fixed-length byte arrays.
func getPrefixValue(key string) (pfx [PfxLen]byte, val [ValLen]byte) {
	copy(pfx[:], key[:PfxLen])
	copy(val[:], key[PfxLen:])
	return
}

// insertAt inserts a value into an array at index i.
func insertAt(data [][ValLen]byte, i int, val [ValLen]byte) [][ValLen]byte {
	if i == len(data) {
		return append(data, val)
	}

	data = append(data[:i+1], data[i:]...)
	data[i] = val
	return data
}
