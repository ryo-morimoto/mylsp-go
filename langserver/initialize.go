package langserver

import (
	"context"

	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

func (h *handler) handleInitialize(
	_ context.Context,
	conn *jsonrpc2.Conn,
	req *jsonrpc2.Request
) (result any, err error) {
	h.conn = conn

	return lsp.InitializeResult {
		Capabilities: lsp.ServerCapabilities{},
	},nil
}