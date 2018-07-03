// GRPCHandler generates handlers which handle gRPC requests which handles given request and response
import (
	"net/http"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

func GRPCHandler(name string, conn *grpc.ClientConn, reqMsg, respMsg proto.Message) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unmarshalRequest(r, reqMsg) // リクエストを与えられたリクエスト型にUnmarshalする

		conn.Invoke(r.Context(), name, reqMsg, respMsg) // gRPCリクエストを実行する

		buf, _ := proto.Marshal(respMsg) // レスポンスをProtobufferにMarshalする

		w.Write(buf) // HTTPレスポンスを書き込む
	})
}

