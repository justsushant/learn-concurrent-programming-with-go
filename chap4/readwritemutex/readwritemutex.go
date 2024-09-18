package main

import "sync"

type ReadWriteMutex struct {
	readersCounter int        // counter for reader goroutines currently
	readersLock    sync.Mutex // mutex for reader sync
	globalLock     sync.Mutex // mutex for writer sync
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++

	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReadWriteMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	if rw.readersLock.TryLock() {
		rw.readersCounter++
		if rw.readersCounter == 0 {
			rw.globalLock.Lock()
		}

		rw.readersLock.Unlock()
		return true
	} else {
		return false
	}
}
