package options

var rhaOptions = &RHAOptions{}

type RHAOptions struct {
}

func GetRHAOptions() *RHAOptions {
	return rhaOptions
}
