package Regex

const (
	DynamicVariableRegex     = `\{{(_)[^}]+\}}`
	JsonDynamicVariableRegex = `\"{{(_)[^}]+\}}"`
	EnvironmentVariableRegex = `{{[a-zA-Z$][a-zA-Z0-9_().-]*}}`
	JsonEnvironmentVarRegex  = `\"{{[a-zA-Z$][a-zA-Z0-9_().-]*}}"`
)
