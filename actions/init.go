package actions

import (
	standardActions "github.com/mehdi-shokohi/inflow-fn-std/actions/standard"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func init() {
	// Commmand Name , Run Function , Arguments Data Model 
	inflowV1.RegisterCommand("null", standardActions.RunNullCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand("db_settings", standardActions.RunCastCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand("isnull", standardActions.RunIsNullCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand("isarray", standardActions.RunIsArrayCommand,standardActions.CastTo{})

}

