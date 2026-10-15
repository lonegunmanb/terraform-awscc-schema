package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesTenant = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the tenant.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reputation_entity": {
        "computed": true,
        "description": "The reputation configuration for the tenant. Sending status remains managed by SES.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy": {
              "computed": true,
              "description": "The reputation policy for the tenant.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "reputation_entity_policy": {
                    "computed": true,
                    "description": "The ARN of an SES-managed reputation policy.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_associations": {
        "computed": true,
        "description": "The list of resources to associate with the tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_arn": {
              "computed": true,
              "description": "The ARN of the resource to associate with the tenant",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "suppression_attributes": {
        "computed": true,
        "description": "The suppression attributes for the tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "suppressed_reasons": {
              "computed": true,
              "description": "The reasons that cause messages to be added to the suppression list.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "suppression_scope": {
              "computed": true,
              "description": "The scope of the suppression list for the tenant.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the key-value tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the key-value tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tenant_name": {
        "computed": true,
        "description": "The name of the tenant.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SES::Tenant",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesTenantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesTenant), &result)
	return &result
}
