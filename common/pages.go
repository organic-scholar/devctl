package common

import (
	"io/fs"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/aymerick/raymond"
)

func ParsePages(f fs.FS, patterns ...string) (map[string]*raymond.Template, error) {
	tmpl := make(map[string]*raymond.Template)
	matches := []string{}
	for _, pattern := range patterns {
		next, err := fs.Glob(f, pattern)
		if err != nil {
			return tmpl, err
		}
		matches = slices.Concat(matches, next)
	}
	for _, match := range matches {
		content, err := fs.ReadFile(f, match)
		if err != nil {
			return tmpl, err
		}
		name := strings.TrimPrefix(match, "pages/")
		tpl, err := raymond.Parse(string(content))
		tmpl[name] = tpl
		if err != nil {
			return tmpl, err
		}
	}
	return tmpl, nil
}

func (m *PageRouter) RegisterHelper() {
	raymond.RegisterHelper("block", func(name string, options *raymond.Options) string {
		out := options.DataFrame().Get("_content_" + name)
		if out == nil {
			return options.Fn()
		}
		return out.(string)
	})
	raymond.RegisterHelper("layout", func(name string, options *raymond.Options) string {
		options.Fn()
		tpl, err := raymond.ParseFile("pages/" + name)
		if err != nil {
			panic(err)
		}
		out, err := tpl.ExecWith(options.Ctx(), options.DataFrame())
		if err != nil {
			panic(err)
		}
		return out
	})
	raymond.RegisterHelper("content", func(name string, options *raymond.Options) string {
		content := options.Fn()
		options.DataFrame().Set("_content_"+name, content)
		return ""
	})
}

func (m *PageRouter) RenderPage(w http.ResponseWriter, name string, data any) error {
	content, err := os.ReadFile("pages/" + name)
	if err != nil {
		panic(err)
	}
	res, err := raymond.Render(string(content), data)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(res))
	return err
}
