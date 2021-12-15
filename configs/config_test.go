package configs

import "testing"

func TestGetDBConnectionInfo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: 	"DBConnectionInfoTest",
			want: "user:pass@tcp(localhost:3306)/app?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDBConnectionInfo(); got != tt.want {
				t.Errorf("GetDBConnectionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServerPort(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestGetServerPort",
			want: ":8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerPort(); got != tt.want {
				t.Errorf("GetServerPort() = %v, want %v", got, tt.want)
			}
		})
	}
}