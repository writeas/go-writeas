package writeas

import (
	"testing"
)

func TestMarkdown(t *testing.T) {
	dwac := NewDevClient()

	in := "This is *formatted* in __Markdown__."
	out := `<p>This is <em>formatted</em> in <strong>Markdown</strong>.</p>
`

	res, err := dwac.Markdown(in, "")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	}

	if res != out {
		t.Errorf(`Got: '%s'
Expected: '%s'`, res, out)
	}
}
