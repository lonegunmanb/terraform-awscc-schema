package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the public key to be imported. The value does not need to be unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_material": {
        "computed": true,
        "description": "The public portion of a customer-generated key pair.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IVS::PublicKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIvsPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsPublicKey), &result)
	return &result
}
