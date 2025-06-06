package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallVpcEndpointAssociation = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_id": {
        "computed": true,
        "description": "An endpoint Id.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_arn": {
        "description": "A resource ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_mapping": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description": "A IPAddressType",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description": "A SubnetId.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
      },
      "vpc_endpoint_association_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallVpcEndpointAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallVpcEndpointAssociation), &result)
	return &result
}
