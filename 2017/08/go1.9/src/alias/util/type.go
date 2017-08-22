package util

// type SpecialError api.SpecialError

// type SpecialError struct {
// 	api.SpecialError
// }

type SpecialError struct {
}

func (*SpecialError) Error() string {
	return "Special error"
}
