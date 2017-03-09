package urlfilter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		URL   string
		valid bool
	}{
		{
			URL:   "google.com",
			valid: false,
		},
		{
			URL:   "http//google.com",
			valid: false,
		},
		{
			URL:   "foo://example.com:8042/over/there?name=ferret#nose",
			valid: true,
		},
		{
			URL:   "http://www.farina.org/blog/categories/tags/about.html",
			valid: true,
		},
	}

	uf := URLFilter{}

	for _, testCase := range testCases {
		require.Equal(t, testCase.valid, uf.Validate(testCase.URL), "url: "+testCase.URL)
	}
}
