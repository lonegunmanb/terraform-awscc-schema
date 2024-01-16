package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayVpcLink = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the VPC link.",
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
        "description": "A name for the VPC link.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
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
      "target_arns": {
        "computed": true,
        "description": "The ARN of network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_link_id": {
        "computed": true,
        "description": "The ID of the instance that backs VPC link.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::VpcLink",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayVpcLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayVpcLink), &result)
	return &result
}
