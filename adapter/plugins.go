package adapter

import "context"

type Plugin interface {
	Handler(ctx context.Context) (bool, error)
}
