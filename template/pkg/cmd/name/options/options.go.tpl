package options

type Option struct {
	Opt1 string
	Opt2 string
	Opt3 string
}

func NewOption() *Option {
	return &Option{}
}
