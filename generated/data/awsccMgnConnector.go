package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMgnConnector = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_id": {
        "computed": true,
        "description": "The unique identifier of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssm_command_config": {
        "computed": true,
        "description": "SSM command configuration for the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_log_group_name": {
              "computed": true,
              "description": "The CloudWatch Logs group name for SSM command output.",
              "description_kind": "plain",
              "type": "string"
            },
            "cloudwatch_output_enabled": {
              "computed": true,
              "description": "Whether SSM command output is sent to CloudWatch Logs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "output_s3_bucket_name": {
              "computed": true,
              "description": "The S3 bucket name for SSM command output.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_output_enabled": {
              "computed": true,
              "description": "Whether SSM command output is stored in S3.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "ssm_instance_id": {
        "computed": true,
        "description": "The SSM instance ID associated with this connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::MGN::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMgnConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnConnector), &result)
	return &result
}
