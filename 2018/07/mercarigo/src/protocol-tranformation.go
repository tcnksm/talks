// GRPCHandler generates handlers which handle gRPC requests which handles given request and response
import (
	"net/http"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

func GRPCHandler(fullName string, conn *grpc.ClientConn, req, resp proto.Message) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unmarshalRequest(r, req)

		conn.Invoke(r.Context(),
			fullName, req, resp,
			grpc.FailFast(false),
		)

		buf, _ := proto.Marshal(resp)
		w.Write(buf)
	})
}

