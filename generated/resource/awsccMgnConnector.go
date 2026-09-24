package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the connector.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "cloudwatch_output_enabled": {
              "computed": true,
              "description": "Whether SSM command output is sent to CloudWatch Logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "output_s3_bucket_name": {
              "computed": true,
              "description": "The S3 bucket name for SSM command output.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_output_enabled": {
              "computed": true,
              "description": "Whether SSM command output is stored in S3.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "ssm_instance_id": {
        "description": "The SSM instance ID associated with this connector.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::MGN::Connector. A Connector provides connectivity between a source environment and Application Migration Service (MGN) via AWS Systems Manager (SSM).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMgnConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnConnector), &result)
	return &result
}
