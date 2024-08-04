package Types

import (
	"crypto/tls"
	"crypto/x509"
)

func init() {

}

func (this *Scenario) Validate() error {

}

func CheckEnvsValidInStep(this *ScenarioStep, DefinedEnv map[string]interface{}) error {

}

func (this *ScenarioStep) Validate(DefinedEnv map[string]interface{}) error {

}

func WrapAsScenarioValidationError(err error) ScenarioValidationError {

}

func validateCaptureConf(this EnvCaptureConf) error {

}

func ParseTLS(CertFile string, KeyFile string) (tls.Certificate, *x509.CertPool, error) {

}

func IsTargetValid(Url string) error {

}
