package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAccessanalyzerAnalyzer = `{
  "block": {
    "attributes": {
      "analyzer_configuration": {
        "computed": true,
        "description": "The configuration for the analyzer",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "unused_access_configuration": {
              "computed": true,
              "description": "The Configuration for Unused Access Analyzer",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "unused_access_age": {
                    "computed": true,
                    "description": "The specified access age in days for which to generate findings for unused access. For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 180 days.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "analyzer_name": {
        "computed": true,
        "description": "Analyzer name",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "archive_rules": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "contains": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "eq": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exists": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "neq": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "property": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "rule_name": {
              "computed": true,
              "description": "The archive rule name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the analyzer",
        "description_kind": "plain",
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
        "description": "An array of key-value pairs to apply to this resource.",
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
      },
      "type": {
        "description": "The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_UNUSED_ACCESS or ORGANIZATION_UNUSED_ACCESS",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAccessanalyzerAnalyzerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccessanalyzerAnalyzer), &result)
	return &result
}
