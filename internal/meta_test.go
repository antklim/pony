package internal_test

import (
	"testing"

	"github.com/antklim/pony/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metaData = `
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

func TestMetaLoad(t *testing.T) {
	expected := &internal.Meta{
		Pages: map[string]internal.Page{
			"index": internal.Page{
				Name: "Home page",
				Path: "/",
				Properties: []internal.Property{
					{Key: "title", Value: "Home Page"},
					{Key: "header", Value: "Welcome to the home page"},
				},
			},
			"about": internal.Page{
				Name:     "About page",
				Path:     "/about",
				Template: "about.html",
				Properties: []internal.Property{
					{Key: "title", Value: "About Page"},
					{Key: "header", Value: "Welcome to the about page"},
				},
			},
		},
		Template: "index.html",
	}

	meta, err := internal.MetaLoad(metaData)
	require.NoError(t, err)
	assert.Equal(t, expected, meta)
}

func TestMetaPageProps(t *testing.T) {
	testCases := []struct {
		desc  string
		data  string
		props internal.Props
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
			meta, err := internal.MetaLoad(tC.data)
			require.NoError(t, err)
			assert.Equal(t, tC.props, meta.Pages["index"].Props())
		})
	}
}
