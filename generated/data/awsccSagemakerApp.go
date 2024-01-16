package data

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
        "computed": true,
        "description": "The name of the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_type": {
        "computed": true,
        "description": "The type of app.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The domain ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "sage_maker_image_arn": {
              "computed": true,
              "description": "The ARN of the SageMaker image that the image version belongs to.",
              "description_kind": "plain",
              "type": "string"
            },
            "sage_maker_image_version_arn": {
              "computed": true,
              "description": "The ARN of the image version created on the instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the app.",
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
          "nesting_mode": "list"
        }
      },
      "user_profile_name": {
        "computed": true,
        "description": "The user profile name.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::App",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerApp), &result)
	return &result
}
