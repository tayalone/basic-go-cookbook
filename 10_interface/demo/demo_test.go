package demo

import "testing"

type fakeRepo struct{}

func (fakeRepo) QueryLang(uint) Langauge {
	return Langauge{
		ID:   1,
		Name: "go",
	}
}

func TestHandle(t *testing.T) {
	h := NewHandler(fakeRepo{})
	s := h.do(1)

	want := "go language\n"

	if s != want {
		t.Errorf("want %q but get %s", want, s)
	}
}
