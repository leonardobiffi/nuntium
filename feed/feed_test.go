package feed

import (
	"testing"
)

func TestFetch(t *testing.T) {
	var diffHours = float64(24)
	type args struct {
		feedURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"AWS", args{"https://aws.amazon.com/blogs/aws/feed"}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNews, err := Fetch(tt.args.feedURL, diffHours)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotNews) == 0 {
				t.Errorf("Fetch() = not got any news")
			}
		})
	}
}
