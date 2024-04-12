package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaLayerVersion = `{
  "block": {
    "attributes": {
      "compatible_architectures": {
        "computed": true,
        "description": "A list of compatible instruction set architectures.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "compatible_runtimes": {
        "computed": true,
        "description": "A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "content": {
        "computed": true,
        "description": "The function layer archive.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description": "The Amazon S3 bucket of the layer archive.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The Amazon S3 key of the layer archive.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description": "For versioned objects, the version of the layer archive object to use.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "layer_name": {
        "computed": true,
        "description": "The name or Amazon Resource Name (ARN) of the layer.",
        "description_kind": "plain",
        "type": "string"
      },
      "layer_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_info": {
        "computed": true,
        "description": "The layer's software license.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::LayerVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaLayerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaLayerVersion), &result)
	return &result
}
