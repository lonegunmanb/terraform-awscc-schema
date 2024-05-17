package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
        "description": "Specifies the alias name. This value must begin with ` + "`" + `` + "`" + `alias/` + "`" + `` + "`" + ` followed by a name, such as ` + "`" + `` + "`" + `alias/ExampleAlias` + "`" + `` + "`" + `. \n  If you change the value of the ` + "`" + `` + "`" + `AliasName` + "`" + `` + "`" + ` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).\n  The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ` + "`" + `` + "`" + `alias/aws/` + "`" + `` + "`" + `. The ` + "`" + `` + "`" + `alias/aws/` + "`" + `` + "`" + ` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_key_id": {
        "description": "Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.\n A valid key ID is required. If you supply a null or empty string value, this operation returns an error.\n For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.\n Specify the key ID or the key ARN of the KMS key.\n For example:\n  +  Key ID: ` + "`" + `` + "`" + `1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  +  Key ARN: ` + "`" + `` + "`" + `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  \n To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::KMS::Alias` + "`" + `` + "`" + ` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).\n  Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.\n  Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *Developer Guide*.\n When specifying an alias, observe the following rules.\n  +  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.\n  +  The alias and its associated KMS key must be in the same AWS-account and Region.\n  +  The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an ` + "`" + `` + "`" + `alias/projectKey` + "`" + `` + "`" + ` in multiple Regions, each of which is associated with a KMS key in its Region.\n  +  Each alias name must begin with ` + "`" + `` + "`" + `alias/` + "`" + `` + "`" + ` followed by a name, such as ` + "`" + `` + "`" + `alias/exampleKey` + "`" + `` + "`" + `. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with ` + "`" + `` + "`" + `alias/aws/` + "`" + `` + "`" + `. That alias name prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).\n  \n  *Regions* \n  KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKmsAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsAlias), &result)
	return &result
}
