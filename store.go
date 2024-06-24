package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))

	hashStr := hex.EncodeToString(hash[:])

	blocksize :=  5
	sliceLen := len(hashStr) / blocksize
}

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.PathTransformFunc(key)

	if err := os.Mkdir(pathName, os.ModePerm); err != nil {
		return err
	}

	fileName := "somefilename"
	pathAndFileName := pathName + "/" + fileName

	f, err := os.Create(pathAndFileName)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)

	if err != nil {
		return err
	}

	log.Printf("writtend %d bytes to disk: %s", n, pathAndFileName)

	return nil
}
