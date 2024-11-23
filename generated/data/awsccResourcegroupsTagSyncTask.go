package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResourcegroupsTagSyncTask = `{
  "block": {
    "attributes": {
      "group": {
        "computed": true,
        "description": "The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the TagSyncTask",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_key": {
        "computed": true,
        "description": "The tag key. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_value": {
        "computed": true,
        "description": "The tag value. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_arn": {
        "computed": true,
        "description": "The ARN of the TagSyncTask resource",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResourceGroups::TagSyncTask",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResourcegroupsTagSyncTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResourcegroupsTagSyncTask), &result)
	return &result
}
