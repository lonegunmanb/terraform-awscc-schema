package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppintegrationsApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_id": {
        "computed": true,
        "description": "The id of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_source_config": {
        "computed": true,
        "description": "Application source config",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_url_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "approved_origins": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
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
      "description": {
        "computed": true,
        "description": "The application description.",
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
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The namespace of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Corresponding tag value for the key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::AppIntegrations::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppintegrationsApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppintegrationsApplication), &result)
	return &result
}
