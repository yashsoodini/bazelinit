package template

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGenerateOutputs(t *testing.T) {
	testCases := []struct {
		name          string
		fs            fs.FS
		templatesRoot string
		templateData  any
		want          map[string][]byte
		shouldPanic   bool
	}{
		{
			name: "simpleCaseWithFlatStructure",
			fs: fstest.MapFS{
				"templates/file.txt":      {Data: []byte("plain file")},
				"templates/hello.tmpl":    {Data: []byte("hello, {{.Name}}")},
				"templates/another.jsonc": {Data: []byte(`{"foo":"bar"}`)},
			},
			templatesRoot: "templates",
			templateData:  struct{ Name string }{"World"},
			want: map[string][]byte{
				"file.txt":      []byte("plain file"),
				"hello":         []byte("hello, World"),
				"another.jsonc": []byte(`{"foo":"bar"}`),
			},
			shouldPanic: false,
		},
		{
			name: "nestedDirectories",
			fs: fstest.MapFS{
				"project/src/main.go":          {Data: []byte("package main")},
				"project/src/helper/util.tmpl": {Data: []byte("package {{.Pkg}}")},
			},
			templatesRoot: "project",
			templateData:  struct{ Pkg string }{"helper"},
			want: map[string][]byte{
				"src/main.go":     []byte("package main"),
				"src/helper/util": []byte("package helper"),
			},
			shouldPanic: false,
		},
		{
			name:          "emptyFilesystem",
			fs:            fstest.MapFS{},
			templatesRoot: ".",
			templateData:  nil,
			want:          map[string][]byte{},
			shouldPanic:   false,
		},
		{
			name: "invalidTemplateSyntax",
			fs: fstest.MapFS{
				"templates/invalid.tmpl": {Data: []byte("hello, {{.Name")},
			},
			templatesRoot: "templates",
			templateData:  nil,
			want:          nil,
			shouldPanic:   true,
		},
		{
			name: "templateExecutionError",
			fs: fstest.MapFS{
				"templates/exec_error.tmpl": {Data: []byte("hello, {{.NonExistent}}")},
			},
			templatesRoot: "templates",
			templateData:  struct{ Name string }{"World"},
			want:          nil,
			shouldPanic:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("GenerateOutputs should have panicked but did not")
					}
				}()
			}

			got := GenerateOutputs(tc.fs, tc.templatesRoot, tc.templateData)

			if !tc.shouldPanic {
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("GenerateOutputs() got = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
