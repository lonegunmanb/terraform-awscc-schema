package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSecurityConfiguration = `{
  "block": {
    "attributes": {
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the security configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_encryption": {
              "computed": true,
              "description": "The encryption configuration for Amazon CloudWatch.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_encryption_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "job_bookmarks_encryption": {
              "computed": true,
              "description": "The encryption configuration for job bookmarks.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "job_bookmarks_encryption_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_encryptions": {
              "computed": true,
              "description": "The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_encryption_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the security configuration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::SecurityConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueSecurityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSecurityConfiguration), &result)
	return &result
}
