package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecondaryNetwork = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_cidr_block": {
        "description": "The IPv4 CIDR block for the secondary network. The CIDR block size must be between /12 and /28.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv_4_cidr_block_associations": {
        "computed": true,
        "description": "Information about the IPv4 CIDR blocks associated with the secondary network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "association_id": {
              "computed": true,
              "description": "The association ID for the IPv4 CIDR block.",
              "description_kind": "plain",
              "type": "string"
            },
            "cidr_block": {
              "computed": true,
              "description": "The IPv4 CIDR block.",
              "description_kind": "plain",
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "The state of the CIDR block association.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "network_type": {
        "description": "The type of secondary network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account that owns the secondary network.",
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the secondary network.",
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_network_id": {
        "computed": true,
        "description": "The ID of the secondary network.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the secondary network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
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
    "description": "Resource Type definition for AWS::EC2::SecondaryNetwork. Creates a secondary network for RDMA (Remote Direct Memory Access) connectivity.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SecondaryNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecondaryNetwork), &result)
	return &result
}
