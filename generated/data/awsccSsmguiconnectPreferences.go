package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmguiconnectPreferences = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_recording_preferences": {
        "computed": true,
        "description": "The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region. This includes details such as which S3 bucket recordings are stored in.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of a AWS KMS key that is used to encrypt data while it is being processed by the service. This key must exist in the same AWS Region as the node you start an RDP connection to.",
              "description_kind": "plain",
              "type": "string"
            },
            "recording_destinations": {
              "computed": true,
              "description": "Determines where recordings of RDP connections are stored.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_buckets": {
                    "computed": true,
                    "description": "The S3 bucket where RDP connection recordings are stored.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description": "The name of the S3 bucket where RDP connection recordings are stored.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description": "The AWS account number that owns the S3 bucket.",
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
      }
    },
    "description": "Data Source schema for AWS::SSMGuiConnect::Preferences",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmguiconnectPreferencesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmguiconnectPreferences), &result)
	return &result
}
