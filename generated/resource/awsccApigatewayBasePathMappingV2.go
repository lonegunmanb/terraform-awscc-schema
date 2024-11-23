package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayBasePathMappingV2 = `{
  "block": {
    "attributes": {
      "base_path": {
        "computed": true,
        "description": "The base path name that callers of the API must provide in the URL after the domain name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_path_mapping_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_arn": {
        "description": "The Arn of an AWS::ApiGateway::DomainNameV2 resource.",
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
      "rest_api_id": {
        "description": "The ID of the API.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description": "The name of the API's stage.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ApiGateway::BasePathMappingV2",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayBasePathMappingV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayBasePathMappingV2), &result)
	return &result
}
