package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbOdbPeeringConnection = `{
  "block": {
    "attributes": {
      "display_name": {
        "computed": true,
        "description": "The name of the ODB peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The unique identifier of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_peering_connection_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ODB peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_peering_connection_id": {
        "computed": true,
        "description": "The unique identifier of the ODB peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the peer network.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_network_id": {
        "computed": true,
        "description": "The unique identifier of the peer network.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Odb peering connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and \".",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ODB::OdbPeeringConnection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOdbOdbPeeringConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbOdbPeeringConnection), &result)
	return &result
}
