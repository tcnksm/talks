import "context"

// START OMIT
type Filter interface {
	Name() string
}

type BeforeFilter interface {
	Filter
	BlacklistedHeaders() []string
	DoBefore(context context.Context) BeforeFilterError
}

type AfterFilter interface {
	Filter
	DoAfter(context context.Context) AfterFilterError
}

// END OMIT
