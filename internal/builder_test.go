package internal_test

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/antklim/pony/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildPage(t *testing.T) {
	d, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(d)
	require.NoError(t, err)

	const ttmpl = `
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>{{ .title }}</title>
		</head>
		<body>
		</body>
		</html>`

	tmpl, err := template.New("index.html").Parse(ttmpl)
	require.NoError(t, err)

	meta := &internal.Meta{
		Pages: map[string]internal.Page{
			"index": internal.Page{
				ID:         "index",
				Name:       "Home page",
				Path:       "/",
				Template:   "index.html",
				Properties: map[string]string{"title": "Home Page", "header": "Welcome to the home page"},
			},
			"about": internal.Page{
				ID:         "about",
				Name:       "About page",
				Path:       "/about",
				Template:   "about.html",
				Properties: map[string]string{"title": "About Page", "header": "Welcome to the about page"},
			},
		},
	}
	b := internal.NewBuilder(meta, tmpl)

	testCases := []struct {
		desc   string
		pageID string
		assert func(t *testing.T, buf bytes.Buffer, err error)
	}{
		{
			desc:   "returns error when unknown page provided",
			pageID: "foo",
			assert: func(t *testing.T, buf bytes.Buffer, err error) {
				assert.EqualError(t, err, "page foo not found in provided configuration")
			},
		},
		{
			desc:   "returns error when no template found for page",
			pageID: "about",
			assert: func(t *testing.T, buf bytes.Buffer, err error) {
				assert.EqualError(t, err, `html/template: "about.html" is undefined`)
			},
		},
		{
			desc:   "writes page to the provided writer",
			pageID: "index",
			assert: func(t *testing.T, buf bytes.Buffer, err error) {
				assert.NoError(t, err)
				assert.Contains(t, buf.String(), "<title>Home Page</title>")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var buf bytes.Buffer
			err := b.BuildPage(tC.pageID, &buf)
			tC.assert(t, buf, err)
		})
	}
}

func TestGeneratePages(t *testing.T) {
	d, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(d)
	require.NoError(t, err)

	const ttmpl = `
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>{{ .title }}</title>
		</head>
		<body>
		</body>
		</html>`

	tmpl, err := template.New("index.html").Parse(ttmpl)
	require.NoError(t, err)

	meta := &internal.Meta{
		Pages: map[string]internal.Page{
			"index": internal.Page{
				ID:         "index",
				Name:       "Home page",
				Path:       "/",
				Template:   "index.html",
				Properties: map[string]string{"title": "Home Page", "header": "Welcome to the home page"},
			},
			"about": internal.Page{
				ID:         "about",
				Name:       "About page",
				Path:       "/about",
				Template:   "index.html",
				Properties: map[string]string{"title": "About Page", "header": "Welcome to the about page"},
			},
		},
	}
	b := internal.NewBuilder(meta, tmpl)
	b.OutDir = d

	err = b.GeneratePages()
	require.NoError(t, err)

	testCases := []struct {
		file     string
		contains string
	}{
		{
			file:     filepath.Join(d, "index.html"),
			contains: "<title>Home Page</title>",
		},
		{
			file:     filepath.Join(d, "/about", "index.html"),
			contains: "<title>About Page</title>",
		},
	}

	for _, tC := range testCases {
		buf, err := ioutil.ReadFile(tC.file)
		require.NoError(t, err)
		assert.Contains(t, string(buf), tC.contains)
	}
}
