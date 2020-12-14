package pony_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/antklim/pony"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func storeTempFile(dir, payload, namePattern string) (string, error) {
	f, err := ioutil.TempFile(dir, namePattern)
	if err != nil {
		return "", err
	}

	if _, err := f.Write([]byte(payload)); err != nil {
		return "", err
	}

	return f.Name(), nil
}

func TestLoadMetadata(t *testing.T) {
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
        value: Welcome to the home page`

	tempDir, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)

	fOk, err := storeTempFile(tempDir, metadata, "pony*.yaml")
	require.NoError(t, err)

	fErr, err := storeTempFile(tempDir, "abc", "pony*.yaml")
	require.NoError(t, err)

	testCases := []struct {
		desc   string
		opts   []pony.Option
		assert func(*testing.T, *pony.Pony, error)
	}{
		{
			desc: "returns error when failed to read metadata file",
			opts: []pony.Option{pony.MetadataFile("abc")},
			assert: func(t *testing.T, p *pony.Pony, err error) {
				assert.EqualError(t, err, "metadata read failed: open abc: no such file or directory")
				assert.False(t, p.MetadataLoaded())
			},
		},
		{
			desc: "returns error when failed to parse metadata file",
			opts: []pony.Option{pony.MetadataFile(fErr)},
			assert: func(t *testing.T, p *pony.Pony, err error) {
				assert.Contains(t, err.Error(), "metadata parse failed: yaml: unmarshal errors:")
				assert.False(t, p.MetadataLoaded())
			},
		},
		{
			desc: "returns no error when loaded metadata file",
			opts: []pony.Option{pony.MetadataFile(fOk)},
			assert: func(t *testing.T, p *pony.Pony, err error) {
				assert.NoError(t, err)
				assert.True(t, p.MetadataLoaded())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := pony.NewPony(tC.opts...)
			assert.False(t, p.MetadataLoaded())

			err := p.LoadMetadata()
			tC.assert(t, p, err)
		})
	}
}

func TestLoadTemplates(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ponytest")
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)

	const tmpl = `
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
  </head>
  <body>
  </body>
</html>`

	_, err = storeTempFile(tempDir, tmpl, "index*.html")
	require.NoError(t, err)

	testCases := []struct {
		desc   string
		opts   []pony.Option
		assert func(*testing.T, *pony.Pony, error)
	}{
		{
			desc: "returns error when failed to parse templates",
			opts: []pony.Option{pony.TemplatesDir("abc")},
			assert: func(t *testing.T, p *pony.Pony, err error) {
				assert.Contains(t, err.Error(), "templates parse failed: ")
				assert.False(t, p.MetadataLoaded())
			},
		},
		{
			desc: "returns error when failed to parse templates",
			opts: []pony.Option{pony.TemplatesDir(tempDir)},
			assert: func(t *testing.T, p *pony.Pony, err error) {
				assert.NoError(t, err)
				assert.True(t, p.TemplatesLoaded())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := pony.NewPony(tC.opts...)
			assert.False(t, p.TemplatesLoaded())

			err := p.LoadTemplates()
			tC.assert(t, p, err)
		})
	}
}
