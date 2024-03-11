package mutex

import (
	"context"
	"time"

	"github.com/bsm/redislock"

	"github.com/razorpay/retail-store/internal/errors"
)

type MutexWrapperMethods interface {
	Obtain(ctx context.Context, key string, retryCount int) (*redislock.Lock, error)
	Release(ctx context.Context, lock *redislock.Lock) error
}

type MutexWrapper struct {
	Client *redislock.Client
}

func InitializeMutex(client redislock.RedisClient) *MutexWrapper {
	var mutex MutexWrapper
	mutex.Client = redislock.New(client)
	return &mutex
}

var Mutex MutexWrapperMethods

type MultipleFuncHandler func([]string, []string) (interface{}, *errors.ErrorData)

func AcquireAndReleaseMultiple(ctx context.Context, keys []string, method MultipleFuncHandler) (interface{}, *errors.ErrorData) {
	var lockAcquireFailedKeys []string
	var lockAcquiredKeys []string
	var locksToRelease []*redislock.Lock
	// in a loop you acquire all of them
	// make a list of the once failed
	// defer call the release on all that went through
	for _, key := range keys {
		lock, err := Mutex.Obtain(ctx, key, 1)
		if err != nil {
			lockAcquireFailedKeys = append(lockAcquireFailedKeys, key)
			continue
		}
		locksToRelease = append(locksToRelease, lock)
		lockAcquiredKeys = append(lockAcquiredKeys, key)
	}

	// after work is done release the lock (defer it)
	defer releaseMultiple(ctx, locksToRelease)

	// execute the critical section method
	return method(lockAcquireFailedKeys, lockAcquiredKeys)
}

func releaseMultiple(ctx context.Context, locks []*redislock.Lock) {
	for _, l := range locks {
		_ = Mutex.Release(ctx, l)
	}
}

func (m *MutexWrapper) Obtain(_ context.Context, key string, retryCount int) (*redislock.Lock, error) {

	lock, err := m.Client.Obtain(key, 500*time.Second, &redislock.Options{RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(time.Second), retryCount)})

	return lock, err
}
func (m *MutexWrapper) Release(_ context.Context, lock *redislock.Lock) error {
	err := lock.Release()
	return err
}
