package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchDomainForOrganization = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the organization domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the organization domain was created.",
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
      "domain_access_role_arn": {
        "description": "The ARN of the customer-provided IAM role in the management account used for domain access.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_endpoint_url": {
        "computed": true,
        "description": "The HTTPS endpoint URL for accessing the organization domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique ID of the organization domain.",
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
        "description": "Identity provider configuration for the organization domain.",
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
        "description": "The identity providers configured for the organization domain.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "A name that identifies the organization domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_id": {
        "computed": true,
        "description": "The ID of the AWS Organization that owns the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "The AWS account ID that owns the organization domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region where this organization domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the organization domain.",
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
        "description": "The timestamp when the organization domain was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::DomainForOrganization. An organization domain provides a single sign-in URL and centralized identity provider configuration for your AWS Organization, removing the need for each account to configure its own domain. Only the management account can create it.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchDomainForOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDomainForOrganization), &result)
	return &result
}
