package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2RegexPatternSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of the WAF entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the entity.",
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
        "description": "Name of the RegexPatternSet.",
        "description_kind": "plain",
        "type": "string"
      },
      "regex_pattern_set_id": {
        "computed": true,
        "description": "Id of the RegexPatternSet",
        "description_kind": "plain",
        "type": "string"
      },
      "regular_expression_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scope": {
        "computed": true,
        "description": "Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.",
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
    "description": "Data Source schema for AWS::WAFv2::RegexPatternSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWafv2RegexPatternSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2RegexPatternSet), &result)
	return &result
}
