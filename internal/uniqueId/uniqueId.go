package uniqueId

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

const (
	// base62 character set
	base string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Random integer ceil value
	maxRandomIntCeil int64 = 9999999999999

	// UUID size
	expectedIDSize int = 14

	// Timestamp of 1st Jan 2014, in nanosecond precision
	firstJan2014EpochTs int64 = 1388534400 * 1000 * 1000 * 1000
)

func init() {
	rand.Seed(int64(randUint32()))
}

func New() (string, error) {
	nanotime := time.Now().UnixNano()

	random := rand.Int63n(maxRandomIntCeil)
	base62Rand := base62Encode(random)

	// We need exactly 4 chars. If greater than 4, strip and use the last 4 chars
	if len(base62Rand) > 4 {
		base62Rand = base62Rand[len(base62Rand)-4:]
	}

	// If less than 4, left pad with '0'
	base62Rand = fmt.Sprintf("%04s", base62Rand)

	b62 := base62Encode(nanotime - firstJan2014EpochTs)
	id := b62 + base62Rand

	if len(id) != expectedIDSize {
		return id, fmt.Errorf("length mismatch when generating a new id: %s", id)
	}

	return id, nil
}

func base62Encode(num int64) string {
	index := base
	res := ""

	for {
		res = string(index[num%62]) + res
		num = int64(num / 62)
		if num == 0 {
			break
		}
	}
	return res
}

// randUint32 returns a random uint32 using crypto/rand which should in turn be
// used for seeding math/rand.
func randUint32() uint32 {
	buf := make([]byte, 4)
	// This panic is very unlikely(refer crypto/rand). Anyway this
	// function should not be used regularly but for one time seeding etc.
	if _, err := cryptorand.Reader.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v;", err))
	}
	// Using BigEndian or LittleEndian does not matter here.
	return binary.BigEndian.Uint32(buf)
}
