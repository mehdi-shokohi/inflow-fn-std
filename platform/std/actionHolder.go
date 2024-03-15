package std

import "github.com/gofiber/fiber/v2"


type ActionHolder struct{
	Run func(*fiber.Ctx)error
	Argument interface{}
}
var actions = map[string]ActionHolder{}
func RegisterCommand2[T any](actionId string,action func(*fiber.Ctx)error,args T){
	if actions ==nil{
		actions = make(map[string]ActionHolder)
	}
	actions[actionId] = ActionHolder{action,args}
}




func GetAction(key string)*ActionHolder{
	if actions==nil{
		actions = make(map[string]ActionHolder)

	}
	if f,ok:= actions[key];ok{
		return &f
	}
	return nil
}
func GetActions()[]string{
	if actions==nil{
		actions = make(map[string]ActionHolder)

	}
	actList:=make([]string,0)
	for  k,_:=range actions{
		actList = append(actList, k)
	}
	return actList
}
