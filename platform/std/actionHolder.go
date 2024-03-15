package std

import (

)


var actions map[string]IAction


func RegisterCommand(actionId string,action IAction,defineHelp Describe){
	if actions==nil{
		actions = make(map[string]IAction)
	}
	// validate action name
	//save help on file
	actions[actionId]=action

}

func GetAction(key string)IAction{
	if actions==nil{
		actions = make(map[string]IAction)

	}
	if f,ok:= actions[key];ok{
		return f
	}
	return nil
}
func GetActions()[]string{
	if actions==nil{
		actions = make(map[string]IAction)

	}
	actList:=make([]string,0)
	for  k,_:=range actions{
		actList = append(actList, k)
	}
	return actList
}