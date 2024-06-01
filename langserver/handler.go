package langserver

import (
	"context"
	"errors"

	"github.com/sourcegraph/jsonrpc2"
)

type handler struct{}

func NewHandler() jsonrpc2.Handler {
	h := &handler{}
	return jsonrpc2.HandlerWithError(h.handle)
}

func (h *handler) handle(
	ctx context.Context,
	conn *jsonrpc2.Conn,
	req *jsonrpc2.Request,
) (result any, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	switch req.Method {
	case "initialize":
		return h.handleInitialize(ctx, conn, req)
	case "initialized":
		return nil, nil
	}
	return nil, errors.New("not implemented")
}
