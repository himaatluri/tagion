package types

type Tags struct {
	Key   string `json:"Key" yaml:"Key"`
	Value string `json:"Value" yaml:"Value"`
}

type Config struct {
	Tags map[string]string `json:"tags"`
}

type CloudFormationTemplate struct {
	Resources map[string]Resource `json:"Resources" yaml:"Resources"`
}

type Resource struct {
	Type       string                 `json:"Type" yaml:"Type"`
	Properties map[string]interface{} `json:"Properties" yaml:"Properties"`
}
