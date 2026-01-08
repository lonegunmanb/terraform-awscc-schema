package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsSubscriptionFilter = `{
  "block": {
    "attributes": {
      "apply_on_transformed_logs": {
        "computed": true,
        "description": "This parameter is valid only for log groups that have an active log transformer. For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).\n If this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the subscription filter is applied on the transformed version of the log events instead of the original ingested log events.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "destination_arn": {
        "description": "The Amazon Resource Name (ARN) of the destination.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "distribution": {
        "computed": true,
        "description": "The method used to distribute log data to the destination, which can be either random or grouped by log stream.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "emit_system_fields": {
        "computed": true,
        "description": "The list of system fields that are included in the log events sent to the subscription destination. Returns the ` + "`" + `` + "`" + `emitSystemFields` + "`" + `` + "`" + ` value if it was specified when the subscription filter was created.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "field_selection_criteria": {
        "computed": true,
        "description": "The filter expression that specifies which log events are processed by this subscription filter based on system fields. Returns the ` + "`" + `` + "`" + `fieldSelectionCriteria` + "`" + `` + "`" + ` value if it was specified when the subscription filter was created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_name": {
        "computed": true,
        "description": "The name of the subscription filter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_pattern": {
        "description": "The filtering expressions that restrict what gets delivered to the destination AWS resource. For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).",
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
      "log_group_name": {
        "description": "The log group to associate with the subscription filter. All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::Logs::SubscriptionFilter` + "`" + `` + "`" + ` resource specifies a subscription filter and associates it with the specified log group. Subscription filters allow you to subscribe to a real-time stream of log events and have them delivered to a specific destination. Currently, the supported destinations are:\n  +  An Amazon Kinesis data stream belonging to the same account as the subscription filter, for same-account delivery.\n  +  A logical destination that belongs to a different account, for cross-account delivery.\n  +  An Amazon Kinesis Firehose delivery stream that belongs to the same account as the subscription filter, for same-account delivery.\n  +  An LAMlong function that belongs to the same account as the subscription filter, for same-account delivery.\n  \n There can be as many as two subscription filters associated with a log group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsSubscriptionFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsSubscriptionFilter), &result)
	return &result
}
