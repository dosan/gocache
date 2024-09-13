package gocache


type CacheInterface interface{
	Set()
	Get()
	Delete()
}
type Obj map[string]interface{}

func New() Obj{
	obj := make(Obj)
	return obj
}

func (o Obj) Set(key string, value interface{}) {
	o[key] = value
}
func (o Obj) Get(key string) interface{} {
	return o[key]
}
func (o Obj) Delete(key string) {
	delete(o, key)
}
