package pony_test

import (
	"testing"

	"github.com/antklim/pony"
	"github.com/stretchr/testify/assert"
)

func TestFromPropertyInput(t *testing.T) {
	pi := []pony.PropertyInput{{
		Key:   "title",
		Value: "Hello Page",
	}, {
		Key:   "header",
		Value: "Welcome",
	}}
	expectedProps := pony.Properties(map[string]string{"title": "Hello Page", "header": "Welcome"})
	props := pony.FromPropertyInput(pi)
	assert.Equal(t, expectedProps, props)
}

func TestFromPageInput(t *testing.T) {
	tmpl := "index"
	pi := pony.PageInput{
		Name:     "home",
		Path:     "/",
		Template: &tmpl,
		Properties: []pony.PropertyInput{{
			Key:   "title",
			Value: "Hello Page",
		}, {
			Key:   "header",
			Value: "Welcome",
		}},
	}

	expectedPage := pony.Page{
		ID:         "homePage",
		Name:       "home",
		Path:       "/",
		Template:   &tmpl,
		Properties: pony.Properties(map[string]string{"title": "Hello Page", "header": "Welcome"}),
	}
	page := pony.FromPageInput("homePage", pi)
	assert.Equal(t, expectedPage, page)
}
