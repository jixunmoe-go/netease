package ip

import (
	"fmt"
	"math/rand"
)

func Format(value uint32) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		0xFF&(value>>24),
		0xFF&(value>>16),
		0xFF&(value>>8),
		0xFF&value,
	)
}

func Make(a, b, c, d uint8) uint32 {
	return (uint32(a) << 24) | (uint32(b) << 16) | (uint32(c) << 8) | uint32(d)
}

func GetRandomIP(begin, end uint32) uint32 {
	return begin + uint32(rand.Intn(int(end-begin)))
}
