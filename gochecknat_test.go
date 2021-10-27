package gochecknat

import (
	"fmt"
	"testing"
)

func TestGetNATInfo(t *testing.T) {
	tests := []struct {
		name     string
		wantInfo NATInfo
		wantErr  bool
	}{
		{
			name: "test",
			wantInfo: NATInfo{
				IP:         "",
				Port:       0,
				Candidates: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := GetNATInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNATInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotInfo.Candidates) == len(tt.wantInfo.Candidates) {
				t.Errorf("GetNATInfo() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotInfo.IP == tt.wantInfo.IP {
				t.Errorf("GetNATInfo() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotInfo.Port == tt.wantInfo.Port {
				t.Errorf("GetNATInfo() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			fmt.Printf("\nPUBLIC IP: %v\n", gotInfo.IP)
			fmt.Printf("ASSIGNED PORT: %v\n", gotInfo.Port)
			fmt.Printf("SYMMERTIC NAT: %v\n\n", gotInfo.Symmetric)
		})
	}
}
