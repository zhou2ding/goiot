package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		s    string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": {s: "babcbef", sep: "b", want: []string{"", "a", "c", "ef"}},
		"case_2": {s: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case_3": {s: "abcef", sep: "bc", want: []string{"a", "ef"}},
		"case_4": {s: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又", "河"}},
	}
	// 原始测试组，无法进行子测试
	// testGroup := []testCase{
	// 	{
	// 		s:    "babcbef",
	// 		sep:  "b",
	// 		want: []string{"", "a", "c", "ef"},
	// 	},
	// 	{
	// 		s:    "a:b:c",
	// 		sep:  ":",
	// 		want: []string{"a", "b", "c"},
	// 	},
	// 	{
	// 		s:    "abcef",
	// 		sep:  "bc",
	// 		want: []string{"a", "ef"},
	// 	},
	// 	{
	// 		s:    "沙河有沙又有河",
	// 		sep:  "有",
	// 		want: []string{"沙河", "沙又", "河"},
	// 	},
	// }
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%v, but got:%v\n", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d", ":")
	}
}
