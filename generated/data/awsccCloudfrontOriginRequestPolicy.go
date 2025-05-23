package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontOriginRequestPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "origin_request_policy_config": {
        "computed": true,
        "description": "The origin request policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the origin request policy. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "cookies_config": {
              "computed": true,
              "description": "The cookies from viewer requests to include in origin requests.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookie_behavior": {
                    "computed": true,
                    "description": "Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No cookies in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any cookies that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + `*are* included in origin requests.\n  +  ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the cookies in viewer requests that are listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cookies": {
                    "computed": true,
                    "description": "Contains a list of cookie names.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "headers_config": {
              "computed": true,
              "description": "The HTTP headers to include in origin requests. These can include headers from viewer requests and additional headers added by CloudFront.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_behavior": {
                    "computed": true,
                    "description": "Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No HTTP headers in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any headers that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + `*are* included in origin requests.\n  +  ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the HTTP headers that are listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `allViewer` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `allViewerAndWhitelistCloudFront` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.\n  +  ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "headers": {
                    "computed": true,
                    "description": "Contains a list of HTTP header names.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "name": {
              "computed": true,
              "description": "A unique name to identify the origin request policy.",
              "description_kind": "plain",
              "type": "string"
            },
            "query_strings_config": {
              "computed": true,
              "description": "The URL query strings from viewer requests to include in origin requests.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_string_behavior": {
                    "computed": true,
                    "description": "Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any query strings that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + `*are* included in origin requests.\n  +  ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the query strings in viewer requests that are listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "query_strings": {
                    "computed": true,
                    "description": "Contains a list of query string names.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "origin_request_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::OriginRequestPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontOriginRequestPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginRequestPolicy), &result)
	return &result
}
