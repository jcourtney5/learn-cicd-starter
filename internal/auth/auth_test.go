package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantApiKey string
		wantErr    bool
	}{
		{
			name: "Valid ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantApiKey: "valid_api_key",
			wantErr:    false,
		},
		{
			name:       "Missing Authorization header",
			headers:    http.Header{},
			wantApiKey: "",
			wantErr:    false,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidApiKey api_key"},
			},
			wantApiKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiKey != tt.wantApiKey {
				t.Errorf("GetAPIKey() gotApiKey = %v, want %v", gotApiKey, tt.wantApiKey)
			}
		})
	}
}
