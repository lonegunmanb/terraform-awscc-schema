package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerImageVersion = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The alias of the image version.",
        "description_kind": "plain",
        "type": "string"
      },
      "aliases": {
        "computed": true,
        "description": "List of aliases for the image version.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "base_image": {
        "computed": true,
        "description": "The registry path of the container image on which this image version is based.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_image": {
        "computed": true,
        "description": "The registry path of the container image that contains this image version.",
        "description_kind": "plain",
        "type": "string"
      },
      "horovod": {
        "computed": true,
        "description": "Indicates Horovod compatibility.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the parent image.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_name": {
        "computed": true,
        "description": "The name of the image this version belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_version_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image version.",
        "description_kind": "plain",
        "type": "string"
      },
      "job_type": {
        "computed": true,
        "description": "Indicates SageMaker job type compatibility.",
        "description_kind": "plain",
        "type": "string"
      },
      "ml_framework": {
        "computed": true,
        "description": "The machine learning framework vended in the image version.",
        "description_kind": "plain",
        "type": "string"
      },
      "processor": {
        "computed": true,
        "description": "Indicates CPU or GPU compatibility.",
        "description_kind": "plain",
        "type": "string"
      },
      "programming_lang": {
        "computed": true,
        "description": "The supported programming language and its version.",
        "description_kind": "plain",
        "type": "string"
      },
      "release_notes": {
        "computed": true,
        "description": "The maintainer description of the image version.",
        "description_kind": "plain",
        "type": "string"
      },
      "vendor_guidance": {
        "computed": true,
        "description": "The availability of the image version specified by the maintainer.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number of the image version.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::ImageVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerImageVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerImageVersion), &result)
	return &result
}
