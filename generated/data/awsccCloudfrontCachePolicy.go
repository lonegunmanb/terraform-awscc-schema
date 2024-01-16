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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_ttl": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_ttl": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_ttl": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "parameters_in_cache_key_and_forwarded_to_origin": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cookies_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookie_behavior": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "cookies": {
                          "computed": true,
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
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "enable_accept_encoding_gzip": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "headers_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header_behavior": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "headers": {
                          "computed": true,
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
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "query_string_behavior": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "query_strings": {
                          "computed": true,
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
