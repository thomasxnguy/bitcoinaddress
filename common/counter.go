package common

import "sync/atomic"

// Count is an atomic counter to generate unique index
type Count uint32

func (i *Count) Inc() uint32 {
	return atomic.AddUint32((*uint32)(i), 1)
}

func (i *Count) Get() uint32 {
	return atomic.LoadUint32((*uint32)(i))
}
