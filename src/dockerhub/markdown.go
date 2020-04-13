package dockerhub

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/charmbracelet/glamour"
)

// RenderMarkdown with display markdown string to the terminal in a nicely formatted style
func RenderMarkdown(md string, writer io.Writer) error {
	const WIDTH = 100

	reader := strings.NewReader(md)

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	// Initialize glamour
	var gs glamour.TermRendererOption
	gs = glamour.WithAutoStyle()

	r, err := glamour.NewTermRenderer(
		gs,
		glamour.WithWordWrap(int(WIDTH)),
	)
	if err != nil {
		return err
	}

	out, err := r.RenderBytes(bytes)
	if err != nil {
		return err
	}

	// Trim lines
	lines := strings.Split(string(out), "\n")

	var content string
	for i, s := range lines {
		content += strings.TrimSpace(s)

		// don't add newline after the last split
		if i+1 < len(lines) {
			content += "\n"
		}
	}

	// display
	fmt.Fprint(writer, content)
	return nil
}
