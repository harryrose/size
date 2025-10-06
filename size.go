package size

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	BytesSuffix     = "B"
	KilobytesSuffix = "KB"
	MegabytesSuffix = "MB"
	GigabytesSuffix = "GB"
	TerabytesSuffix = "TB"
	PetabytesSuffix = "PB"
)

const (
	Kilobyte Size = 1024
	Megabyte      = 1024 * Kilobyte
	Gigabyte      = 1024 * Megabyte
	Terabyte      = 1024 * Gigabyte
	Petabyte      = 1024 * Terabyte
)

// ParseSize parses a size string (e.g., "10MB", "1.5 GB") and returns the size in bytes.
// If no suffix is provided, bytes are assumed.  If a suffix is provided, it must be one of
// "B", "KB", "MB", "GB", "TB", or "PB" (case-insensitive).  The 'B' component of the suffix
// is optional (e.g. 1K == 1KB).  A space between the number and the suffix is allowed but
// not required.
func ParseSize(s string) (int64, error) {
	reg, _ := regexp.Compile(`^([0-9]+(\.[0-9]+)?)[[:space:]]*([KkMmGgTtPp]?B?)$`)
	matches := reg.FindStringSubmatch(s)
	if len(matches) != 4 {
		return 0, fmt.Errorf("invalid size: %v", s)
	}
	number, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid size: %v", s)
	}

	switch strings.ToUpper(matches[3]) {
	case "K", "KB":
		return int64(number * float64(Kilobyte)), nil
	case "M", "MB":
		return int64(number * float64(Megabyte)), nil
	case "G", "GB":
		return int64(number * float64(Gigabyte)), nil
	case "T", "TB":
		return int64(number * float64(Terabyte)), nil
	case "P", "PB":
		return int64(number * float64(Petabyte)), nil
	case "B", "":
		return int64(number), nil
	default:
		return 0, fmt.Errorf("invalid size suffix: %v", matches[3])
	}
}

// Size represents a size in bytes.
type Size int64

// Bytes returns the size in bytes.
func (s Size) Bytes() int64 {
	return int64(s)
}

// Kilobytes returns the size in kilobytes.
func (s Size) Kilobytes() float64 {
	return float64(s) / float64(Kilobyte)
}

// Megabytes returns the size in megabytes.
func (s Size) Megabytes() float64 {
	return float64(s) / float64(Megabyte)
}

// Gigabytes returns the size in gigabytes.
func (s Size) Gigabytes() float64 {
	return float64(s) / float64(Gigabyte)
}

// Terabytes returns the size in terabytes.
func (s Size) Terabytes() float64 {
	return float64(s) / float64(Terabyte)
}

// Petabytes returns the size in petabytes.
func (s Size) Petabytes() float64 {
	return float64(s) / float64(Petabyte)
}

// Abs returns the absolute value of the size.
func (s Size) Abs() Size {
	if s < 0 {
		return -s
	}
	return s
}

// String returns a human-readable representation of the size, using the largest
// appropriate unit (B, KB, MB, GB, TB, PB) with two decimal places of precision.
func (s Size) String() string {
	abs := s.Abs()
	suffix := BytesSuffix
	val := 0.0
	switch {
	case abs < Kilobyte:
		//shortcut here as we don't need decimal places for bytes
		return fmt.Sprintf("%d %s", s.Bytes(), BytesSuffix)

	case abs < Megabyte:
		suffix = KilobytesSuffix
		val = s.Kilobytes()

	case abs < Gigabyte:
		suffix = MegabytesSuffix
		val = s.Megabytes()

	case abs < Terabyte:
		suffix = GigabytesSuffix
		val = s.Gigabytes()

	case abs < Petabyte:
		suffix = TerabytesSuffix
		val = s.Terabytes()

	default:
		suffix = PetabytesSuffix
		val = s.Petabytes()
	}

	return fmt.Sprintf("%.2f %s", val, suffix)
}
