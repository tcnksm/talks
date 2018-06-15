// START OMIT
{
	Name:     "mercari.platform.echo.v1.Echo",
	Endpoint: regionConfig.EchoServiceEndpoint,
	MethodMap: map[string]*gateway.GRPCMethod{
		"/services/echo/say": &gateway.GRPCMethod{
			Name: "Say",
			RequestMessage:  &echo_pb.SayRequest{}, // Request message
			ResponseMessage: &echo_pb.SayRequest{}, // Response message
		},
	},
}
// END OMIT

