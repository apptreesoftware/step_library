package dataset

type Configuration struct {
	Name       string
	Success    bool
	Attributes []Attribute
}

type Attribute struct {
	Name       string
	Index      int
	Type       string
	Attributes []Attribute
}

type ConfigInput struct {
	Model Configuration
}
