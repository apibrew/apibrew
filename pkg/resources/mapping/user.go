package mapping

import (
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func UserToRecord(user *model.User) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["username"] = structpb.NewStringValue(user.Username)
	properties["password"] = structpb.NewStringValue(user.Password)
	if user.Details != nil {
		properties["details"] = structpb.NewStructValue(user.Details)
	}

	properties["securityContext"] = SecurityContextToValue(user.SecurityContext)

	MapSpecialColumnsToRecord(user, &properties)

	return &model.Record{
		Id:         user.Id,
		Properties: properties,
	}
}

func UserFromRecord(record *model.Record) *model.User {
	if record == nil {
		return nil
	}

	var user = &model.User{
		Id: record.Id,
	}

	if record.Properties["username"] != nil {
		user.Username = record.Properties["username"].GetStringValue()
	}

	if record.Properties["password"] != nil {
		user.Password = record.Properties["password"].GetStringValue()
	}

	user.SecurityContext = SecurityContextFromValue(record.Properties["securityContext"])

	if record.Properties["details"] != nil {
		user.Details = record.Properties["details"].GetStructValue()
	}

	MapSpecialColumnsFromRecord(user, &record.Properties)

	return user
}
