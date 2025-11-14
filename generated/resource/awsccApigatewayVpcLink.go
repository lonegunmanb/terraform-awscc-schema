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
              "description": "A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
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
