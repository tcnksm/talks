package main

// START OMIT
type ClientTrace struct {
	// GetConn is called before a connection is created.
	GetConn func(hostPort string)

	// GotConn is called after a successful connection is　obtained.
	GotConn func(GotConnInfo)

	// WroteHeaders is called after the Transport has written
	// the request headers.
	WroteHeaders func()

	// WroteRequest is called with the result of writing the
	// request and any body.
	WroteRequest func(WroteRequestInfo)

	// And more
}

// END OMIT
