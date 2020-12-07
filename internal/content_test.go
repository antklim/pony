package internal_test

import (
	"testing"

	"github.com/antklim/pony/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metaData = `
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
    properties: 
      - key: title
        value: About Page
      - key: header
        value: Welcome to the about page`

func TestContentLoad(t *testing.T) {
	expected := &internal.Content{
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
				Name: "About page",
				Path: "/about",
				Properties: []internal.Property{
					{Key: "title", Value: "About Page"},
					{Key: "header", Value: "Welcome to the about page"},
				},
			},
		},
	}

	meta, err := internal.ContentLoad(metaData)
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
			meta, err := internal.ContentLoad(tC.data)
			require.NoError(t, err)
			assert.Equal(t, tC.props, meta.Pages["index"].Props())
		})
	}
}
