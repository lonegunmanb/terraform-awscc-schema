package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveCloudwatchAlarmTemplateGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "A cloudwatch alarm template group's ARN (Amazon Resource Name)",
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_alarm_template_group_id": {
        "computed": true,
        "description": "A cloudwatch alarm template group's id. AWS provided template groups have ids that start with ` + "`" + `aws-` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A resource's optional description.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Represents the tags associated with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplateGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMedialiveCloudwatchAlarmTemplateGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveCloudwatchAlarmTemplateGroup), &result)
	return &result
}
