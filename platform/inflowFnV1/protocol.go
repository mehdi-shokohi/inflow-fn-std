package inflowV1

import (

	"github.com/gofiber/fiber/v2"

)
func GetBodyAs[T any](ctx *fiber.Ctx)(*ProtocolBodyV1[T],error){
	inputForm := ProtocolBodyV1[T]{}
	if err := ctx.BodyParser(&inputForm); err != nil {
		return nil,err
	}

	return &inputForm,nil
}

func GetHeadersAs[T any](ctx *fiber.Ctx)(*T,error){
	inputForm := ProtocolHeaderV1[T]{}
	if err := ctx.BodyParser(&inputForm); err != nil {
		return nil,err
	}

	return &inputForm.Headers,nil
}
