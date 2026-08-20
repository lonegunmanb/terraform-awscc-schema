package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAccountaccessApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The ARN of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the application was created",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_source": {
        "computed": true,
        "description": "The identity source for the application",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identity_center": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_arn": {
                    "computed": true,
                    "description": "The ARN of the associated Identity Center application",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_arn": {
                    "computed": true,
                    "description": "The ARN of the Identity Center instance",
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
      "status": {
        "computed": true,
        "description": "The status of the application",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "tenant_id": {
        "computed": true,
        "description": "The tenant ID of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the application was last updated",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AccountAccess::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAccountaccessApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccountaccessApplication), &result)
	return &result
}
