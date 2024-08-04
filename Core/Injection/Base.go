package Injection

import (
	"regexp"
	"sync"
)

type BodyPiece struct {
	Start      int
	End        int
	Injectable bool
	Value      string
}

type TaxonimBodyReader struct {
	Body   string
	Pieces []BodyPiece
	Pi     int
	Vi     int
}

type EnvironmentInjector struct {
	R   *regexp.Regexp
	JR  *regexp.Regexp
	DR  *regexp.Regexp
	JDR *regexp.Regexp
	MU  sync.Mutex
}

type EnvMatch struct {
	Regex string
	Found []int
}

type EnvMatchSlice []EnvMatch
