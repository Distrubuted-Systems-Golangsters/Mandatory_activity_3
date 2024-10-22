package util

type Process struct {
	Timestamp int32
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}