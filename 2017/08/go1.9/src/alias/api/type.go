package api

type ID int

type SpecialError struct {
}

func (*SpecialError) Error() string {
	return "Special error"
}
