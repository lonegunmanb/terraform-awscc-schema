package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsPublicKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Key-pair identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Key-pair identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the public key to be imported. The value does not need to be unique.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key_material": {
        "computed": true,
        "description": "The public portion of a customer-generated key pair. This field is required to create the AWS::IVS::PublicKey resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::IVS::PublicKey",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsPublicKey), &result)
	return &result
}
