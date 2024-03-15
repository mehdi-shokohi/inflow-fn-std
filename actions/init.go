package actions

import (
	standardActions "github.com/mehdi-shokohi/inflow-fn-std/actions/standard"
	settingsHandler "github.com/mehdi-shokohi/inflow-fn-std/actions/settings"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)

func init() {
	std.RegisterCommand("null", &standardActions.StandardFunc{}, std.Describe{})
	std.RegisterCommand("cast", &standardActions.StandardFunc{}, std.Describe{})
	std.RegisterCommand("settings", &settingsHandler.PersistSettings{}, std.Describe{})
	std.RegisterCommand("isnull", &standardActions.StandardFunc{}, std.Describe{})
	std.RegisterCommand("isarray", &standardActions.StandardFunc{}, std.Describe{})
	std.RegisterCommand("timer", &standardActions.StandardFunc{}, std.Describe{})


}
