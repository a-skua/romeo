package romeo

// ResponseValueWrapper is the method
// that wraps argument for response body
type ResponseValueWrapper func(interface{}) interface{}

type defaultValue struct {
	Data interface{} `json:"data"`
}

// DefaultWrap implements  ResponseValueWrapper
func DefaultWrap(data interface{}) interface{} {
	return &defaultValue{data}
}

// NoWrap implements ResponseValueWrapper
// that not wraps
func NoWrap(data interface{}) interface{} {
	return data
}
