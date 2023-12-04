package obj

type CustomObject interface {
	Insert(...interface{})
	List() []interface{}
}
