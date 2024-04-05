package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontOriginRequestPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "origin_request_policy_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cookies_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookie_behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cookies": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "headers_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "headers": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "query_strings_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_string_behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "query_strings": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "origin_request_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::OriginRequestPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontOriginRequestPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginRequestPolicy), &result)
	return &result
}
