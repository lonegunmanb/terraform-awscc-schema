package resource

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
        "optional": true,
        "type": "string"
      },
      "bucket_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket_region": {
        "computed": true,
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
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "bucket_prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sync_format": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sync_format": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sync_name": {
        "description_kind": "plain",
        "required": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "include_future_regions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "source_regions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "source_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sync_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SSM::ResourceDataSync",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmResourceDataSyncSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmResourceDataSync), &result)
	return &result
}
