package engine

type Result struct {
	Requests []Request
	Contents []interface{}
}

func (r Result) NilFunc([]byte) Result {
	return Result{}
}

type Request struct {
	Url        string
	HandleFunc func([]byte) Result
}
