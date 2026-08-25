package storage

import (
	"testing"

	"github.com/minio/minio-go/v7"
)

func TestParseBucketLookup(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  minio.BucketLookupType
	}{
		{name: "default", input: "", want: minio.BucketLookupAuto},
		{name: "auto", input: "auto", want: minio.BucketLookupAuto},
		{name: "path", input: " path ", want: minio.BucketLookupPath},
		{name: "dns", input: "DNS", want: minio.BucketLookupDNS},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBucketLookup(tt.input)
			if err != nil {
				t.Fatalf("ParseBucketLookup(%q) returned an error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("ParseBucketLookup(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseBucketLookupRejectsUnknownMode(t *testing.T) {
	if _, err := ParseBucketLookup("unsupported"); err == nil {
		t.Fatal("ParseBucketLookup should reject an unsupported mode")
	}
}
