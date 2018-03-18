package watcher

import (
	"testing"
)

func TestGenerateRegexp(t *testing.T) {
	tests := []struct {
		name   string
		filter string
		want   string
	}{
		{
			name:   "no filter",
			filter: "",
			want:   "",
		},
		{
			name:   "simple filter",
			filter: "filters",
			want:   "(filters)",
		},
		{
			name:   "two filters",
			filter: "java,laravel",
			want:   "(java)|(laravel)",
		},
		{
			name:   "test white spaces",
			filter: "java , laravel ",
			want:   "(java)|(laravel)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRegexp(tt.filter); got != tt.want {
				t.Errorf("GenerateRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainKeywords(t *testing.T) {
	type args struct {
		pattern    string
		paragraphs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check basic usage should return true",
			args: args{"(java)", []string{"java is the new c++"}},
			want: true,
		},
		{
			name: "check basic usage",
			args: args{"(swift)", []string{"java is the new c++"}},
			want: false,
		},
		{
			name: "check basic usage",
			args: args{"(swift)|(apple)", []string{"Apple releases new iphone"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainKeywords(tt.args.pattern, tt.args.paragraphs...); got != tt.want {
				t.Errorf("ContainKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}
