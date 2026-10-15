package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreConsentPortal = `{
  "block": {
    "attributes": {
      "consent_portal_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "consent_portal_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "idp_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audience": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "credential_provider_arn": {
              "description": "The ARN of an AWS::BedrockAgentCore::OAuth2CredentialProvider resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scopes": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portal_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sources": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the consent portal",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::ConsentPortal",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreConsentPortalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreConsentPortal), &result)
	return &result
}
