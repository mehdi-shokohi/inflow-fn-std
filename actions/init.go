package actions

import (
	standardActions "github.com/mehdi-shokohi/inflow-fn-std/actions/standard"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)

func init() {
	std.RegisterCommand2("null", standardActions.RunNullCommand,standardActions.CastTo{})
	std.RegisterCommand2("global_settings", standardActions.RunCastCommand,standardActions.CastTo{})
	std.RegisterCommand2("isnull", standardActions.RunIsNullCommand,standardActions.CastTo{})
	std.RegisterCommand2("isarray", standardActions.RunIsArrayCommand,standardActions.CastTo{})

}

