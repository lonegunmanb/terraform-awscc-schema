package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackageAsset = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the Asset was initially submitted for Ingest.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_endpoints": {
        "computed": true,
        "description": "The list of egress endpoints available for the Asset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "packaging_configuration_id": {
              "computed": true,
              "description": "The ID of the PackagingConfiguration being applied to the Asset.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "The URL of the parent manifest for the repackaged Asset.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "packaging_group_id": {
        "computed": true,
        "description": "The ID of the PackagingGroup for the Asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The resource ID to include in SPEKE key requests.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "ARN of the source object in S3.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_role_arn": {
        "computed": true,
        "description": "The IAM role_arn used to access the source S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
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
      }
    },
    "description": "Data Source schema for AWS::MediaPackage::Asset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackageAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackageAsset), &result)
	return &result
}
