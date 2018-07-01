// START OMIT
{
	Name:     "mercari.platform.echo.v1.Echo",
	Endpoint: regionConfig.EchoServiceEndpoint,
	MethodMap: map[string]*gateway.GRPCMethod{
		"/services/echo/say": &gateway.GRPCMethod{  // エンドポイント（Path）
			Name: "Say",                            // gRPCメソッドの名前
			RequestMessage:  &echo_pb.SayRequest{}, // リクエストの型
			ResponseMessage: &echo_pb.SayRequest{}, // レスポンスの型
		},
	},
}
// END OMIT

