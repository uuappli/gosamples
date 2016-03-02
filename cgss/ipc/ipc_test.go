package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{method, params}
}

func (server *EchoServer) Name() string {
	return "EhcoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("From", "client1")
	resp2, _ := client2.Call("From", "client2")
	if resp1 == nil || resp2 == nil {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
