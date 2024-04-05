package resource

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
        "description": "Application source config",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_url_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "approved_origins": {
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "description": {
        "description": "The application description.",
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
      "name": {
        "description": "The name of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The namespace of the application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Corresponding tag value for the key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS:AppIntegrations::Application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppintegrationsApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppintegrationsApplication), &result)
	return &result
}
