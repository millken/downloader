package core

import (
	"bytes"
	"sync"
)

var (
	// BufferPool1k is the buffer pool with 1K initialized capacity.
	BufferPool1k = NewBufferPool(1024)

	// BufferPool2k is the buffer pool with 2K initialized capacity.
	BufferPool2k = NewBufferPool(2048)

	// BufferPool4k is the buffer pool with 4K initialized capacity.
	BufferPool4k = NewBufferPool(4096)

	// BufferPool8k is the buffer pool with 8K initialized capacity.
	BufferPool8k = NewBufferPool(8192)
)

// BufferPool is the bytes.Buffer wrapper of sync.Pool.
type BufferPool struct {
	pool sync.Pool
}

func makeBuffer(size int) (b *bytes.Buffer) {
	b = bytes.NewBuffer(make([]byte, size))
	b.Reset()
	return
}

// NewBufferPool returns a new bytes.Buffer pool.
func NewBufferPool(size int) *BufferPool {
	newf := func() interface{} { return makeBuffer(size) }
	return &BufferPool{pool: sync.Pool{New: newf}}
}

// Get returns a bytes.Buffer.
func (p *BufferPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

// Put places a bytes.Buffer to the pool.
func (p *BufferPool) Put(b *bytes.Buffer) {
	if b != nil {
		b.Reset()
		p.pool.Put(b)
	}
}
