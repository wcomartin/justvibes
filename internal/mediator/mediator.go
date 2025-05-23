package mediator

import "fmt"

type Request interface{}

type Response interface{}

type Handler interface {
    Handle(Request) (Response, error)
}

type Mediator interface {
    Send(Request) (Response, error)
}

type SimpleMediator struct {
    handlers map[string]Handler
}

func NewSimpleMediator() *SimpleMediator {
    return &SimpleMediator{
        handlers: make(map[string]Handler),
    }
}

func (m *SimpleMediator) Register(requestName string, handler Handler) {
    m.handlers[requestName] = handler
}

func (m *SimpleMediator) Send(req Request) (Response, error) {
    reqName := fmt.Sprintf("%T", req)
    handler, ok := m.handlers[reqName]
    if !ok {
        return nil, fmt.Errorf("no handler found for %s", reqName)
    }
    return handler.Handle(req)
}