package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSecurityConfiguration = `{
  "block": {
    "attributes": {
      "encryption_configuration": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
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
            "job_bookmarks_encryption": {
              "computed": true,
              "description": "The encryption configuration for job bookmarks.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "job_bookmarks_encryption_mode": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  },
                  "s3_encryption_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name for the security configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::SecurityConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueSecurityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSecurityConfiguration), &result)
	return &result
}
