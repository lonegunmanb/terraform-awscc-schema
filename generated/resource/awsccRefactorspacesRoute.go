package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRefactorspacesRoute = `{
  "block": {
    "attributes": {
      "application_identifier": {
        "description_kind": "plain",
        "required": true,
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "environment_identifier": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "uri_path_route": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "activation_state": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "append_source_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_child_paths": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "methods": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "source_path": {
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
    "description": "Definition of AWS::RefactorSpaces::Route Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRefactorspacesRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRefactorspacesRoute), &result)
	return &result
}
