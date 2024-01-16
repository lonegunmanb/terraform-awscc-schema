package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2VpcLink = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the VPC link.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "A list of security group IDs for the VPC link.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "description": "A list of subnet IDs to include in the VPC link.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "optional": true,
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGatewayV2::VpcLink` + "`" + `` + "`" + ` resource creates a VPC link. Supported only for HTTP APIs. The VPC link status must transition from ` + "`" + `` + "`" + `PENDING` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `AVAILABLE` + "`" + `` + "`" + ` to successfully create a VPC link, which can take up to 10 minutes. To learn more, see [Working with VPC Links for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vpc-links.html) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2VpcLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2VpcLink), &result)
	return &result
}
