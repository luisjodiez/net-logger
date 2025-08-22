package core

type Result struct {
	Target    string
	ConnType  string
	Timestamp int64
	Datetime  string
	Status    string // ok/ko
}

type ResultsRepository interface {
	Add(Result)
	All() []Result
}

type Pinger interface {
	Ping(target string) string
}

type Reporter interface {
	Report(results []Result) error
}
