package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAccountaccessEntitlement = `{
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
        "description": "The timestamp when the entitlement was created",
        "description_kind": "plain",
        "type": "string"
      },
      "entitlement": {
        "computed": true,
        "description": "The entitlement details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "principal_role": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account": {
                    "computed": true,
                    "description": "The AWS account ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "principal": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "identity_center": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "group_id": {
                                "computed": true,
                                "description": "The ID of the group",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "user_id": {
                                "computed": true,
                                "description": "The ID of the user",
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
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the IAM role",
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
      "entitlement_id": {
        "computed": true,
        "description": "The ID of the entitlement",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AccountAccess::Entitlement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAccountaccessEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccountaccessEntitlement), &result)
	return &result
}
