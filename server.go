package hw3

import (
	"net/http"
)

type RemoteMethodStubProvider interface {
	CreateObjectStub(serverObject any) http.HandlerFunc
}

var _ RemoteMethodStubProvider = (*RmiStubProvider)(nil)

type RmiStubProvider struct {
}

func NewRmiStubProvider() *RmiStubProvider

func (r *RmiStubProvider) CreateObjectStub(serverObject any) http.HandlerFunc
