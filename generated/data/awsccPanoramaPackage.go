package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "package_name": {
        "computed": true,
        "description": "A name for the package.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_location": {
        "computed": true,
        "description": "A storage location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "binary_prefix_location": {
              "computed": true,
              "description": "The location's binary prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "bucket": {
              "computed": true,
              "description": "The location's bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "generated_prefix_location": {
              "computed": true,
              "description": "The location's generated prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_prefix_location": {
              "computed": true,
              "description": "The location's manifest prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "repo_prefix_location": {
              "computed": true,
              "description": "The location's repo prefix.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags for the package.",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Panorama::Package",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPanoramaPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaPackage), &result)
	return &result
}
