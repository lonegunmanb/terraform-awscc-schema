package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53QueryLoggingConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the query logging configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_logs_log_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group in us-east-1 that Amazon Route 53 publishes query logs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description": "The ID of the public hosted zone that Amazon Route 53 logs queries for.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_logging_config_id": {
        "computed": true,
        "description": "The ID that Amazon Route 53 assigns to the configuration for DNS query logging.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53::QueryLoggingConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53QueryLoggingConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53QueryLoggingConfig), &result)
	return &result
}
