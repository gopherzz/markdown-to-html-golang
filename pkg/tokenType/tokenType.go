package tokenType

const (
	Text = iota
	H1
	H2
	H3
	H4
	H5
	H6

	Bold
	Italic
	BoldItalic

	Link
	Code

	NewLine
	EOF
)

var TypesString = map[int]string{
	Text: "Text",
	H1: "H1",
	H2: "H2",
	H3: "H3",
	H4: "H4",
	H5: "H5",
	H6: "H6",

	Bold: "Bold",
	Italic: "Italic",
	BoldItalic: "BoldItalic",

	Link: "Link",
	Code: "Code",

	NewLine: "NL",
	EOF: "EOF",
}
