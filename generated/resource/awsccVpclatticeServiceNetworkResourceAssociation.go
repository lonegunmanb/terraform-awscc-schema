package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_network_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
    "description": "VpcLattice ServiceNetworkResourceAssociation CFN resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVpclatticeServiceNetworkResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeServiceNetworkResourceAssociation), &result)
	return &result
}
