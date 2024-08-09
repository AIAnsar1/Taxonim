package Evaluator

import "time"

var (
	LessChan = func(Variable int64, Limit int64) bool {

	}

	GreaterThan = func(Variable int64, Limit int64) bool {

	}

	Not = func(b bool) bool {

	}

	Percentile = func(Arr []int64, Num int64) (int64, error) {

	}

	Min = func(Arr []int64, Num int64) (int64, error) {

	}

	Max = func(Arr []int64, Num int64) (int64, error) {

	}

	Avg = func(Arr []float64, Num int64) (int64, error) {

	}

	Equals = func(A interface{}, B interface{}) (bool, error) {

	}

	In = func(A interface{}, B interface{}) (bool, error) {

	}

	Contains = func(Source string, SubStr string) bool {

	}

	Time = func(time string) (time.Time, error) {

	}

	Ranger = func(X float64, Low float64, Hi float64) bool {

	}

	JsonExtract = func(Source interface{}, JPath string) (interface{}, string) {

	}

	XmlExtract = func(Source interface{}, XPath string) (interface{}, string) {

	}

	HtmlExtract = func(Source interface{}, HPath string) (interface{}, string) {

	}

	RegexExtract = func(Source interface{}, RPath string, MatchNo int64) (interface{}, string) {

	}


	EqualsOnFile = func(Source interface{}. FPath string) (bool, error) {
		
	}
)
