package substrings

import (
	"reflect"
	"testing"
)

func TestFindIndices(t *testing.T) {
	type args struct {
		textToSearch string
		subtext      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty textToSearch",
			args: args{
				textToSearch: "",
				subtext:      "foo",
			},
			want: nil,
		},
		{
			name: "empty subtext",
			args: args{
				textToSearch: "foo",
				subtext:      "",
			},
			want: nil,
		},
		{
			name: "subtext longer than textToSearch",
			args: args{
				textToSearch: "foo",
				subtext:      "foobar",
			},
			want: nil,
		},
		{
			name: "subtext not found - z",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "z",
			},
			want: nil,
		},
		{
			name: "subtext not found - Peterz",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "Peterz",
			},
			want: nil,
		},
		{
			name: "subtext found - Peter",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "Peter",
			},
			want: []int{1, 20, 75},
		},
		{
			name: "subtext found - peter",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "peter",
			},
			want: []int{1, 20, 75},
		},
		{
			name: "subtext found - pick",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "pick",
			},
			want: []int{30, 58},
		},
		{
			name: "subtext found - pi",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "pi",
			},
			want: []int{30, 37, 43, 51, 58},
		},
		{
			name: "subtext found - overlapping",
			args: args{
				textToSearch: "abababab",
				subtext:      "aba",
			},
			want: []int{1, 3, 5},
		},
		{
			name: "subtext found - 9 digit",
			args: args{
				textToSearch: "123456789",
				subtext:      "9",
			},
			want: []int{9},
		},
		{
			name: "subtext found - !",
			args: args{
				textToSearch: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
				subtext:      "!",
			},
			want: []int{92},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndices(tt.args.textToSearch, tt.args.subtext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLowerASCII(t *testing.T) {
	type args struct {
		char byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "A",
			args: args{
				char: 'A',
			},
			want: 'a',
		},
		{
			name: "Z",
			args: args{
				char: 'Z',
			},
			want: 'z',
		},
		{
			name: "a",
			args: args{
				char: 'a',
			},
			want: 'a',
		},
		{
			name: "z",
			args: args{
				char: 'z',
			},
			want: 'z',
		},
		{
			name: "1",
			args: args{
				char: '1',
			},
			want: '1',
		},
		{
			name: "!",
			args: args{
				char: '!',
			},
			want: '!',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerASCII(tt.args.char); got != tt.want {
				t.Errorf("toLowerASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
