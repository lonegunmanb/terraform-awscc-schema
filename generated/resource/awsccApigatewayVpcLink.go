package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayVpcLink = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the VPC link.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name used to label and identify the VPC link.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of arbitrary tags (key-value pairs) to associate with the VPC link.",
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
      "target_arns": {
        "description": "The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS-account of the API owner.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_link_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::VpcLink` + "`" + `` + "`" + ` resource creates an API Gateway VPC link for a REST API to access resources in an Amazon Virtual Private Cloud (VPC). For more information, see [vpclink:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateVpcLink.html) in the ` + "`" + `` + "`" + `Amazon API Gateway REST API Reference` + "`" + `` + "`" + `.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayVpcLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayVpcLink), &result)
	return &result
}
