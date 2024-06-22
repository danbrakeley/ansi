package sgr

// SGR stands for Select Graphic Rendition, and is used to alter display attributes.
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
type SGR string

const (
	Reset         SGR = "0"
	Bold          SGR = "1"
	Underline     SGR = "4"
	Reverse       SGR = "7"
	UnderlineOff  SGR = "24"
	ReverseOff    SGR = "27"
	FgBlack       SGR = "30"
	FgDarkRed     SGR = "31"
	FgDarkGreen   SGR = "32"
	FgDarkYellow  SGR = "33"
	FgDarkBlue    SGR = "34"
	FgDarkMagenta SGR = "35"
	FgDarkCyan    SGR = "36"
	FgLightGray   SGR = "37"
	FgReset       SGR = "39" // sets foreground color to default
	BgBlack       SGR = "40"
	BgDarkRed     SGR = "41"
	BgDarkGreen   SGR = "42"
	BgDarkYellow  SGR = "43"
	BgDarkBlue    SGR = "44"
	BgDarkMagenta SGR = "45"
	BgDarkCyan    SGR = "46"
	BgLightGray   SGR = "47"
	BgReset       SGR = "49" // sets background color to default
	FgDarkGray    SGR = "90"
	FgRed         SGR = "91"
	FgGreen       SGR = "92"
	FgYellow      SGR = "93"
	FgBlue        SGR = "94"
	FgMagenta     SGR = "95"
	FgCyan        SGR = "96"
	FgWhite       SGR = "97"
	BgDarkGray    SGR = "100"
	BgRed         SGR = "101"
	BgGreen       SGR = "102"
	BgYellow      SGR = "103"
	BgBlue        SGR = "104"
	BgMagenta     SGR = "105"
	BgCyan        SGR = "106"
	BgWhite       SGR = "107"
)
