package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPanoramaPackage = `{
  "block": {
    "attributes": {
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "package_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "package_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "binary_prefix_location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "generated_prefix_location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_prefix_location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "repo_prefix_location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
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
    "description": "Schema for Package CloudFormation Resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPanoramaPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaPackage), &result)
	return &result
}
