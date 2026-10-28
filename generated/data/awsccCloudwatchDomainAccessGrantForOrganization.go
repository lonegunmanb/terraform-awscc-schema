package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchDomainAccessGrantForOrganization = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the organization access grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the grant was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The principal that created the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique identifier of the organization domain the grant applies to.",
        "description_kind": "plain",
        "type": "string"
      },
      "grant_id": {
        "computed": true,
        "description": "The unique identifier of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "grant_type": {
        "computed": true,
        "description": "Whether the grant is managed by the service or by the customer.",
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
        "description": "The name of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission": {
        "computed": true,
        "description": "The permission level the grant confers. Organization grants are administrative only.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The principal the grant is issued to. Supply PrincipalId, or PrincipalAttributes for an attribute-based grant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "principal_attributes": {
              "computed": true,
              "description": "Attribute key-value pairs that select the principals the grant applies to.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The attribute name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The attribute value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "principal_id": {
              "computed": true,
              "description": "The identifier of the principal, such as an IAM role ARN or an Identity Center user ID.",
              "description_kind": "plain",
              "type": "string"
            },
            "principal_type": {
              "computed": true,
              "description": "The type of principal the grant is issued to.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
        "description": "The time the grant was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::DomainAccessGrantForOrganization",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchDomainAccessGrantForOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDomainAccessGrantForOrganization), &result)
	return &result
}
