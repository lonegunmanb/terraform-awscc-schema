package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2RoutingRule = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The resulting action based on matching a routing rules condition. Only InvokeApi is supported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "invoke_api": {
              "description": "Represents an InvokeApi action.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_id": {
                    "description": "The API identifier of the target API.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stage": {
                    "description": "The name of the target stage.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "strip_base_path": {
                    "computed": true,
                    "description": "The strip base path setting. When true, API Gateway strips the incoming matched base path when forwarding the request to the target API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "conditions": {
        "description": "The conditions of the routing rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "match_base_paths": {
              "computed": true,
              "description": "The base path to be matched.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "any_of": {
                    "computed": true,
                    "description": "The string of the case sensitive base path to be matched.",
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
              "optional": true
            },
            "match_headers": {
              "computed": true,
              "description": "The headers to be matched.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "any_of": {
                    "computed": true,
                    "description": "The header name and header value glob to be matched. The matchHeaders condition is matched if any of the header name and header value globs are matched.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "computed": true,
                          "description": "The case insensitive header name to be matched. The header name must be less than 40 characters and the only allowed characters are ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, and the following special characters: ` + "`" + `` + "`" + `*?-!#$%\u0026'.^_` + "`" + `|~.` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value_glob": {
                          "computed": true,
                          "description": "The case sensitive header glob value to be matched against entire header value. The header glob value must be less than 128 characters and the only allowed characters are ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, and the following special characters: ` + "`" + `` + "`" + `*?-!#$%\u0026'.^_` + "`" + `|~` + "`" + `` + "`" + `. Wildcard matching is supported for header glob values but must be for ` + "`" + `` + "`" + `*prefix-match` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `suffix-match*` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `*infix*-match` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "domain_name_arn": {
        "description": "The ARN of the domain name.",
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
      "priority": {
        "description": "The order in which API Gateway evaluates a rule. Priority is evaluated from the lowest value to the highest value. Rules can't have the same priority. Priority values 1-1,000,000 are supported.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "routing_rule_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "routing_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a routing rule. When the incoming request to a domain name matches the conditions for a rule, API Gateway invokes a stage of a target API. Supported only for REST APIs.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2RoutingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2RoutingRule), &result)
	return &result
}
