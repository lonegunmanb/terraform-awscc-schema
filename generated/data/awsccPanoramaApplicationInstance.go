package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPanoramaApplicationInstance = `{
  "block": {
    "attributes": {
      "application_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_instance_id_to_replace": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "default_runtime_context_device": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_runtime_context_device_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "health_status": {
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
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "manifest_overrides_payload": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "manifest_payload": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of tags",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Panorama::ApplicationInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPanoramaApplicationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaApplicationInstance), &result)
	return &result
}
