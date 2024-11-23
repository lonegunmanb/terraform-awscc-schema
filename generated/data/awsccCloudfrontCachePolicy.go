package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontCachePolicy = `{
  "block": {
    "attributes": {
      "cache_policy_config": {
        "computed": true,
        "description": "The cache policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the cache policy. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "default_ttl": {
              "computed": true,
              "description": "The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send ` + "`" + `` + "`" + `Cache-Control` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 86400 seconds (one day). If the value of ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` is more than 86400 seconds, then the default value for this field is the same as the value of ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_ttl": {
              "computed": true,
              "description": "The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value only when the origin sends ` + "`" + `` + "`" + `Cache-Control` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 31536000 seconds (one year). If the value of ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DefaultTTL` + "`" + `` + "`" + ` is more than 31536000 seconds, then the default value for this field is the same as the value of ` + "`" + `` + "`" + `DefaultTTL` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_ttl": {
              "computed": true,
              "description": "The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.",
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "A unique name to identify the cache policy.",
              "description_kind": "plain",
              "type": "string"
            },
            "parameters_in_cache_key_and_forwarded_to_origin": {
              "computed": true,
              "description": "The HTTP headers, cookies, and URL query strings to include in the cache key. The values included in the cache key are also included in requests that CloudFront sends to the origin.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookies_config": {
                    "computed": true,
                    "description": "An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and in requests that CloudFront sends to the origin.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookie_behavior": {
                          "computed": true,
                          "description": "Determines whether any cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No cookies in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any cookies that are listed in an ` + "`" + `` + "`" + `OriginRequestPolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the cookies in viewer requests that are listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* for those that are listed in the ` + "`" + `` + "`" + `CookieNames` + "`" + `` + "`" + ` type, which are not included.\n  +   ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
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
                  "enable_accept_encoding_brotli": {
                    "computed": true,
                    "description": "A flag that can affect whether the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ` + "`" + `` + "`" + `EnableAcceptEncodingGzip` + "`" + `` + "`" + ` field. If one or both of these fields is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` *and* the viewer request includes the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and this cache behavior also has an origin request policy attached, do not include the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header in the origin request policy. CloudFront always includes the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header in origin requests when the value of this field is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, so including this header in an origin request policy has no effect.\n If both of these fields are ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, then CloudFront treats the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` to the headers whitelist like any other HTTP header.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "enable_accept_encoding_gzip": {
                    "computed": true,
                    "description": "A flag that can affect whether the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ` + "`" + `` + "`" + `EnableAcceptEncodingBrotli` + "`" + `` + "`" + ` field. If one or both of these fields is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` *and* the viewer request includes the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and this cache behavior also has an origin request policy attached, do not include the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header in the origin request policy. CloudFront always includes the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header in origin requests when the value of this field is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, so including this header in an origin request policy has no effect.\n If both of these fields are ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, then CloudFront treats the ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ` + "`" + `` + "`" + `Accept-Encoding` + "`" + `` + "`" + ` to the headers whitelist like any other HTTP header.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "headers_config": {
                    "computed": true,
                    "description": "An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and in requests that CloudFront sends to the origin.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header_behavior": {
                          "computed": true,
                          "description": "Determines whether any HTTP headers are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No HTTP headers are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any headers that are listed in an ` + "`" + `` + "`" + `OriginRequestPolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the HTTP headers that are listed in the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + ` type are included in the cache key and in requests that CloudFront sends to the origin.",
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
                  "query_strings_config": {
                    "computed": true,
                    "description": "An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and in requests that CloudFront sends to the origin.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "query_string_behavior": {
                          "computed": true,
                          "description": "Determines whether any URL query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` ? No query strings in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, any query strings that are listed in an ` + "`" + `` + "`" + `OriginRequestPolicy` + "`" + `` + "`" + ` *are* included in origin requests.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` ? Only the query strings in viewer requests that are listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +   ` + "`" + `` + "`" + `allExcept` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* those that are listed in the ` + "`" + `` + "`" + `QueryStringNames` + "`" + `` + "`" + ` type, which are not included.\n  +   ` + "`" + `` + "`" + `all` + "`" + `` + "`" + ` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "cache_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
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
      }
    },
    "description": "Data Source schema for AWS::CloudFront::CachePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontCachePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontCachePolicy), &result)
	return &result
}
