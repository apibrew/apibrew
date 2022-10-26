package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func UserToRecord(user *model.User) *model.Record {
	properties := make(map[string]interface{})

	properties["username"] = user.Username
	properties["password"] = user.Password
	if user.Details != nil {
		properties["details"] = user.Details.AsMap()
	}
	properties["scopes"] = strings.Join(user.Scopes, ",")

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         user.Id,
		Resource:   system.UserResource.Name,
		Type:       user.Type,
		Properties: structProperties,
		AuditData:  user.AuditData,
		Version:    user.Version,
	}
}

func UserFromRecord(record *model.Record) *model.User {
	if record == nil {
		return nil
	}

	var user = &model.User{
		Id:        record.Id,
		Type:      record.Type,
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if record.Properties.AsMap()["username"] != nil {
		user.Username = record.Properties.AsMap()["username"].(string)
	}

	if record.Properties.AsMap()["password"] != nil {
		user.Password = record.Properties.AsMap()["password"].(string)
	}

	if record.Properties.AsMap()["scopes"] != nil {
		user.Scopes = strings.Split(record.Properties.AsMap()["scopes"].(string), ",")
	}

	if record.Properties.AsMap()["details"] != nil {
		user.Details = record.Properties.AsMap()["details"].(*structpb.Value).GetStructValue()
	}

	return user
}
