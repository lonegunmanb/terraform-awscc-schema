package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerNotebookInstanceLifecycleConfig = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notebook_instance_lifecycle_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the lifecycle configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "notebook_instance_lifecycle_config_name": {
        "computed": true,
        "description": "The name of the lifecycle configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "on_create": {
        "computed": true,
        "description": "A shell script that runs only once, when you create a notebook instance. The shell script must be a base64-encoded string.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "on_start": {
        "computed": true,
        "description": "A shell script that runs every time you start a notebook instance, including when you create the notebook instance. The shell script must be a base64-encoded string.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::NotebookInstanceLifecycleConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerNotebookInstanceLifecycleConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerNotebookInstanceLifecycleConfig), &result)
	return &result
}
