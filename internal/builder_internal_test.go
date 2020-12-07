package internal

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildPage(t *testing.T) {
	dir, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(dir)
	require.NoError(t, err)

	const ttmpl = `
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ .foo }}</title>
</head>
<body>
</body>
</html>`

	tmpl, err := template.New("testTemplate").Parse(ttmpl)
	require.NoError(t, err)

	testCases := []struct {
		desc         string
		id           string
		page         Page
		expectedTmpl string
	}{
		{
			desc: "builds page to the root directory",
			id:   "index",
			page: Page{
				Name:       "Pony root",
				Path:       "/",
				Properties: []Property{{Key: "foo", Value: "root"}},
			},
			expectedTmpl: "<title>root</title>",
		},
		{
			desc: "builds page to the sub-directory",
			id:   "test",
			page: Page{
				Name:       "Pony test",
				Path:       "/test",
				Properties: []Property{{Key: "foo", Value: "test"}},
			},
			expectedTmpl: "<title>test</title>",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := buildPage(tC.id, tC.page, tmpl, dir)
			require.NoError(t, err)

			fname := filepath.Join(dir, tC.page.Path, tC.id+".html")
			finfo, err := os.Stat(fname)
			require.NoError(t, err)
			assert.True(t, finfo.Size() > 0)
			assert.False(t, finfo.IsDir())

			buf, err := ioutil.ReadFile(fname)
			require.NoError(t, err)
			assert.True(t, strings.Contains(string(buf), tC.expectedTmpl))
		})
	}
}
