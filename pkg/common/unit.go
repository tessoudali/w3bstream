package common

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KiB kilobyte
	KiB = 1 << (10 * iota) // 1 << (10*1)
	// MiB megabyte
	MiB // 1 << (10*2)
	// GiB gigabyte
	GiB // 1 << (10*3)
	// TiB terabyte
	TiB // 1 << (10*4)
)
