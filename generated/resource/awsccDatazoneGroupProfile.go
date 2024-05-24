package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneGroupProfile = `{
  "block": {
    "attributes": {
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the group profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description": "The identifier of the Amazon DataZone domain in which the group profile would be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_identifier": {
        "description": "The ID of the group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The group-name of the Group Profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_profile_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone group profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the group profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Group profiles represent groups of Amazon DataZone users. Groups can be manually created, or mapped to Active Directory groups of enterprise customers. In Amazon DataZone, groups serve two purposes. First, a group can map to a team of users in the organizational chart, and thus reduce the administrative work of a Amazon DataZone project owner when there are new employees joining or leaving a team. Second, corporate administrators use Active Directory groups to manage and update user statuses and so Amazon DataZone domain administrators can use these group memberships to implement Amazon DataZone domain policies.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneGroupProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneGroupProfile), &result)
	return &result
}
