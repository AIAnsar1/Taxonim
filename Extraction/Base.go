package Extraction

import "regexp"

type ExtractionError struct {
	Message      string
	WrappedError error
}
type RegexExtractor struct {
	R *regexp.Regexp
}
type HtmlExtractor struct {
}

type JsonExtractor struct {
}
type XmlExtractor struct {
}
