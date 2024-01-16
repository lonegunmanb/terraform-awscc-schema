package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRefactorspacesRoute = `{
  "block": {
    "attributes": {
      "application_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_route": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "activation_state": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "environment_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path_resource_to_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "uri_path_route": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "activation_state": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "append_source_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "include_child_paths": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "methods": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "source_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::RefactorSpaces::Route",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRefactorspacesRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRefactorspacesRoute), &result)
	return &result
}
