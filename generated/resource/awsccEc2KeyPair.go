package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2KeyPair = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_format": {
        "computed": true,
        "description": "The format of the key pair.\n Default: ` + "`" + `` + "`" + `pem` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_name": {
        "description": "A unique name for the key pair.\n Constraints: Up to 255 ASCII characters",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_pair_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_type": {
        "computed": true,
        "description": "The type of key pair. Note that ED25519 keys are not supported for Windows instances.\n If the ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` property is specified, the ` + "`" + `` + "`" + `KeyType` + "`" + `` + "`" + ` property is ignored, and the key type is inferred from the ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` value.\n Default: ` + "`" + `` + "`" + `rsa` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key_material": {
        "computed": true,
        "description": "The public key material. The ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` property is used to import a key pair. If this property is not specified, then a new key pair will be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the key pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Specifies a key pair for use with an EC2long instance as follows:\n  +  To import an existing key pair, include the ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` property.\n  +  To create a new key pair, omit the ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` property.\n  \n When you import an existing key pair, you specify the public key material for the key. We assume that you have the private key material for the key. CFNlong does not create or return the private key material when you import a key pair.\n When you create a new key pair, the private key is saved to SYSlong Parameter Store, using a parameter with the following name: ` + "`" + `` + "`" + `/ec2/keypair/{key_pair_id}` + "`" + `` + "`" + `. For more information about retrieving private key, and the required permissions, see [Create a key pair using](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-key-pairs.html#create-key-pair-cloudformation) in the *User Guide*.\n When CFN deletes a key pair that was created or imported by a stack, it also deletes the parameter that was used to store the private key material in Parameter Store.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2KeyPairSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2KeyPair), &result)
	return &result
}
