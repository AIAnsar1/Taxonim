package Injection

func (Tbr TaxonimBodyReader) Close() error {

}

func (Tbr TaxonimBodyReader) Read(Dst []byte) (N int, err error) {

}

func (this *EnvironmentInjector) Init() {

}

func TruncateTag(Tag string, Rx string) string {

}

func InjectEnv(Text string, Envs map[string]interface{}) (string, error) {

}

func (this *EnvironmentInjector) GetEnv(Input map[string]interface{}, Output string) (interface{}, error) {

}

func UnifyErrors(errors []error) error {

}

func StringToBytes(S string) (B []byte) {

}

func GetInjectStrFunc(Rx string, self *EnvironmentInjector, Input map[string]interface{}, errors *[]error) func(string) string {
	return func(string) string {

	}
}

func GetInjectJsonFunc(Rx string, self *EnvironmentInjector, Input map[string]interface{}, errors *[]error) func(S []byte) []byte {
	return func(S []byte) []byte {

	}
}
func (this EnvMatchSlice) Len() int {

}

func (this EnvMatchSlice) Swap(I int, J int) {

}
func (this EnvMatchSlice) Less(I int, J int) bool {

}

func (this *EnvironmentInjector) GenerateBodyPieces(Body string, Input map[string]interface{}) []BodyPiece {

}

func GetContentLength(this []BodyPiece) int {

}
