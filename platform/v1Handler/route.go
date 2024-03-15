package v1Handler

import(

	"github.com/gofiber/fiber/v2"

) 


func RegisterInflowFnV1(api fiber.Router) {
	api.Use(recoveryHandler())
	api.Post("/:action/run",actionHandlers2) // run method of IAction
	api.Get("/:action/load",getActionApplication) // get dialog form of action command
	api.Get("",describeHandler) // doc based on Descibe arg on registerAction
	api.Get("/help/md/:name",helpMdHandler)
	api.Get("/help/o/:name",helpObjHandler)

}



func recoveryHandler()func(c  *fiber.Ctx)error{
	return func(c *fiber.Ctx) error {
		defer func(){
			if r := recover(); r != nil {

				if e, ok := r.(error); ok {
					c.SendStatus(400)
					c.JSON("error occurred" ,e.Error())

				} else if v, ok := r.(map[string]error); ok {
					t := ""
					for eKey, errMessage := range v {
						t += eKey + errMessage.Error() + " , "
					}
					c.SendStatus(400)
					c.JSON("error occurred")
				}
				c.SendStatus(500)
				c.JSON("error occurred")

			}

		}()
		return c.Next()
	}

}