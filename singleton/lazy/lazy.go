package lazy

import "sync"

type Instance struct{}

var (
	i    *Instance
	once = &sync.Once{}
)

func GetInstance() *Instance {
	if i == nil {
		once.Do(func() {
			i = &Instance{}
		})
	}
	return i
}
