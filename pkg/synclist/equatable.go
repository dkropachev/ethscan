package synclist

import "sync"

type equatable[T any] interface {
	Equal(T) bool
}

type EquatableList[T equatable[T]] struct {
	lock sync.RWMutex
	list []T
}

func (l *EquatableList[T]) AppendIfNotExists(item ...T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	var ret bool

outer:
	for _, it := range item {
		for _, i := range l.list {
			if i.Equal(it) {
				continue outer
			}
		}
		ret = true
		l.list = append(l.list, it)
	}
	return ret
}

func (l *EquatableList[T]) Append(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = append(l.list, item)
}

func (l *EquatableList[T]) Get() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list
}

func (l *EquatableList[T]) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.list)
}

func (l *EquatableList[T]) GetAll() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	out := make([]T, len(l.list))
	copy(out, l.list)
	return out
}

func (l *EquatableList[T]) Contains(item ...T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for _, i := range l.list {
		for _, it := range item {
			if i.Equal(it) {
				return true
			}
		}
	}
	return false
}
