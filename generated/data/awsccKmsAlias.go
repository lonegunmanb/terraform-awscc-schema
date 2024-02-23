package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
        "computed": true,
        "description": "Specifies the alias name. This value must begin with ` + "`" + `` + "`" + `alias/` + "`" + `` + "`" + ` followed by a name, such as ` + "`" + `` + "`" + `alias/ExampleAlias` + "`" + `` + "`" + `. \n  If you change the value of the ` + "`" + `` + "`" + `AliasName` + "`" + `` + "`" + ` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).\n  The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ` + "`" + `` + "`" + `alias/aws/` + "`" + `` + "`" + `. The ` + "`" + `` + "`" + `alias/aws/` + "`" + `` + "`" + ` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_key_id": {
        "computed": true,
        "description": "Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.\n A valid key ID is required. If you supply a null or empty string value, this operation returns an error.\n For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.\n Specify the key ID or the key ARN of the KMS key.\n For example:\n  +  Key ID: ` + "`" + `` + "`" + `1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  +  Key ARN: ` + "`" + `` + "`" + `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  \n To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::KMS::Alias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKmsAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsAlias), &result)
	return &result
}
