package core

import (
	"sync"

	"github.com/pkg/errors"
)

// NewHistory creates an history.
func NewHistory() *History {
	return &History{
		index:   -1,
		history: make([]string, 0, 32),
	}
}

// History represents a store that contains URL ordered chronologically.
type History struct {
	mutex   sync.RWMutex
	index   int
	history []string
}

// Len returns the number of entries recorded in the history.
func (h *History) Len() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return len(h.history)
}

// Current returns the current entry.
func (h *History) Current() (url string, err error) {
	if h.Len() == 0 {
		return "", errors.New("history does not have entries")
	}

	h.mutex.RLock()
	defer h.mutex.RUnlock()

	url = h.history[h.index]
	return url, nil
}

// NewEntry adds an entry to the history.
// The entry is added after the current one.
// All the entries that was after the current one are removed.
func (h *History) NewEntry(url string) {
	l := h.Len()

	h.mutex.Lock()
	defer h.mutex.Unlock()

	var history []string
	if l == 0 {
		history = h.history
	} else {
		history = h.history[:h.index+1]
	}

	h.history = append(history, url)
	h.index++
}

// CanPrevious reports whether there is a previous entry.
func (h *History) CanPrevious() bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.index > 0
}

// Previous returns the previous entry.
func (h *History) Previous() (url string, err error) {
	if !h.CanPrevious() {
		return "", errors.New("history does not have a previous entry to return")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.index--
	url = h.history[h.index]
	return url, nil
}

// CanNext reports whether there is a next entry.
func (h *History) CanNext() bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.index < h.Len()-1
}

// Next returns the next entry.
func (h *History) Next() (url string, err error) {
	if !h.CanNext() {
		return "", errors.New("history does not have a next entry to return")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.index++
	url = h.history[h.index]
	return url, nil
}
