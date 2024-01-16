package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2VpcLink = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the VPC link.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "A list of security group IDs for the VPC link.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description": "A list of subnet IDs to include in the VPC link.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_link_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::VpcLink",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2VpcLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2VpcLink), &result)
	return &result
}
