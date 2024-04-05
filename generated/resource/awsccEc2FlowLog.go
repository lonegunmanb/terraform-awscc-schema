package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2FlowLog = `{
  "block": {
    "attributes": {
      "deliver_cross_account_role": {
        "computed": true,
        "description": "The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deliver_logs_permission_arn": {
        "computed": true,
        "description": "The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_format": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hive_compatible_partitions": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "per_hour_partition": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "flow_log_id": {
        "computed": true,
        "description": "The Flow Log ID",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_destination": {
        "computed": true,
        "description": "Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_destination_type": {
        "computed": true,
        "description": "Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_format": {
        "computed": true,
        "description": "The fields to include in the flow log record, in the order in which they should appear.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_group_name": {
        "computed": true,
        "description": "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_aggregation_interval": {
        "computed": true,
        "description": "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_id": {
        "description": "The ID of the subnet, network interface, or VPC for which you want to create a flow log.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "description": "The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the flow logs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "traffic_type": {
        "computed": true,
        "description": "The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2FlowLogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2FlowLog), &result)
	return &result
}
