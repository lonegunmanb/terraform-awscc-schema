package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferHostKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The unique Amazon Resource Name (ARN) for the host key.",
        "description_kind": "plain",
        "type": "string"
      },
      "date_imported": {
        "computed": true,
        "description": "The date on which the host key was added to the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The text description for this host key.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_key_body": {
        "computed": true,
        "description": "The private key portion of an SSH key pair. Transfer Family accepts RSA, ECDSA, and ED25519 keys.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_key_fingerprint": {
        "computed": true,
        "description": "The public key fingerprint, which is a short sequence of bytes used to identify the longer public key.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_key_id": {
        "computed": true,
        "description": "A unique identifier for the host key.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_id": {
        "computed": true,
        "description": "The identifier of the server that contains the host key.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for host keys.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The encryption algorithm that is used for the host key.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Transfer::HostKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferHostKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferHostKey), &result)
	return &result
}
