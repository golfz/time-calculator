package timeutils

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all 00:00:00, expect 00:00:00",
			args: args{
				a: "00:00:00",
				b: "00:00:00",
			},
			want: "00:00:00",
		},
		{
			name: "no overflow, expect 12:12:12",
			args: args{
				a: "02:04:07",
				b: "10:08:05",
			},
			want: "12:12:12",
		},
		{
			name: "second overflow, expect 12:12:12",
			args: args{
				a: "02:03:17",
				b: "10:08:55",
			},
			want: "12:12:12",
		},
		{
			name: "minute overflow, expect 12:12:12",
			args: args{
				a: "01:14:07",
				b: "10:58:05",
			},
			want: "12:12:12",
		},
		{
			name: "minute, second overflow, expect 12:12:12",
			args: args{
				a: "01:13:17",
				b: "10:58:55",
			},
			want: "12:12:12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTimeList(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no input, expect 00:00:00",
			args: args{s: []string{}},
			want: "00:00:00",
		},
		{
			name: "1 input, expect 01:02:03",
			args: args{s: []string{"01:02:03"}},
			want: "01:02:03",
		},
		{
			name: "2 input, expect 02:03:04",
			args: args{s: []string{"01:02:03", "01:01:01"}},
			want: "02:03:04",
		},
		{
			name: "3 input with overflow, expect 12:12:12",
			args: args{s: []string{"01:40:30", "01:40:50", "08:50:52"}},
			want: "12:12:12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTimeList(tt.args.s); got != tt.want {
				t.Errorf("SumTimeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
