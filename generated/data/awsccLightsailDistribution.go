package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDistribution = `{
  "block": {
    "attributes": {
      "able_to_update_bundle": {
        "computed": true,
        "description": "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle.",
        "description_kind": "plain",
        "type": "bool"
      },
      "bundle_id": {
        "computed": true,
        "description": "The bundle ID to use for the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_behavior_settings": {
        "computed": true,
        "description": "An object that describes the cache behavior settings for the distribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_http_methods": {
              "computed": true,
              "description": "The HTTP methods that are processed and forwarded to the distribution's origin.",
              "description_kind": "plain",
              "type": "string"
            },
            "cached_http_methods": {
              "computed": true,
              "description": "The HTTP method responses that are cached by your distribution.",
              "description_kind": "plain",
              "type": "string"
            },
            "default_ttl": {
              "computed": true,
              "description": "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
              "description_kind": "plain",
              "type": "number"
            },
            "forwarded_cookies": {
              "computed": true,
              "description": "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookies_allow_list": {
                    "computed": true,
                    "description": "The specific cookies to forward to your distribution's origin.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "option": {
                    "computed": true,
                    "description": "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "forwarded_headers": {
              "computed": true,
              "description": "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "headers_allow_list": {
                    "computed": true,
                    "description": "The specific headers to forward to your distribution's origin.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "option": {
                    "computed": true,
                    "description": "The headers that you want your distribution to forward to your origin and base caching on.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "forwarded_query_strings": {
              "computed": true,
              "description": "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "option": {
                    "computed": true,
                    "description": "Indicates whether the distribution forwards and caches based on query strings.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "query_strings_allow_list": {
                    "computed": true,
                    "description": "The specific query strings that the distribution forwards to the origin.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maximum_ttl": {
              "computed": true,
              "description": "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_ttl": {
              "computed": true,
              "description": "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cache_behaviors": {
        "computed": true,
        "description": "An array of objects that describe the per-path cache behavior for the distribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "behavior": {
              "computed": true,
              "description": "The cache behavior for the specified path.",
              "description_kind": "plain",
              "type": "string"
            },
            "path": {
              "computed": true,
              "description": "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "certificate_name": {
        "computed": true,
        "description": "The certificate attached to the Distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_cache_behavior": {
        "computed": true,
        "description": "An object that describes the default cache behavior for the distribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "behavior": {
              "computed": true,
              "description": "The cache behavior of the distribution.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "distribution_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_name": {
        "computed": true,
        "description": "The name for the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description": "The IP address type for the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_enabled": {
        "computed": true,
        "description": "Indicates whether the distribution is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "origin": {
        "computed": true,
        "description": "An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the origin resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "protocol_policy": {
              "computed": true,
              "description": "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The AWS Region name of the origin resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Lightsail::Distribution",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDistribution), &result)
	return &result
}
