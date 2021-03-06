package backend

import (
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type Backend interface {
	Process(iCode intermediate.ICode, symTab intermediate.SymTab) error
	GetMessageHandler() *message.MessageHandler
}

type BaseBackend struct {
	MessageHandler *message.MessageHandler

	symTab *intermediate.SymTab
	iCode  *intermediate.ICode
}

func NewBaseBackend(mh *message.MessageHandler) BaseBackend {
	return BaseBackend{
		MessageHandler: mh,
		symTab:         nil,
		iCode:          nil,
	}
}

func (bb *BaseBackend) GetMessageHandler() *message.MessageHandler {
	return bb.MessageHandler
}
