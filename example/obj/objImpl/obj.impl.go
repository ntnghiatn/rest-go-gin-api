package objimpl

import "github.com/ntnghiatn/rest-go-gin-api/example/obj"

type CustomObjectImpl struct {
	data []interface{}
}

// Insert implements obj.CustomObject.
func (da *CustomObjectImpl) Insert(element ...interface{}) {
	// panic("unimplemented")
	da.data = append(da.data, element...)
}

// List implements obj.CustomObject.
func (da *CustomObjectImpl) List() []interface{} {
	// panic("unimplemented")
	return da.data
}

func NewCustomObject() obj.CustomObject {
	return &CustomObjectImpl{
		data: []interface{}{},
	}
}
