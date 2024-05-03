package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2KeyPair = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": "string"
      },
      "key_name": {
        "computed": true,
        "description": "A unique name for the key pair.\n Constraints: Up to 255 ASCII characters",
        "description_kind": "plain",
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
        "type": "string"
      },
      "public_key_material": {
        "computed": true,
        "description": "The public key material. The ` + "`" + `` + "`" + `PublicKeyMaterial` + "`" + `` + "`" + ` property is used to import a key pair. If this property is not specified, then a new key pair will be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the key pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::KeyPair",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2KeyPairSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2KeyPair), &result)
	return &result
}
