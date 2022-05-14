//Package markdown is responsible for rendering the contents of a
//given CHANGELOG.md file and displaying it in the terminal.
package markdown

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

func Render(path string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithEmoji(),
		glamour.WithWordWrap(140), // TODO: make this configurable
	)

	data, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return errors.New("changelog not found. Check your configuration or run gh changelog new")
	}

	content, err := r.Render(string(data))
	if err != nil {
		return err
	}

	return start(content)
}
