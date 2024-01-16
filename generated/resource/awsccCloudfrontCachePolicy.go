package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontCachePolicy = `{
  "block": {
    "attributes": {
      "cache_policy_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_ttl": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "max_ttl": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_ttl": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameters_in_cache_key_and_forwarded_to_origin": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
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
                  "enable_accept_encoding_brotli": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_accept_encoding_gzip": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::CachePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontCachePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontCachePolicy), &result)
	return &result
}
