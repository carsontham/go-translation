package service

import (
	"context"
)

type TranslateAPI interface {
	Translate(ctx context.Context, query string, target string) (string, error)
}
