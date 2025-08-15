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
      "application_config": {
        "computed": true,
        "description": "The application configuration. Cannot be used when IsService is true.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "contact_handling": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scope": {
                    "computed": true,
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
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
      "iframe_config": {
        "computed": true,
        "description": "The iframe configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sandbox": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "initialization_timeout": {
        "computed": true,
        "description": "The initialization timeout in milliseconds. Required when IsService is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "is_service": {
        "computed": true,
        "description": "Indicates if the application is a service",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The name of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description": "The namespace of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description": "The configuration of events or requests that the application has access to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Corresponding tag value for the key.",
              "description_kind": "plain",
              "optional": true,
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
