package inflowV1

import (

	"github.com/gofiber/fiber/v2"

)
type IAction interface{
	Run(*fiber.Ctx)(error)
	// Arguments(*fiber.Ctx)(error)
	Settings(*fiber.Ctx)(error)
}


//command for some action

//access to db 
//find db by name 
//restart policy on doc
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

func GetRAWBodyAs[T any](ctx *fiber.Ctx)(*T,error){
	inputForm := new(T)
	if err := ctx.BodyParser(&inputForm); err != nil {
		return nil,err
	}

	return inputForm,nil
}