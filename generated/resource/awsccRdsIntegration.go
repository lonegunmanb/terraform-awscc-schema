package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsIntegration = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description": "An optional set of non-secret key?value pairs that contains additional contextual information about the data. For more information, see [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the *Key Management Service Developer Guide*.\n You can only include this parameter if you specify the ` + "`" + `` + "`" + `KMSKeyId` + "`" + `` + "`" + ` parameter.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the integration.",
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
      "integration_arn": {
        "computed": true,
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
        "description": "The AWS Key Management System (AWS KMS) key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, RDS uses a default AWS owned key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_arn": {
        "description": "The Amazon Resource Name (ARN) of the database to use as the source for replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
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
        "description": "The ARN of the Redshift data warehouse to use as the target for replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A zero-ETL integration with Amazon Redshift.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsIntegration), &result)
	return &result
}
