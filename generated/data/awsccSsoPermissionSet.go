package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoPermissionSet = `{
  "block": {
    "attributes": {
      "customer_managed_policy_references": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "The permission set description.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inline_policy": {
        "computed": true,
        "description": "The inline policy to put in permission set.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The sso instance arn that the permission set is owned.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The name you want to assign to this permission set.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_set_arn": {
        "computed": true,
        "description": "The permission set that the policy will be attached to",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_boundary": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_managed_policy_reference": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "managed_policy_arn": {
              "computed": true,
              "description": "The managed policy to attach.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "relay_state_type": {
        "computed": true,
        "description": "The relay state URL that redirect links to any service in the AWS Management Console.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_duration": {
        "computed": true,
        "description": "The length of time that a user can be signed in to an AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SSO::PermissionSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsoPermissionSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoPermissionSet), &result)
	return &result
}
