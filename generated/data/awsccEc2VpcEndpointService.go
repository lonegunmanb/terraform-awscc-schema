package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEndpointService = `{
  "block": {
    "attributes": {
      "acceptance_required": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "contributor_insights_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "gateway_load_balancer_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_load_balancer_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "payer_responsibility": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to add to the VPC endpoint service.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::VPCEndpointService",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcEndpointServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpointService), &result)
	return &result
}
