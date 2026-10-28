package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchDomain = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_endpoint_urls": {
        "computed": true,
        "description": "Additional endpoint URLs derived from the domain name.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "domain_endpoint_url": {
        "computed": true,
        "description": "The HTTPS endpoint URL for accessing the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique ID of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_application_arn": {
        "computed": true,
        "description": "The ARN of the Identity Center application. Absent for IAM-only domains.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_configuration": {
        "computed": true,
        "description": "Identity provider configuration for the domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identity_center_configuration": {
              "computed": true,
              "description": "Identity Center configuration. Required when identityProviders includes IDC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "identity_center_instance_arn": {
                    "computed": true,
                    "description": "Identity Center instance ARN",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "identity_providers": {
        "description": "The identity providers configured for the domain.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "A name that identifies the domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region where this domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the domain.",
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
        "description": "The timestamp when the domain was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The account-level boundary for CloudWatch Omni. A domain configures the\nidentity providers users sign in through and exposes the endpoint URLs used\nto reach CloudWatch Omni. Spaces are created within a domain.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDomain), &result)
	return &result
}
