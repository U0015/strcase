package strcase

import "testing"

func TestToCamel(t *testing.T) {
	cases := [...]struct{ in, exp string }{
		{``, ``},
		{`abcabc`, `Abcabc`},
		{`abcAbc`, `AbcAbc`},
		{`ABCabc`, `ABCabc`},
		{`abc abc`, `AbcAbc`},
		{`abc_abc`, `AbcAbc`},
		{`abc-abc`, `AbcAbc`},
		{`  abc  abc    `, `AbcAbc`},
		{`__abc__abc____`, `AbcAbc`},
		{`--abc--abc----`, `AbcAbc`},
		{`абвабв`, `Абвабв`},
		{`абвАбв`, `АбвАбв`},
	}
	for _, c := range cases {
		if out := ToCamel(c.in); out != c.exp {
			t.Errorf("for: %q, exp: %q, got: %q", c.in, c.exp, out)
		}
	}
}

func BenchmarkToCamel(b *testing.B) {
	c, r := `abcABCabc abc_abcAbc`, ``

	for n := 0; n < b.N; n++ {
		r = ToCamel(c)
	}
	_ = r
}

func BenchmarkToCamelWithEncodeRune(b *testing.B) {
	c, r := `abcABCabc abc_абвАбв`, ``

	for n := 0; n < b.N; n++ {
		r = ToCamel(c)
	}
	_ = r
}