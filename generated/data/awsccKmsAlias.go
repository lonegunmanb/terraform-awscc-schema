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
        "description": "Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.",
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
        "description": "Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
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
