package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Regular test 1",
			args: args{
				str: "a4bc2d5e",
			},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name: "Regular test 2",
			args: args{
				str: "abcd",
			},
			want:    "abcd",
			wantErr: false,
		},
		{
			name: "Regular test 3",
			args: args{
				str: `qwe\4\5`,
			},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name: "Regular test 4",
			args: args{
				str: `qwe\45`,
			},
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name: "Regular test 5",
			args: args{
				str: `qwe\\5`,
			},
			want:    `qwe\\\\\`,
			wantErr: false,
		},
		{
			name: "Regular test 6",
			args: args{
				str: ``,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Regular test 6",
			args: args{
				str: `п10ривет`,
			},
			want:    "ппппппппппривет",
			wantErr: false,
		},
		{
			name: "Incorrect stroke 1",
			args: args{
				str: "45",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Incorrect stroke 2",
			args: args{
				str: `eq2\`,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Incorrect stroke 3",
			args: args{
				str: `1rw`,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() = |%v|, want |%v|", got, tt.want)
			}
		})
	}
}
