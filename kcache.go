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

// KCache represents a set membership object
type KCache map[[PfxLen]byte][][ValLen]byte

// In returns true if the given key is present inside the cache,
// false otherwise.
func (k KCache) In(key string) bool {
	var prefix [PfxLen]byte
	var value [ValLen]byte
	copy(prefix[:], key[:PfxLen])
	copy(value[:], key[PfxLen:])

	prefixCache, ok := k[prefix]
	if !ok {
		return false
	}

	i := sort.Search(len(prefixCache), func(i int) bool {
		return string(prefixCache[i][:]) >= string(value[:])
	})
	return i < len(prefixCache) && prefixCache[i] == value
}

// LoadCache loads a KCache with keys from a readable interface
func LoadCache(r io.Reader) (cache KCache, err error) {
	// Scan IDs line-by-line, divided into prefix/value
	cache = make(KCache)
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		var prefix [PfxLen]byte
		var value [ValLen]byte

		// Split ID into fixed-length byte arrays
		copy(prefix[:], scanner.Bytes()[:PfxLen])
		copy(value[:], scanner.Bytes()[PfxLen:])
		cache[prefix] = append(cache[prefix], value)
	}

	// Sort sub-caches to maintain binary search invariance
	for _, prefixCache := range cache {
		sort.Slice(prefixCache, func(i, j int) bool {
			return string(prefixCache[i][:]) < string(prefixCache[j][:])
		})
	}
	return cache, nil
}

// LoadCacheFromFile is a wrapper for loading keys from a given file
func LoadCacheFromFile(filename string) (KCache, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return LoadCache(f)
}
