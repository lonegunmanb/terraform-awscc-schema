package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchSpace = `{
  "block": {
    "attributes": {
      "agent_core_evaluation_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role used for AgentCore evaluation.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the space was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role the space uses to access observability data.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the domain the space belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the space.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_strategy": {
              "computed": true,
              "description": "Whether the space is encrypted with an AWS owned key or a customer managed key.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the customer managed KMS key. Required when encryptionStrategy is CUSTOMER_MANAGED.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "The AWS account that owns the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The AWS Region the space resides in.",
        "description_kind": "plain",
        "type": "string"
      },
      "space_id": {
        "computed": true,
        "description": "The unique identifier of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
      },
      "updated_at": {
        "computed": true,
        "description": "The time the space was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::Space",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchSpace), &result)
	return &result
}
