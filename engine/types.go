package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []interface{}
}

type ParserNewsListTotal struct {
	total int
}

func NilParserFunc(content []byte) ParserResult {
	return ParserResult{}
}