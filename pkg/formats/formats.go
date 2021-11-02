package formats

import "MTDGo/pkg/tokenType"

var Format = map[int]string {
	tokenType.Text: "%s",
	tokenType.NewLine: "<br>",
	tokenType.H1: "<h1>%s</h1>",
	tokenType.H2: "<h2>%s</h2>",
	tokenType.H3: "<h3>%s</h3>",
	tokenType.H4: "<h4>%s</h4>",
	tokenType.H5: "<h5>%s</h5>",
	tokenType.H6: "<h6>%s</h6>",
	tokenType.Bold: "<b>%s</b>",
	tokenType.Italic: "<i>%s</i>",
	tokenType.BoldItalic: "<b><i>%s</i></b>",
}

