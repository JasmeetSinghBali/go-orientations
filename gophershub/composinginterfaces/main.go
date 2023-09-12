package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello gophers!")
	incorrecthashAndBroadcast(bytes.NewReader(payload))
	hashAndBroadcast(NewHashReader(payload))
}

func incorrecthashAndBroadcast(r io.Reader) error {
	// read everything from reader in bytes
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// hash the read bytes
	hash := sha1.Sum(b)

	// print the hash in encoded string format
	fmt.Println(hex.EncodeToString(hash[:]))
	// Wrong📌 pipe/pass-on the reader broadcast instance to broadcast function
	// as the reader was already read, hence incorrectbroadcast has nothing to read
	return incorrectbroadcast(r)
}

// 💫
// composed
// here the reader is of type HashReader that is interface with embedded reader & hash
// so to test this hashAndBroadcast functionality both reader and hash is required
func hashAndBroadcast(r HashReader) error {
	// 💫
	// here r.hash/Read has both functionality
	hash := r.hash()
	// print the hash that is already in string format
	fmt.Println(hash)

	return broadcast(r)
}

func incorrectbroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes: ", string(b))

	return nil
}

// 💫
// composed
// here the reader is of type io.Reader and its only task is to read not hash it
// so to test this broadcast functionality only a Reader is required
func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes: ", string(b))

	return nil
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

// 💡
// composable hash reader struct with embedded byte reader and buffer
func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// 💫
// composite hashReader that reads bytes and a hash function to hash those bytes embedded in a single interface named HashReader
// now this interface can have this hash function implementation in sha1,sha256,hmac....
type HashReader interface {
	io.Reader
	hash() string
}
