package minimal

// Route struct for HTTP handlers
type Route struct {
	URI     string
	Action  string
	Handler func() interface{}
	Before  []func() (interface{}, bool, int)
}
