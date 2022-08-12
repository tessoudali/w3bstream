package vm

type Config struct {
	SpecVersion string       `yaml:"specVersion"`
	RepoAddr    string       `yaml:"repository"`
	Desc        string       `yaml:"description"`
	Schema      Schema       `yaml:"schema"`
	DataSources []DataSource `yaml:"dataSources"`
}

type Schema struct {
	File string `yml:"file"`
}

type DataSource struct {
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
	Network string `yaml:"network"`
	Source  `yaml:"source"`
	Mapping `yaml:"mapping"`
}

type Source struct {
	Address string `yaml:"address"`
	AbiName string `yaml:"abi"`
}

type Mapping struct {
	Kind          string         `yaml:"kind"`
	APIVersion    string         `yaml:"apiVersion"`
	Language      string         `yaml:"language"`
	Entities      []string       `yaml:"entities"`
	ABIs          []ABI          `yaml:"abis"`
	EventHandlers []EventHandler `yaml:"eventHandlers"`
	File          string         `yaml:"file"`
}

type ABI struct {
	Name string `yml:"name"`
	File string `yml:"file"`
}

type EventHandler struct {
	Event   string `yml:"event"`
	Handler string `yml:"handler"`
}
