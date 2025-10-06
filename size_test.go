package size

import (
	"math"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		input   string
		want    int64
		wantErr bool
	}{
		{"1234", 1234, false},
		{"1024B", 1024, false},
		{"1024 B", 1024, false},

		{"1KB", 1024, false},
		{"1 KB", 1024, false},

		{"1K", 1024, false},
		{"1 K", 1024, false},

		{"1.5KB", 1536, false},
		{"1.5 KB", 1536, false},

		{"1.5K", 1536, false},
		{"1.5 K", 1536, false},

		{"2MB", 2 * 1024 * 1024, false},
		{"2 MB", 2 * 1024 * 1024, false},

		{"2M", 2 * 1024 * 1024, false},
		{"2 M", 2 * 1024 * 1024, false},

		{"0.5GB", 512 * 1024 * 1024, false},
		{"0.5 GB", 512 * 1024 * 1024, false},

		{"0.5G", 512 * 1024 * 1024, false},
		{"0.5 G", 512 * 1024 * 1024, false},

		{"1TB", 1024 * 1024 * 1024 * 1024, false},
		{"1 TB", 1024 * 1024 * 1024 * 1024, false},

		{"1T", 1024 * 1024 * 1024 * 1024, false},
		{"1 T", 1024 * 1024 * 1024 * 1024, false},

		{"1PB", 1024 * 1024 * 1024 * 1024 * 1024, false},
		{"1 PB", 1024 * 1024 * 1024 * 1024 * 1024, false},

		{"1P", 1024 * 1024 * 1024 * 1024 * 1024, false},
		{"1 P", 1024 * 1024 * 1024 * 1024 * 1024, false},

		{"invalid", 0, true},
		{"1XB", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseSize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSize(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseSize(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestSizeString(t *testing.T) {
	tests := []struct {
		input Size
		want  string
	}{
		{Size(512), "512 B"},
		{Size(1024), "1.00 KB"},
		{Size(1536), "1.50 KB"},
		{Size(1048576), "1.00 MB"},
		{Size(1073741824), "1.00 GB"},
		{Size(1099511627776), "1.00 TB"},
		{Size(1125899906842624), "1.00 PB"},
		{Size(-2048), "-2.00 KB"},
	}

	for _, tt := range tests {
		got := tt.input.String()
		if got != tt.want {
			t.Errorf("Size(%d).String() = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestSizeMethods(t *testing.T) {
	tests := []struct {
		input     Size
		bytes     int64
		kilobytes float64
		megabytes float64
		gigabytes float64
		terabytes float64
		petabytes float64
	}{
		{Size(1024), 1024, 1, 1.0 / 1024, 1.0 / (1024 * 1024), 1.0 / (1024 * 1024 * 1024), 1.0 / (1024 * 1024 * 1024 * 1024)},
		{Size(1048576), 1048576, 1024, 1, 1.0 / 1024, 1.0 / (1024 * 1024), 1.0 / (1024 * 1024 * 1024)},
		{Size(1073741824), 1073741824, 1048576, 1024, 1, 1.0 / 1024, 1.0 / (1024 * 1024)},
		{Size(1099511627776), 1099511627776, 1073741824, 1048576, 1024, 1, 1.0 / 1024},
		{Size(1125899906842624), 1125899906842624, 1099511627776, 1073741824, 1048576, 1024, 1},
		{Size(-2048), -2048, -2, -2.0 / 1024, -2.0 / (1024 * 1024), -2.0 / (1024 * 1024 * 1024), -2.0 / (1024 * 1024 * 1024 * 1024)},
	}

	for _, tt := range tests {
		if got := tt.input.Bytes(); got != tt.bytes {
			t.Errorf("Size(%d).Bytes() = %d, want %d", tt.input, got, tt.bytes)
		}
		if got := tt.input.Kilobytes(); math.Abs(got-tt.kilobytes) > 1e-6 {
			t.Errorf("Size(%d).Kilobytes() = %f, want %f", tt.input, got, tt.kilobytes)
		}
		if got := tt.input.Megabytes(); math.Abs(got-tt.megabytes) > 1e-6 {
			t.Errorf("Size(%d).Megabytes() = %f, want %f", tt.input, got, tt.megabytes)
		}
		if got := tt.input.Gigabytes(); math.Abs(got-tt.gigabytes) > 1e-6 {
			t.Errorf("Size(%d).Gigabytes() = %f, want %f", tt.input, got, tt.gigabytes)
		}
		if got := tt.input.Terabytes(); math.Abs(got-tt.terabytes) > 1e-6 {
			t.Errorf("Size(%d).Terabytes() = %f, want %f", tt.input, got, tt.terabytes)
		}
		if got := tt.input.Petabytes(); math.Abs(got-tt.petabytes) > 1e-6 {
			t.Errorf("Size(%d).Petabytes() = %f, want %f", tt.input, got, tt.petabytes)
		}
	}
}
