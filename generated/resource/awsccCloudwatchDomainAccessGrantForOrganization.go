package resource

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
        "description": "The unique identifier of the organization domain the grant applies to.",
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the grant.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission": {
        "description": "The permission level the grant confers. Organization grants are administrative only.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The attribute value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "principal_id": {
              "computed": true,
              "description": "The identifier of the principal, such as an IAM role ARN or an Identity Center user ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "principal_type": {
              "description": "The type of principal the grant is issued to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
        "description": "The time the grant was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::DomainAccessGrantForOrganization. An organization access grant authorizes a principal to administer an organization-scoped CloudWatch Omni domain.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchDomainAccessGrantForOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDomainAccessGrantForOrganization), &result)
	return &result
}
