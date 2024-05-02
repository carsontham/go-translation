package service_test

import (
	"go-translation/app/infrastructure/service"
	"net/http"
	"testing"
)

type translationClientConfig struct {
	translateEndpoint string
}

// TODO
func setUpTranslationClient(
	t *testing.T,
	config translationClientConfig,
	stubResponses map[string]*http.Response,
) (apiClient *service.Client) {
	return
}
