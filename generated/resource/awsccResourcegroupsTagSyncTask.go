package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResourcegroupsTagSyncTask = `{
  "block": {
    "attributes": {
      "group": {
        "description": "The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the ApplicationGroup for which the TagSyncTask is created",
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The Name of the application group for which the TagSyncTask is created",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the TagSyncTask",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_key": {
        "description": "The tag key. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_value": {
        "description": "The tag value. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_arn": {
        "computed": true,
        "description": "The ARN of the TagSyncTask resource",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Schema for ResourceGroups::TagSyncTask",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResourcegroupsTagSyncTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResourcegroupsTagSyncTask), &result)
	return &result
}
