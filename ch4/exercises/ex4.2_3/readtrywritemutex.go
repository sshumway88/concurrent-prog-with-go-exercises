package listing4_12

import "sync"

// listng 4.12
type ReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

// listing 4.13
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

func (rw *ReadWriteMutex) ReaderUnlock() {
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

// exercise 4.2...this was...rather simple
func (rw *ReadWriteMutex) TryWriteLock() bool {
	return rw.globalLock.TryLock()
}

// exercise 4.3 - implement TryReadLock
func (rw *ReadWriteMutex) TryReadLock() bool {
	if rw.readersLock.TryLock() {
        globalLockObtained := true
        if rw.readersCounter == 0 {
            globalLockObtained = rw.globalLock.TryLock()
        }
        if globalLockObtained {
            rw.readersCounter++
        }
        rw.readersLock.Unlock()
        return globalLockObtained
    } 
	return false
}