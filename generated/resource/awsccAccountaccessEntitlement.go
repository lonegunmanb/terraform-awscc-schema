package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAccountaccessEntitlement = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description": "The ARN of the application",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the entitlement was created",
        "description_kind": "plain",
        "type": "string"
      },
      "entitlement": {
        "description": "The entitlement details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "principal_role": {
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
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "identity_center": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "group_id": {
                                "computed": true,
                                "description": "The ID of the group",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_id": {
                                "computed": true,
                                "description": "The ID of the user",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "role_arn": {
                    "description": "The ARN of the IAM role",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "entitlement_id": {
        "computed": true,
        "description": "The ID of the entitlement",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AccountAccess::Entitlement specifying an entitlement for account access",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAccountaccessEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccountaccessEntitlement), &result)
	return &result
}
