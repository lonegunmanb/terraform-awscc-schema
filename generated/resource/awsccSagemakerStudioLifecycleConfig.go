package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerStudioLifecycleConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "studio_lifecycle_config_app_type": {
        "description": "The App type that the Lifecycle Configuration is attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "studio_lifecycle_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "studio_lifecycle_config_content": {
        "description": "The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "studio_lifecycle_config_name": {
        "description": "The name of the Amazon SageMaker Studio Lifecycle Configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::SageMaker::StudioLifecycleConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerStudioLifecycleConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerStudioLifecycleConfig), &result)
	return &result
}
