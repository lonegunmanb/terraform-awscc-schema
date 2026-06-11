package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRtbfabricLinkRoutingRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "conditions": {
        "description": "Conditions for a routing rule. All non-null fields must match (AND logic). At least one field must be set. HostHeader and HostHeaderWildcard are mutually exclusive. PathPrefix and PathExact are mutually exclusive.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "host_header": {
              "computed": true,
              "description": "Exact host match ? RFC 3986 unreserved characters. Mutually exclusive with HostHeaderWildcard.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_header_wildcard": {
              "computed": true,
              "description": "Wildcard host pattern (e.g., *.example.com) ? RFC 3986 unreserved characters plus *. Mutually exclusive with HostHeader.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path_exact": {
              "computed": true,
              "description": "Exact path match ? must start with /. Mutually exclusive with PathPrefix.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path_prefix": {
              "computed": true,
              "description": "Path prefix matching ? strict starts-with, must start with /. Mutually exclusive with PathExact.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_string_equals": {
              "computed": true,
              "description": "Query string key=value pair match (single pair).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "Query string key ? RFC 3986 unreserved characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "Query string value ? RFC 3986 unreserved characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "query_string_exists": {
              "computed": true,
              "description": "Query string key presence check (any value accepted).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_id": {
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
      "link_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the LinkRoutingRule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::RTBFabric::LinkRoutingRule. A routing rule on a link within RTB Fabric that controls request routing based on conditions such as host headers, path matching, and query string parameters.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRtbfabricLinkRoutingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRtbfabricLinkRoutingRule), &result)
	return &result
}
