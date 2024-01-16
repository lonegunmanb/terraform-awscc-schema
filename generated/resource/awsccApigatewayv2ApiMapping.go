package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2ApiMapping = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The identifier of the API.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_mapping_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_mapping_key": {
        "computed": true,
        "description": "The API mapping key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "The domain name.",
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
      "stage": {
        "description": "The API stage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGatewayV2::ApiMapping` + "`" + `` + "`" + ` resource contains an API mapping. An API mapping relates a path of your custom domain name to a stage of your API. A custom domain name can have multiple API mappings, but the paths can't overlap. A custom domain can map only to APIs of the same protocol type. For more information, see [CreateApiMapping](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html#CreateApiMapping) in the *Amazon API Gateway V2 API Reference*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2ApiMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2ApiMapping), &result)
	return &result
}
