package hungry

type Instance struct{}

var i *Instance

func init() {
	i = &Instance{}
}

func GetInstance() *Instance {
	return i
}
