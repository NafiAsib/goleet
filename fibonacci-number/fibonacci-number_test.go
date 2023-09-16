package leetcode

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
