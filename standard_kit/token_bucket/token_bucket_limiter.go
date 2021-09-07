package main

import (
	"sync"
	"sync/atomic"
	"time"
)

// 每分钟 匀速放入60个token
type TokenBucketLimiter struct {
	duration time.Duration // 每隔多久放入1个token
	tokenSize uint64
	ticker *time.Ticker
	tokenLimit uint64
	mutex sync.RWMutex
}

func newLimiter(tokenLimit uint64, duration time.Duration) *TokenBucketLimiter {
	l := &TokenBucketLimiter{
		tokenLimit: tokenLimit,
		duration: duration,
		ticker: time.NewTicker(duration),
	}

	go func() {
		for {
			select {
			case <-l.ticker.C:
				if l.tokenSize < l.tokenLimit {
					atomic.AddUint64(&l.tokenSize, 1)
					//fmt.Println("produce token:", l.tokenSize)
				}
			}
		}
	}()
	return l
}

func (l *TokenBucketLimiter)allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	//fmt.Println("token size:", l.tokenSize)
	if l.tokenSize > 0 {
		l.tokenSize--
		//fmt.Println("consume token:", l.tokenSize)
		return true
	} else {
		return false
	}
}
