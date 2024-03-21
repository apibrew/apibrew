package helper

import (
	"github.com/apibrew/apibrew/pkg/model"
)

func IsPropertyOmitted(property *model.ResourceProperty) bool {
	if property.Type == model.ResourceProperty_LIST && property.Item != nil && property.Item.Type == model.ResourceProperty_REFERENCE && property.Item.BackReference != nil {
		// skip back references as they will be populated on service layer
		return true
	}

	if property.Virtual {
		return true
	}

	if property.BackReference != nil {
		return true
	}

	return false
}

// SELECT "t"."email" as "t_email", "t"."picture" as "t_picture", "t"."surname" as "t_surname", "t"."jobTitle" as "t_jobTitle", "t"."skills" as "t_skills", "t"."cv" as "t_cv", "t"."user" as "t_user", "t"."id" as "t_id", "t"."username" as "t_username", "t"."about" as "t_about", "t"."auditData" as "t_auditData", "t"."savedJobs" as "t_savedJobs", "t"."educations" as "t_educations", "t"."version" as "t_version", "t"."appliedJobs" as "t_appliedJobs", "t"."verifiedBy" as "t_verifiedBy", "t"."slug" as "t_slug", "t"."sector" as "t_sector", "t"."name" as "t_name", "t"."status" as "t_status"
// FROM "talent" as t
//ORDER BY t.id ASC LIMIT 10 OFFSET 0 ; Bind Params:
