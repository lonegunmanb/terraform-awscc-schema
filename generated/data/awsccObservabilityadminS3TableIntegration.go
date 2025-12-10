package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccObservabilityadminS3TableIntegration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the S3 Table Integration",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption": {
        "computed": true,
        "description": "Encryption configuration for the S3 Table Integration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key used to encrypt the S3 Table Integration",
              "description_kind": "plain",
              "type": "string"
            },
            "sse_algorithm": {
              "computed": true,
              "description": "The server-side encryption algorithm used to encrypt the S3 Table(s) data",
              "description_kind": "plain",
              "type": "string"
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
      "log_sources": {
        "computed": true,
        "description": "The CloudWatch Logs data sources to associate with the S3 Table Integration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "The ID of the CloudWatch Logs data source association",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the CloudWatch Logs data source",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the CloudWatch Logs data source",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the role used to access the S3 Table Integration",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::ObservabilityAdmin::S3TableIntegration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccObservabilityadminS3TableIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccObservabilityadminS3TableIntegration), &result)
	return &result
}
