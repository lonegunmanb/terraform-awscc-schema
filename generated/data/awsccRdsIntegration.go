package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsIntegration = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description": "An optional set of non-secret key–value pairs that contains additional contextual information about the data.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_filter": {
        "computed": true,
        "description": "The data filter for the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_arn": {
        "computed": true,
        "description": "The ARN of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_name": {
        "computed": true,
        "description": "The name of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "An optional AWS Key Management System (AWS KMS) key ARN for the key used to to encrypt the integration. The resource accepts the key ID and the key ARN forms. The key ID form can be used if the KMS key is owned by te same account. If the KMS key belongs to a different account than the calling account, the full key ARN must be specified. Do not use the key alias or the key alias ARN as this will cause a false drift of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Aurora DB cluster to use as the source for replication.",
        "description_kind": "plain",
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
      "target_arn": {
        "computed": true,
        "description": "The ARN of the Redshift data warehouse to use as the target for replication.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RDS::Integration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsIntegration), &result)
	return &result
}
