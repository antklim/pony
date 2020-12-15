package pony_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antklim/pony"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	metadata = `
template: %s
pages:
  index:
    name: Home page
    path: /
    properties: 
      - key: title
        value: Home Page
      - key: header
        value: Welcome to the home page`

	tmpl = `
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
  </head>
  <body>
  </body>
</html>`
)

func setupFiles() (tempDir, fMeta, tmplName string, err error) {
	tempDir, err = ioutil.TempDir("", "ponytest")
	if err != nil {
		return tempDir, fMeta, tmplName, err
	}

	tmplPath, err := storeTempFile(tempDir, tmpl, "index*.html")
	if err != nil {
		return tempDir, fMeta, tmplName, err
	}

	tmplFile := filepath.Base(tmplPath)
	tmplName = strings.TrimSuffix(tmplFile, filepath.Ext(tmplFile))

	fMeta, err = storeTempFile(tempDir, fmt.Sprintf(metadata, tmplName), "pony*.yaml")
	if err != nil {
		return tempDir, fMeta, tmplName, err
	}

	return tempDir, fMeta, tmplName, err
}

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

func TestLoadAll(t *testing.T) {
	tempDir, fMeta, _, err := setupFiles()
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)

	testCases := []struct {
		desc   string
		opts   []pony.Option
		assert func(*testing.T, []error)
	}{
		{
			desc: "returns a list of errors when failed",
			opts: []pony.Option{
				pony.MetadataFile("abc"),
				pony.TemplatesDir("xyz"),
			},
			assert: func(t *testing.T, errs []error) {
				assert.Len(t, errs, 2)
			},
		},
		{
			desc: "returns nil when succeed",
			opts: []pony.Option{
				pony.MetadataFile(fMeta),
				pony.TemplatesDir(tempDir),
			},
			assert: func(t *testing.T, errs []error) {
				assert.Nil(t, errs)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := pony.NewPony(tC.opts...)
			errs := p.LoadAll()
			tC.assert(t, errs)
		})
	}
}

func TestLoadMetadata(t *testing.T) {
	tempDir, fOk, _, err := setupFiles()
	defer os.RemoveAll(tempDir)
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
	tempDir, _, _, err := setupFiles()
	defer os.RemoveAll(tempDir)
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

func TestRenderPages(t *testing.T) {
	tempDir, fMeta, _, err := setupFiles()
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)

	opts := []pony.Option{
		pony.MetadataFile(fMeta),
		pony.TemplatesDir(tempDir),
	}
	p := pony.NewPony(opts...)
	errs := p.LoadAll()
	require.Nil(t, errs)

	bufs := make(map[string]bytes.Buffer)
	pageWriter := func(id, path string) (io.Writer, error) {
		var buf bytes.Buffer
		bufs[id] = buf
		return &buf, nil
	}

	err = p.RenderPages(pageWriter)
	require.NoError(t, err)
	assert.Len(t, bufs, 1)

	_, ok := bufs["index"]
	assert.True(t, ok)

	_, nok := bufs["index1"]
	assert.False(t, nok)
}

func TestRenderPage(t *testing.T) {
	tempDir, fMeta, tmplName, err := setupFiles()
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)

	page := pony.Page{
		ID:         "index",
		Name:       "Home Page",
		Template:   &tmplName,
		Properties: map[string]string{"title": "Foo Bar"},
	}

	opts := []pony.Option{
		pony.MetadataFile(fMeta),
		pony.TemplatesDir(tempDir),
	}
	p := pony.NewPony(opts...)
	errs := p.LoadAll()
	require.Nil(t, errs)

	var buf bytes.Buffer
	err = p.RenderPage(page, &buf)
	require.NoError(t, err)

	renderedPage := `
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Foo Bar</title>
  </head>
  <body>
  </body>
</html>`

	assert.Equal(t, renderedPage, buf.String())
}
