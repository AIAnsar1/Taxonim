package Config

import "Taxonim/Core/Types"

func Init() {

}

func NewConfigReader(Config []byte, ConfigType string) (Reader ConfigReader, Err error) {

}

func StepToScenarioStep(Step Step) (Types.ScenarioStep, error) {

}

func PrepareMultipartPayload(Parts []MultipartFormData) (Body string, ContentType string, Err error) {

}

func preparePayloadFile(Url string) (Body string, Err error) {

}
func (this *Step) UnMarshalJson(Data []byte) error {

}

func (this *Tag) UnMarshalJson(Data []byte) error {

}

func (this *CsvConf) UnMarshalJson(Data []byte) error {

}

func (this *JsonReader) UnMarshalJson(Data []byte) error {

}

func (this *JsonReader) Init(JsonByte []byte) (err error) {

}

func (this *JsonReader) CreateHammer() (self Types.Hammer, err error) {

}

func (this RemoteMultipartError) Error() string {

}

func (this RemoteMultipartError) Unwrap() error {

}
