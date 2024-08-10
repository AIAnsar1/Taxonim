package Data

import "Taxonim/Core/Types"

func ValidateConf(conf Types.CsvConf) error {

}

func ReadCsv(conf Types.CsvConf) ([]map[string]interface{}, error) {

}

func EmptyLine(Row []string) bool {

}

func WrapAsCsvError(Message string, err error) RemoteCsvError {

}
func (this RemoteCsvError) Error() string {

}

func (this RemoteCsvError) Unwrap() error {

}
