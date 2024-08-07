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
        "description": "The ID of an application instance to replace with the new instance.",
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
        "description": "The device's ID.",
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
        "description": "A description for the application instance.",
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
        "description": "Setting overrides for the application manifest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
              "description": "The overrides document.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "manifest_payload": {
        "computed": true,
        "description": "The application's manifest document.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
              "description": "The application manifest.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the application instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_role_arn": {
        "computed": true,
        "description": "The ARN of a runtime role for the application instance.",
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
        "description": "Tags for the application instance.",
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
