package infrastructure

type Container struct {
	Config *Configuration
}

var C *Container

func init() {
	conf = NewConfiguration("conf", "sys")

	C = &Container{
		Config: conf,
	}
}
