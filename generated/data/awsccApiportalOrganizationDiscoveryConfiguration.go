package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalOrganizationDiscoveryConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the organization discovery configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the organization discovery configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_configuration_status": {
        "computed": true,
        "description": "The status of the organization discovery configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_features": {
        "computed": true,
        "description": "Optional discovery capabilities to enable.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "discovery_filters": {
        "computed": true,
        "description": "Scopes organization-wide discovery. Exactly one of ExclusionFilters or InclusionFilters is set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclusion_filters": {
              "computed": true,
              "description": "Resources matching any of these filters are excluded from discovery.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_ids": {
                    "computed": true,
                    "description": "AWS account IDs to exclude from discovery.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "organizational_unit_ids": {
                    "computed": true,
                    "description": "Organizational unit IDs to exclude from discovery.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "resource_arns": {
                    "computed": true,
                    "description": "Resource ARNs to exclude from discovery.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "inclusion_filters": {
              "computed": true,
              "description": "Only resources matching any of these filters are included in discovery.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_ids": {
                    "computed": true,
                    "description": "AWS account IDs to include in discovery.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "organizational_unit_ids": {
                    "computed": true,
                    "description": "Organizational unit IDs to include in discovery.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "discovery_sources": {
        "computed": true,
        "description": "The AWS services whose resources are scanned to discover APIs.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the organization discovery configuration.",
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
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the organization discovery configuration was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::OrganizationDiscoveryConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalOrganizationDiscoveryConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalOrganizationDiscoveryConfiguration), &result)
	return &result
}
