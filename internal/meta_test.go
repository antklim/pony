package internal_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/antklim/pony/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadMeta(t *testing.T) {
	const metadata = `
template: index
pages:
  index:
    name: Home page
    path: /
    properties: 
      - key: title
        value: Home Page
      - key: header
        value: Welcome to the home page
  about:
    name: About page
    path: /about
    template: about
    properties: 
      - key: title
        value: About Page
      - key: header
        value: Welcome to the about page`

	d, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(d)
	require.NoError(t, err)

	f, err := ioutil.TempFile(d, "pony*.yaml")
	require.NoError(t, err)

	_, err = f.Write([]byte(metadata))
	require.NoError(t, err)

	meta, err := internal.LoadMeta(f.Name())
	require.NoError(t, err)

	expected := &internal.Meta{
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
	assert.Equal(t, expected, meta)
}
