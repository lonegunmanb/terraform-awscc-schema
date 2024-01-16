package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmResourceDataSync = `{
  "block": {
    "attributes": {
      "bucket_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_destination": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sync_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sync_format": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_source": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_organizations_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "organization_source_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "organizational_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "include_future_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "source_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "source_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sync_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::ResourceDataSync",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmResourceDataSyncSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmResourceDataSync), &result)
	return &result
}
