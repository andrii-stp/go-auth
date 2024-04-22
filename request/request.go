package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

var httpMethods = []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
	http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace}

func NewHTTP[T any](method, rawURL string, headers map[string]string, queryParams url.Values, body io.Reader, response T) (T, error) {
	if !slices.Contains(httpMethods, method) {
		return response, fmt.Errorf("invalid http method")
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return response, err
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Set(k, strings.Join(v, ","))
		}

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return response, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}

	slog.Debug("request sent", slog.String("description", fmt.Sprintf("%s %s", method, req.URL.String())))

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		slog.Error("error unmarshaling response", slog.String("description", err.Error()))
		return response, err
	}

	return response, nil
}
