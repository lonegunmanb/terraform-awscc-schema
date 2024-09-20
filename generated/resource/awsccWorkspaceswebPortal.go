package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebPortal = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "authentication_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "browser_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "browser_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_access_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_concurrent_sessions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "network_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "portal_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "portal_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renderer_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_provider_saml_metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
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
      "trust_store_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_access_logging_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::WorkSpacesWeb::Portal Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspaceswebPortalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebPortal), &result)
	return &result
}
