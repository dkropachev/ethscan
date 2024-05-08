package synclist

import "sync"

type ComparableList[T comparable] struct {
	lock sync.RWMutex
	list []T
}

func (l *ComparableList[T]) AppendIfNotExists(item T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	for _, i := range l.list {
		if i == item {
			return false
		}
	}

	l.list = append(l.list, item)
	return true
}

func (l *ComparableList[T]) Append(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = append(l.list, item)
}

func (l *ComparableList[T]) Get() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list
}

func (l *ComparableList[T]) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.list)
}

func (l *ComparableList[T]) GetAll() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	var out []T
	copy(out, l.list)
	return out
}

func (l *ComparableList[T]) Contains(item ...T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for _, i := range l.list {
		for _, it := range item {
			if i == it {
				return true
			}
		}
	}
	return false
}

func (l *ComparableList[T]) Count(item ...T) int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	cnt := 0

	for _, i := range l.list {
		for _, it := range item {
			if i == it {
				cnt++
			}
		}
	}
	return cnt
}
