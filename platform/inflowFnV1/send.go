package inflowV1

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)
const (
	defaultPage    = 1
	defaultPerPage = 5
)
// type Response struct {
// 	Data  interface{} `json:"data"`
// 	Error interface{} `json:"error"`
// }

// type Error struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

type ValidationError struct {
	Field string
	Rule  string
	Param string
	Message string
}

type PaginationParams struct {
	Page    uint   `query:"page"`
	PerPage uint   `query:"per_page"`
	Search  string `query:"search"`
}
func (i *PaginationParams) Validate() error {
	i.Search = strings.Trim(i.Search," ")
	if i.Page == 0 {
		i.Page = defaultPage
	}

	if i.PerPage == 0 {
		i.PerPage = defaultPerPage
	}
	if i.Search!="" && len(i.Search)<5{
		return errors.New("bad request")
	}
	return nil
}


func Send(c *fiber.Ctx, Code int, Data interface{}) error {
	root := map[string]interface{}{}
	if Data!=nil{
		root["_data"] = Data
	}
	return c.Status(Code).JSON(root)
}