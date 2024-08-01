package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerApp = `{
  "block": {
    "attributes": {
      "app_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_name": {
        "description": "The name of the app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_type": {
        "description": "The type of app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_id": {
        "description": "The domain ID.",
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
      "resource_spec": {
        "computed": true,
        "description": "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instance_type": {
              "computed": true,
              "description": "The instance type that the image version runs on.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lifecycle_config_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sage_maker_image_arn": {
              "computed": true,
              "description": "The ARN of the SageMaker image that the image version belongs to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sage_maker_image_version_arn": {
              "computed": true,
              "description": "The ARN of the image version created on the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the app.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_profile_name": {
        "description": "The user profile name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::App",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerApp), &result)
	return &result
}
