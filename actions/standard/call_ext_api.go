package standardActions

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"strings"

	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"

	inflowV1 "github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func HttpRequest(c *fiber.Ctx) error {

	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	bodyByte, _ := json.Marshal(body.Body)

	if v, ok := body.InlineParams["with_body"]; ok {
		if v != nil {
			bodyByte, err = json.Marshal(body.InlineParams["with_body"])
			if err != nil {
				panic(err)
			}

		}
	}
	headers := c.GetReqHeaders()
	response, err := Request(strings.ToUpper(body.InlineParams["method"].(string)), body.InlineParams["call_url"].(string), bodyByte, headers)
	if err != nil || response == nil {
		panic(err)
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(response, &resp)
	if err != nil {
		panic(err)
	}
	return inflowV1.Send(c, fiber.StatusOK, resp)
}

func Request(method, requestURL string, data []byte, headers map[string][]string) ([]byte, error) {
	bodyReader := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err

	}
	defer res.Body.Close()
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		reader = res.Body
	}

	resBody, err := io.ReadAll(reader)
	if err != nil {
		return nil, err

	}

	return resBody, nil
}
