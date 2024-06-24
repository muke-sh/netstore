package main

import ("testing"
"bytes"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpeg"))
	if err := s.writeStream("mypic", data); err != nil {
		t.Error(err)
	}
}
