package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigAggregationAuthorization = `{
  "block": {
    "attributes": {
      "aggregation_authorization_arn": {
        "computed": true,
        "description": "The ARN of the AggregationAuthorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorized_account_id": {
        "description": "The 12-digit account ID of the account authorized to aggregate data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorized_aws_region": {
        "description": "The region authorized to collect aggregated data.",
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
      "tags": {
        "computed": true,
        "description": "The tags for the AggregationAuthorization.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Config::AggregationAuthorization",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConfigAggregationAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigAggregationAuthorization), &result)
	return &result
}
