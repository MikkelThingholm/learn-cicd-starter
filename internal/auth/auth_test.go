package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input   http.Header
		want    string
		wantErr bool
	}{
		{
			input:   map[string][]string{"Authorization": {"ApiKey abc"}},
			want:    "abc",
			wantErr: false,
		},
		{
			input:   map[string][]string{"Authorizatio": {"ApiKey abc"}},
			want:    "",
			wantErr: true,
		},
		{
			input:   map[string][]string{"Authorization": {"abc"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)

		if tc.wantErr && err == nil {
			t.Fatalf("expected error")
		} else if !tc.wantErr && err != nil {
			t.Fatalf("unexpected error")
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got %v", tc.want, got)
		}
	}
}
