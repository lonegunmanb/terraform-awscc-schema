package resource

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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "compatible_runtimes": {
        "computed": true,
        "description": "A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "content": {
        "description": "The function layer archive.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "description": "The Amazon S3 bucket of the layer archive.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description": "The Amazon S3 key of the layer archive.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description": "For versioned objects, the version of the layer archive object to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "The description of the version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "layer_name": {
        "computed": true,
        "description": "The name or Amazon Resource Name (ARN) of the layer.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::LayerVersion",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaLayerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaLayerVersion), &result)
	return &result
}
