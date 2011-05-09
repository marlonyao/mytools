package mytools

import "testing"

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest {
	ReverseTest{"ABCD", "DCBA"},
	ReverseTest{"Hello, 世界", "界世 ,olleH"},
}

func TestReverse(t *testing.T) {
	for _, r := range ReverseTests {
		o := Reverse(r.in)
		if o != r.out {
			t.Errorf("%q expected %q got %q", r.in, r.out, o)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}
