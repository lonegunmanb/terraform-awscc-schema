package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerHub = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The date and time that the hub was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the hub.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_description": {
        "computed": true,
        "description": "A description of the hub.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_display_name": {
        "computed": true,
        "description": "The display name of the hub.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_name": {
        "computed": true,
        "description": "The name of the hub.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_search_keywords": {
        "computed": true,
        "description": "The searchable keywords for the hub.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "hub_status": {
        "computed": true,
        "description": "The status of the hub.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date and time that the hub was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_storage_config": {
        "computed": true,
        "description": "The Amazon S3 storage configuration for the hub.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_output_path": {
              "computed": true,
              "description": "The Amazon S3 bucket prefix for hosting hub content.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the hub.",
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
    "description": "Data Source schema for AWS::SageMaker::Hub",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerHubSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerHub), &result)
	return &result
}
