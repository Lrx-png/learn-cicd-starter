package auth

import (
	"net/http"
	"testing"
	"errors" // Add this line
)

// TestGetAPIKey tests the GetAPIKey function with various scenarios.
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantKey     string
		wantErr     error
	}{
		{
			name:    "No Authorization Header",
			headers: http.Header{},
			wantKey: "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "Malformed Authorization Header",
			headers: http.Header{"Authorization": []string{"Bearer token"}},
			wantKey: "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": []string{"ApiKey mysecretkey"}},
			wantKey: "mysecretkey",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotErr := GetAPIKey(tt.headers)

			// Check if the returned key matches the expected key
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}

			// Check if the error matches the expected error
			if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) || 
			   (gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
