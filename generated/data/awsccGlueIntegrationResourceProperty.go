package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueIntegrationResourceProperty = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The connection ARN of the source, or the database ARN of the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_property_arn": {
        "computed": true,
        "description": "The integration resource property ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_processing_properties": {
        "computed": true,
        "description": "The resource properties associated with the integration source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "computed": true,
              "description": "The IAM role to access the Glue connection.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_processing_properties": {
        "computed": true,
        "description": "The resource properties associated with the integration target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_name": {
              "computed": true,
              "description": "The Glue network connection to configure the Glue job running in the customer VPC.",
              "description_kind": "plain",
              "type": "string"
            },
            "event_bus_arn": {
              "computed": true,
              "description": "The ARN of an Eventbridge event bus to receive the integration status notification.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_arn": {
              "computed": true,
              "description": "The ARN of the KMS key used for encryption.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The IAM role to access the Glue database.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::IntegrationResourceProperty",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueIntegrationResourcePropertySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueIntegrationResourceProperty), &result)
	return &result
}
