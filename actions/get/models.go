package standardActions

import "go.mongodb.org/mongo-driver/bson/primitive"

type IsoMapping struct {
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SourceFramework string   `json:"source_framework" bson:"source_framework"`
	SourceControl   []string `json:"source_control" bson:"source_control"`
	TargetFramework string   `json:"target_framework" bson:"target_framework"`
	TargetControl   []string `json:"target_control" bson:"target_control"`
}

type IsoMappingArgs struct{
	SourceFramework string   `json:"source_framework" bson:"source_framework"`
	SourceControl   string `json:"source_control" bson:"source_control"`
	TargetFramework string   `json:"target_framework" bson:"target_framework"`
}