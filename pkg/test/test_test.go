package test

import (
	"testing"

	"github.com/jonnay101/domain-data-parser/pkg/emaildomainstats"
)

func Benchmark_emaildomainstats_GetDomainData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		emaildomainstats.GetDomainStats("big_test_data.csv")
	}
}

func TestGetDomainNameFrequency(t *testing.T) {
	dd, err := emaildomainstats.GetDomainStats("test_data.csv")
	if err != nil {
		t.Error(err)
		return
	}
	type args struct {
		domName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find bigcartel.com", args{"bigcartel.com"}, 2},
		{"domain not there", args{"not.there"}, 0},
		{"find goo.gl", args{"goo.gl"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dd.GetDomainNameFrequency(tt.args.domName)
			if got != tt.want {
				t.Errorf("wanted: %d but got: %d", tt.want, got)
			}
		})
	}
}
