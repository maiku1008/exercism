package paasio

import (
	"io"
	"sync"
)

// return a new counter
func newCounter() counter {
	return counter{}
}

// A concurrency-safe counter
type counter struct {
	sync.Mutex
	bytes int64
	ops   int
}

// increments byte and ops count in a concurrency-safe manner
func (counter *counter) addBytes(bytes int64) {
	counter.Lock()
	defer counter.Unlock()
	counter.bytes += bytes
	counter.ops++
}

// gets byte and ops count in a concurrency-safe manner
func (counter *counter) getCount() (bytes int64, ops int) {
	counter.Lock()
	defer counter.Unlock()
	return counter.bytes, counter.ops
}

// NewReadCounter initializes a new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		reader:  r,
		counter: newCounter(),
	}
}

// readCounter wraps an io.reader and stores bytes and number of operations
type readCounter struct {
	reader io.Reader
	counter
}

// Read wraps r.reader's Read method, and saves data about bytes read
func (r *readCounter) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.addBytes(int64(n))
	return n, err
}

// ReadCount returns the number of bytes read and operations performed
func (r *readCounter) ReadCount() (n int64, nops int) {
	return r.getCount()
}

// NewWriteCounter initializes a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		writer:  w,
		counter: newCounter(),
	}
}

// writeCounter wraps an io.reader and stores bytes and number of operations
type writeCounter struct {
	writer io.Writer
	counter
}

// Write wraps r.reader's Write method, and saves data about bytes read
func (w *writeCounter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.addBytes(int64(n))
	return n, err
}

// WriteCount returns the number of bytes read and operations performed
func (w *writeCounter) WriteCount() (n int64, nops int) {
	return w.getCount()
}

// NewReadWriteCounter initializes a new ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		ReadCounter:  NewReadCounter(rw),
		WriteCounter: NewWriteCounter(rw),
	}
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}
