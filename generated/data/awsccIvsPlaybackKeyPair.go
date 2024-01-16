package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsPlaybackKeyPair = `{
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
        "description": "An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.",
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
    "description": "Data Source schema for AWS::IVS::PlaybackKeyPair",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIvsPlaybackKeyPairSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsPlaybackKeyPair), &result)
	return &result
}
