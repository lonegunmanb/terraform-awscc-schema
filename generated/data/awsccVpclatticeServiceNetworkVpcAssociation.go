package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeServiceNetworkVpcAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "private_dns_preference": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_dns_specified_domains": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_dns_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "service_network_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_vpc_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::VpcLattice::ServiceNetworkVpcAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVpclatticeServiceNetworkVpcAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeServiceNetworkVpcAssociation), &result)
	return &result
}
