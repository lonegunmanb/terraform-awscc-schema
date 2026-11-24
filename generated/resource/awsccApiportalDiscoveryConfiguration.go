package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalDiscoveryConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the discovery configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the discovery configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_configuration_status": {
        "computed": true,
        "description": "The status of the discovery configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_features": {
        "computed": true,
        "description": "Optional discovery capabilities to enable.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "discovery_sources": {
        "description": "The AWS services whose resources are scanned to discover APIs.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "excluded_resources": {
        "computed": true,
        "description": "ARNs of resources to exclude from discovery.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the discovery configuration.",
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
        "description": "The timestamp when the discovery configuration was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::DiscoveryConfiguration. Represents an API discovery configuration that controls how APIs are discovered in the account",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalDiscoveryConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalDiscoveryConfiguration), &result)
	return &result
}
