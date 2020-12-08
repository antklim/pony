package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metadata = `
template: index.html
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
    template: about.html
    properties: 
      - key: title
        value: About Page
      - key: header
        value: Welcome to the about page`

func TestParseInMeta(t *testing.T) {
	aboutTmpl := "about.html"
	expected := &inMeta{
		Pages: map[string]inPage{
			"index": inPage{
				Name: "Home page",
				Path: "/",
				Properties: []inProperty{
					{Key: "title", Value: "Home Page"},
					{Key: "header", Value: "Welcome to the home page"},
				},
			},
			"about": inPage{
				Name:     "About page",
				Path:     "/about",
				Template: &aboutTmpl,
				Properties: []inProperty{
					{Key: "title", Value: "About Page"},
					{Key: "header", Value: "Welcome to the about page"},
				},
			},
		},
		Template: "index.html",
	}

	meta, err := parseInMeta([]byte(metadata))
	require.NoError(t, err)
	assert.Equal(t, expected, meta)
}

func TestInPageProps(t *testing.T) {
	testCases := []struct {
		desc  string
		data  string
		props Props
	}{
		{
			desc: "returns nil when no properties found",
			data: `
        pages:
          index:
            name: Home page
            path: /`,
			props: nil,
		},
		{
			desc: "returns properties map",
			data: `
        pages:
          index:
            name: Home page
            path: /
            properties:
              - key: title
                value: Home Page
              - key: header
                value: Welcome to the home page`,
			props: map[string]string{
				"title":  "Home Page",
				"header": "Welcome to the home page",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			meta, err := parseInMeta([]byte(tC.data))
			require.NoError(t, err)
			page, ok := meta.Pages["index"]
			assert.True(t, ok)
			assert.Equal(t, tC.props, page.Props())
		})
	}
}
