package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchSpace = `{
  "block": {
    "attributes": {
      "agent_core_evaluation_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role used by AgentCore online evaluation. Absent when the space was created without one.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The timestamp when the space was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "description": "The ARN of the IAM role used for data access.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_id": {
        "description": "The ID of the domain to create the space in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "How the space's data at rest is encrypted. Always populated.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_strategy": {
              "computed": true,
              "description": "Which kind of key to use. Required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "Customer managed KMS key ARN. Required when encryptionStrategy is CUSTOMER_MANAGED, and must be omitted when it is AWS_OWNED. Must be a symmetric ENCRYPT_DECRYPT key in the caller's account and region.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "A name that identifies the space.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "AWS account ID that owns this space.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region where this space was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "space_id": {
        "computed": true,
        "description": "The unique ID of the space.",
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
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the space was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::Space. A space is where you work in CloudWatch Omni. It gives you access to telemetry from the account and Region in which the space is hosted, and controls who can see it.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchSpace), &result)
	return &result
}
