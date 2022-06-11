package main

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	size   int
	oldest int
	newest int
	data   []byte
}

func NewBuffer(size int) *Buffer {
	buffer := Buffer{size: size, oldest: -1, newest: -1}
	buffer.data = make([]byte, size)
	return &buffer
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.oldest >= 0 {
		ret := b.data[b.oldest]
		if b.oldest == b.newest {
			b.Reset()
		} else {
			b.oldest++
			if b.oldest == b.size {
				b.oldest = 0
			}
		}
		return ret, nil

	} else {
		return byte(0), errors.New("empty buffer")
	}
}

func (b *Buffer) WriteByte(c byte) error {
	b.newest++
	if b.newest == 0 {
		b.oldest = 0
		b.data[b.newest] = c
		return nil
	}

	if b.newest == b.size {
		b.newest = 0
	}

	if b.newest != b.oldest {
		b.data[b.newest] = c
		return nil
	} else {
		return errors.New("attempt to write to full buffer")
	}
}

func (b *Buffer) Overwrite(c byte) {
	b.newest++
	if b.newest == 0 {
		b.oldest = 0
		b.data[b.newest] = c
		return
	}

	if b.newest == b.size {
		b.newest = 0
	}

	if b.newest == b.oldest {
		b.oldest++
		if b.oldest == b.size {
			b.oldest = 0
		}
	}
	b.data[b.newest] = c
}

func (b *Buffer) Reset() {
	b.oldest = -1
	b.newest = -1
}
