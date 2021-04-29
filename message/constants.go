package messagepacker

import "regexp"

const (
	EncodingChar   = "CHAR" // for Number
	EncodingHex    = "HEX"
	EncodingEbcdic = "EBCDIC"
	EncodingAscii  = "ASCII" // for characters
	EncodingBcd    = "BCD"   // packed bcd
	EncodingRBcd   = "RBCD"

	EncodingCatNumber    = "number"
	EncodingCatBinary    = "binary"
	EncodingCatCharacter = "character"

	HeaderSize         int = 4
	MtiSize            int = 4
	HeaderHexFormatter     = "%04x"
	HexHeaderSize      int = 4
	HexMtiSize         int = 8
)

var (
	RegexAlphabetic          = regexp.MustCompile(`^[a-z A-Z]+$`).MatchString
	RegexNumeric             = regexp.MustCompile(`^[0-9]+$`).MatchString
	RegexSpecial             = regexp.MustCompile(`^[$&+,:;=?@#|'<>.^*()%! -]+$`).MatchString
	RegexIndicate            = regexp.MustCompile(`^[C|D]{1}$`).MatchString
	RegexAlphaNumeric        = regexp.MustCompile(`^[a-z A-Z0-9]*$`).MatchString
	RegexIndicateNumeric     = regexp.MustCompile(`^[C|D][0-9]+$`).MatchString
	RegexAlphaSpecial        = regexp.MustCompile(`^[a-zA-Z$&+,:;=?@#|'<>.^*()%! -]+$`).MatchString
	RegexBinary              = regexp.MustCompile(`^[0|1]+$`).MatchString
	RegexNumericSpecial      = regexp.MustCompile(`^[0-9$&+,:;=?@#|'<>.^*()%! -]+$`).MatchString
	RegexAlphaNumericSpecial = regexp.MustCompile(`^[0-9a-zA-Z$&+,:;=?@#|'<>.^*()%! -]+$`).MatchString
	RegexMagnetic            = regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString

	RegexTimeHHMMSS     = regexp.MustCompile(`(2[0-3]|[01][0-9])[0-5][0-9][0-5][0-9]`).MatchString
	RegexDateYYMM       = regexp.MustCompile(`((\d{2})(0[1-9]|10|12))`).MatchString
	RegexDateMMDD       = regexp.MustCompile(`((0[1-9]|10|11|12)(0[1-9]|[12][0-9]|3[01]))`).MatchString
	RegexDateYYMMDD     = regexp.MustCompile(`((\d{2})(0[1-9]|10|11|12)(0[1-9]|[12][0-9]|3[01]))`).MatchString
	RegexDateMMDDHHMMSS = regexp.MustCompile(`((0[1-9]|10|11|12)(0[1-9]|[12][0-9]|3[01])(2[0-3]|[01][0-9])[0-5][0-9][0-5][0-9])`).MatchString
)
