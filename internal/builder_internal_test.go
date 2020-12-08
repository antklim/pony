package internal

/*
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
		page         Page
		expectedTmpl string
	}{
		{
			desc: "builds page to the root directory",
			page: Page{
				ID:         "index",
				Name:       "Pony root",
				Path:       "/",
				Properties: map[string]string{"foo": "root"},
			},
			expectedTmpl: "<title>root</title>",
		},
		{
			desc: "builds page to the sub-directory",
			page: Page{
				ID:         "test",
				Name:       "Pony test",
				Path:       "/test",
				Properties: map[string]string{"foo": "test"},
			},
			expectedTmpl: "<title>test</title>",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := buildPage(tC.page, tmpl, dir)
			require.NoError(t, err)

			fname := filepath.Join(dir, tC.page.Path, tC.page.ID+".html")
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
*/
