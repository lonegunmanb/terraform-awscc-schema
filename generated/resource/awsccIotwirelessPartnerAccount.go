package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessPartnerAccount = `{
  "block": {
    "attributes": {
      "account_linked": {
        "computed": true,
        "description": "Whether the partner account is linked to the AWS account.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description": "PartnerAccount arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "The fingerprint of the Sidewalk application server private key.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "partner_account_id": {
        "computed": true,
        "description": "The partner account ID to disassociate from the AWS account",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partner_type": {
        "computed": true,
        "description": "The partner type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sidewalk": {
        "computed": true,
        "description": "The Sidewalk account credentials.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_server_private_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sidewalk_response": {
        "computed": true,
        "description": "The Sidewalk account credentials.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fingerprint": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sidewalk_update": {
        "computed": true,
        "description": "The Sidewalk account credentials.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_server_private_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the destination.",
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
    "description": "Create and manage partner account",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessPartnerAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessPartnerAccount), &result)
	return &result
}
