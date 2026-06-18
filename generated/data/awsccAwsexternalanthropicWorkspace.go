package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAwsexternalanthropicWorkspace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the workspace was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_residency": {
        "computed": true,
        "description": "Data residency configuration for the workspace. WorkspaceGeo is immutable after creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_inference_geos": {
              "computed": true,
              "description": "Permitted inference geo values. Omit to allow all geos (the service default of 'unrestricted'); otherwise list specific geos.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "default_inference_geo": {
              "computed": true,
              "description": "Default inference geo applied when requests omit the parameter. Defaults to 'global' if omitted. Must be a member of AllowedInferenceGeos unless AllowedInferenceGeos is omitted.",
              "description_kind": "plain",
              "type": "string"
            },
            "workspace_geo": {
              "computed": true,
              "description": "Geographic region for workspace data storage. Immutable after creation. Defaults to 'us' if omitted.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "workspace_id": {
        "computed": true,
        "description": "The unique identifier of the workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AWSExternalAnthropic::Workspace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAwsexternalanthropicWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAwsexternalanthropicWorkspace), &result)
	return &result
}
