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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_data": {
              "computed": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_role_arn": {
        "computed": true,
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
        "description": "List of tags",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Schema for ApplicationInstance CloudFormation Resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPanoramaApplicationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaApplicationInstance), &result)
	return &result
}
