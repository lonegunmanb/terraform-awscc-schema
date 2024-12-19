package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftIntegration = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description": "An optional set of non-secret key–value pairs that contains additional contextual information about the data.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The time (UTC) when the integration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_name": {
        "computed": true,
        "description": "The name of the integration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_arn": {
        "description": "The Amazon Resource Name (ARN) of the database to use as the source for replication",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_arn": {
        "description": "The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Integration from a source AWS service to a Redshift cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftIntegration), &result)
	return &result
}
