package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeServiceNetworkResourceAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_resource_association_id": {
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
      }
    },
    "description": "Data Source schema for AWS::VpcLattice::ServiceNetworkResourceAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVpclatticeServiceNetworkResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeServiceNetworkResourceAssociation), &result)
	return &result
}
