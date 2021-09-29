package main

import (
	"bytes"
	"encoding/gob"
	"github.com/nate-droid/core/symphony"
)

func encodeSong(song *symphony.Song) ([]byte, error){
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(song)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func decodeSong(song []byte) (*symphony.Song, error) {
	r := bytes.NewReader(song)
	dec := gob.NewDecoder(r)
	var n *symphony.Song
	err := dec.Decode(&n)
	if err != nil {
		return nil, err
	}

	return n, nil
}

