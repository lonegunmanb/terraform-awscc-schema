package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontFieldLevelEncryptionProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the field-level encryption profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "field_level_encryption_profile_config": {
        "computed": true,
        "description": "The configuration of a field-level encryption profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "caller_reference": {
              "computed": true,
              "description": "A unique value that identifies the creation request. Caller references are unique within an AWS account and cannot be changed after creation.",
              "description_kind": "plain",
              "type": "string"
            },
            "comment": {
              "computed": true,
              "description": "An optional comment describing the field-level encryption profile.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption_entities": {
              "computed": true,
              "description": "The encryption entities of the field-level encryption profile. At least one entity is required.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "field_patterns": {
                    "computed": true,
                    "description": "The request-body field names to encrypt. A pattern is either a full field name or leading characters followed by a wildcard (*). Patterns are case-sensitive and must not overlap.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "provider_id": {
                    "computed": true,
                    "description": "The provider associated with the public key. The same value must be supplied with the private key for an application to decrypt the data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "public_key_id": {
                    "computed": true,
                    "description": "The identifier of the CloudFront public key used to encrypt the fields that match the patterns.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the field-level encryption profile. Names are unique within an AWS account.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "field_level_encryption_profile_id": {
        "computed": true,
        "description": "The identifier that CloudFront assigns to the field-level encryption profile.",
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
        "description": "The time the field-level encryption profile was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::FieldLevelEncryptionProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontFieldLevelEncryptionProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFieldLevelEncryptionProfile), &result)
	return &result
}
