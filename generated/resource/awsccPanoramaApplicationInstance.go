package resource

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
        "optional": true,
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
        "description": "The device's ID.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "health_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "manifest_payload": {
        "description": "The application's manifest document.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
              "description": "The application manifest.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "computed": true,
        "description": "A name for the application instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_role_arn": {
        "computed": true,
        "description": "The ARN of a runtime role for the application instance.",
        "description_kind": "plain",
        "optional": true,
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Creates an application instance and deploys it to a device.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPanoramaApplicationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaApplicationInstance), &result)
	return &result
}
