package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsSubscriptionFilter = `{
  "block": {
    "attributes": {
      "apply_on_transformed_logs": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "destination_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "distribution": {
        "computed": true,
        "description": "The method used to distribute log data to the destination, which can be either random or grouped by log stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_name": {
        "computed": true,
        "description": "The name of the subscription filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_pattern": {
        "computed": true,
        "description": "The filtering expressions that restrict what gets delivered to the destination AWS resource. For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_group_name": {
        "computed": true,
        "description": "The log group to associate with the subscription filter. All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::SubscriptionFilter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsSubscriptionFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsSubscriptionFilter), &result)
	return &result
}
