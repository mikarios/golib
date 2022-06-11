package hid

import (
	"bytes"
	"math/big"
	"sync"
	"time"
)

var (
	randomKeeper     = &rand{}
	DefaultSeparator = '-'
)

type rand struct {
	lastNumber int64
	sync.Mutex
}

// New creates a new human-readable ID that consists of numbers and letters. If splitEvery is <=0 then separator
// can be nil. If splitEvery is >0 and separator is nil then the DefaultSeparator will be used.
func New(includeCaps bool, splitEvery int, separator *rune) string {
	randomKeeper.Lock()
	defer randomKeeper.Unlock()

	now := time.Now().UnixNano()
	if now == randomKeeper.lastNumber {
		time.Sleep(1 * time.Nanosecond)
		now = time.Now().UnixNano()
	}

	randomKeeper.lastNumber = now

	base := 62
	if !includeCaps {
		base = 36
	}

	res := big.NewInt(randomKeeper.lastNumber).Text(base)

	if splitEvery <= 0 {
		return res
	}

	if separator == nil {
		separator = &DefaultSeparator
	}

	return insertNth(res, splitEvery, *separator)
}

func insertNth(str string, splitEvery int, separator rune) string {
	var (
		buffer      bytes.Buffer
		splitEvery1 = splitEvery - 1
		len1        = len(str) - 1
	)

	for i, r := range str {
		buffer.WriteRune(r)

		if i%splitEvery == splitEvery1 && i != len1 {
			buffer.WriteRune(separator)
		}
	}

	return buffer.String()
}
