package actions

import (
	standardActions "github.com/mehdi-shokohi/inflow-fn-std/actions/standard"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func init() {
	// Commmand Name , Run Function , Arguments Data Model
	inflowV1.RegisterCommand("null", standardActions.RunNullCommand, standardActions.Keys{Keys: []string{"k1", "k2"}})
	inflowV1.RegisterCommand("delete", standardActions.RunDeleteCommand, standardActions.DeleteKeys{Keys: []string{"k1", "k2"}})
	inflowV1.RegisterCommand("cast", standardActions.RunCastCommand, standardActions.CastTo{CastTo: "int|string", Keys: []string{"k1", "k2"}})
	inflowV1.RegisterCommand("isnull", standardActions.RunIsNullCommand, standardActions.IsNull{})
	inflowV1.RegisterCommand("isarray", standardActions.RunIsArrayCommand, standardActions.IsArray{})
	inflowV1.RegisterCommand("timer", standardActions.RunTimerCommand, standardActions.Scheduler{})
	inflowV1.RegisterCommand("http_request", standardActions.HttpRequest, standardActions.HTTPRequest{URL: "http://root:mate123@localhost:8529/_api/cursor",Method: "POST",Body: map[string]interface{}{"query" : "FOR p IN items filter p.terminologyId in ['ICPC2P','CPT'] limit 12000,1000 RETURN p "	}})
	inflowV1.RegisterCommand("db_settings", standardActions.RunCastCommand, standardActions.DBSettings{})
}
