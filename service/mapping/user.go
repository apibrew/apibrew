package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func UserToRecord(user *model.User) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["username"] = structpb.NewStringValue(user.Username)
	properties["password"] = structpb.NewStringValue(user.Password)
	if user.Details != nil {
		properties["details"] = structpb.NewStructValue(user.Details)
	}
	properties["scopes"] = structpb.NewStringValue(strings.Join(user.Scopes, ","))

	return &model.Record{
		Id:         user.Id,
		Resource:   system.UserResource.Name,
		DataType:   user.Type,
		Properties: properties,
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
		Type:      record.DataType,
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if record.Properties["username"] != nil {
		user.Username = record.Properties["username"].GetStringValue()
	}

	if record.Properties["password"] != nil {
		user.Password = record.Properties["password"].GetStringValue()
	}

	if record.Properties["scopes"] != nil {
		user.Scopes = strings.Split(record.Properties["scopes"].GetStringValue(), ",")
	}

	if record.Properties["details"] != nil {
		user.Details = record.Properties["details"].GetStructValue()
	}

	return user
}
