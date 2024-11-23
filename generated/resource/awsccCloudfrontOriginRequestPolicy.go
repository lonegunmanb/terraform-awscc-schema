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
        "description": "The origin request policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the origin request policy. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cookies_config": {
              "description": "The cookies from viewer requests to include in origin requests.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookie_behavior": {
                    "description": "Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No cookies in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any cookies that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the cookies in viewer requests that are listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cookies": {
                    "computed": true,
                    "description": "Contains a list of cookie names.",
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
              "description": "The HTTP headers to include in origin requests. These can include headers from viewer requests and additional headers added by CloudFront.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_behavior": {
                    "description": "Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No HTTP headers in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any headers that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the HTTP headers that are listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allViewer` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allViewerAndWhitelistCloudFront` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.\n  +   ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "headers": {
                    "computed": true,
                    "description": "Contains a list of HTTP header names.",
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
              "description": "A unique name to identify the origin request policy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "query_strings_config": {
              "description": "The URL query strings from viewer requests to include in origin requests.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_string_behavior": {
                    "description": "Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any query strings that are listed in a ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the query strings in viewer requests that are listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type, which are not included.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "query_strings": {
                    "computed": true,
                    "description": "Contains a list of query string names.",
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
    "description": "An origin request policy.\n When it's attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:\n  +  The request body and the URL path (without the domain name) from the viewer request.\n  +  The headers that CloudFront automatically includes in every origin request, including ` + "`" + `` + "`" + `Host` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `User-Agent` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `X-Amz-Cf-Id` + "`" + `` + "`" + `.\n  +  All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.\n  \n CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use ` + "`" + `` + "`" + `CachePolicy` + "`" + `` + "`" + `.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontOriginRequestPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginRequestPolicy), &result)
	return &result
}
