package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseWorkspace = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time the workspace was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "description": "The encryption configuration for the workspace.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "description": "The type of encryption.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ARN of the AWS KMS key used for KMS_BASED_ENCRYPTION. Required when EncryptionConfiguration.EncryptionType is KMS_BASED_ENCRYPTION.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current state of the workspace.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time the workspace was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_arn": {
        "computed": true,
        "description": "The ARN of the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_description": {
        "computed": true,
        "description": "A description of the workspace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_name": {
        "description": "The name of the workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Represents an AWS IoT SiteWise workspace that provides logical isolation for tasks and pipelines.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseWorkspace), &result)
	return &result
}
