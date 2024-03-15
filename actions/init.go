package actions

import (
	standardActions "github.com/mehdi-shokohi/inflow-fn-std/actions/standard"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func init() {
	inflowV1.RegisterCommand2("null", standardActions.RunNullCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand2("db_settings", standardActions.RunCastCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand2("isnull", standardActions.RunIsNullCommand,standardActions.CastTo{})
	inflowV1.RegisterCommand2("isarray", standardActions.RunIsArrayCommand,standardActions.CastTo{})

}

