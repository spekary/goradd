package strings

import "testing"

func TestShiftPath(t *testing.T) {
	tests := []struct {
		name              string
		path              string
		wantTopName       string
		wantRemainingPath string
	}{
		{"Empty", "", "", ""},
		{"Root", "/", "/", ""},
		{"One", "/one", "/one", ""},
		{"RelOne", "one", "one", ""},
		{"RelOneDir", "one/", "one", "/"},
		{"Two", "/one/two", "/one", "/two"},
		{"RelTwo", "one/two", "one", "/two"},
		{"Three", "/one/two/three", "/one", "/two/three"},
		{"RelThree", "one/two/three", "one", "/two/three"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTopName, gotRemainingPath := ShiftPath(tt.path)
			if gotTopName != tt.wantTopName {
				t.Errorf("ShiftPath() gotTopName = %v, want %v", gotTopName, tt.wantTopName)
			}
			if gotRemainingPath != tt.wantRemainingPath {
				t.Errorf("ShiftPath() gotRemainingPath = %v, want %v", gotRemainingPath, tt.wantRemainingPath)
			}
		})
	}
}
