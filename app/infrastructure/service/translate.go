package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-translation/app/httpapi"
	"io"
	"net/http"
)

// Translate ...
func (c *Client) Translate(ctx context.Context, s string, target string) (string, error) {

	translateForm := TranslateRequest{
		Text:   s,
		Target: target,
	}

	api := &translateAPI{
		endpoint: c.translateEndpoint,
		decorReq: httpapi.DecorateRequest(
			httpapi.ContentType,
			httpapi.RapidAPIKey,
			httpapi.RapidAPIHost,
			httpapi.AcceptEncoding,
		),
		req: translateForm,
	}

	fmt.Println("sending translation request...")
	successful, err := c.httpClient.Call(ctx, api)
	if successful {
		fmt.Println("translation success")
		return api.resp.Data.Translations.TranslatedText, nil
	} else {
		return "", err
	}
}

var _ httpapi.API = new(translateAPI)

type translateAPI struct {
	endpoint string
	decorReq httpapi.DecorateRequestFunc
	req      TranslateRequest
	resp     *TranslateResponse
}

func (api *translateAPI) BuildRequest(ctx context.Context) (*http.Request, error) {
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, api.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return api.decorReq(req)
}

func (api *translateAPI) ParseResponse(_ context.Context, _ *http.Request, resp *http.Response) error {
	var respData TranslateResponse
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, statusCode: %d, body: %s", errors.New("status not ok"), resp.StatusCode, resp.Body)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &respData); err != nil {
		return err
	}

	api.resp = &respData
	return nil
}
