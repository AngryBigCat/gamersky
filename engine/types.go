package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests  []Request
	Items     []interface{}
	StoreFunc func()
}

func NilParserFunc(content []byte) ParserResult {
	return ParserResult{}
}
