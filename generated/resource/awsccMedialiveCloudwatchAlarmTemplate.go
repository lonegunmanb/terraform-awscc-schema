package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveCloudwatchAlarmTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "A cloudwatch alarm template's ARN (Amazon Resource Name)",
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_alarm_template_id": {
        "computed": true,
        "description": "A cloudwatch alarm template's id. AWS provided templates have ids that start with ` + "`" + `aws-` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "description": "The comparison operator used to compare the specified statistic and the threshold.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "datapoints_to_alarm": {
        "computed": true,
        "description": "The number of datapoints within the evaluation period that must be breaching to trigger the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "A resource's optional description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_periods": {
        "computed": true,
        "description": "The number of periods over which data is compared to the specified threshold.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "group_id": {
        "computed": true,
        "description": "A cloudwatch alarm template group's id. AWS provided template groups have ids that start with ` + "`" + `aws-` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "group_identifier": {
        "description": "A cloudwatch alarm template group's identifier. Can be either be its id or current name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metric_name": {
        "description": "The name of the metric associated with the alarm. Must be compatible with targetResourceType.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "computed": true,
        "description": "The period, in seconds, over which the specified statistic is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "statistic": {
        "description": "The statistic to apply to the alarm's metric data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Represents the tags associated with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_resource_type": {
        "description": "The resource type this template should dynamically generate cloudwatch metric alarms for.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "threshold": {
        "computed": true,
        "description": "The threshold value to compare with the specified statistic.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "treat_missing_data": {
        "description": "Specifies how missing data points are treated when evaluating the alarm's condition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::MediaLive::CloudWatchAlarmTemplate Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveCloudwatchAlarmTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveCloudwatchAlarmTemplate), &result)
	return &result
}
