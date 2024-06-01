package langserver

import (
	"context"
	"errors"

	"github.com/sourcegraph/jsonrpc2"
)

type handler struct {
	conn  *jsonrpc2.conn
	cache *cache
}

func NewHandler() jsonrpc2.Handler {
	h := &handler{
		cache: newCache(),
	}
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
	case "textDocument/didOpen":
		return h.handleTextDocumentDidOpen(ctx, conn, req)
	case "textDocument/didChange":
		return h.handleTextDocumentDidChange(ctx, conn, req)
	case "textDocument/didClose":
		return h.handleTextDocumentDidClose(ctx, conn, req)
	case "textDocument/didSave":
		return h.handleTextDocumentDidSave(ctx, conn, req)
	}
	return nil, errors.New("not implemented")
}
