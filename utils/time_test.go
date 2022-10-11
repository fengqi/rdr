package utils

import (
	"testing"
)

func TestTimestampToTime(t *testing.T) {
	cases := map[int64]string{
		1665555454199: "2022-10-12 14:17:34",
	}

	for k, want := range cases {
		tt := TimestampToTime(k)
		give := tt.Format("2006-01-02 15:04:05")
		if want != give {
			t.Errorf("TimestampToTime(%d) give: %s, want: %s", k, give, want)
		}
	}
}
