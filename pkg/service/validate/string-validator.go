package validate

import "regexp"

var NamePattern, _ = regexp.Compile("^[$a-zA-Z][$a-zA-Z0-9_-]*$")
