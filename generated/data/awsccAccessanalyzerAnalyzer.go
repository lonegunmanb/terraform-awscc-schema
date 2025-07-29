package data

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
            "internal_access_configuration": {
              "computed": true,
              "description": "Specifies the configuration of an internal access analyzer for an AWS organization or account. This configuration determines how the analyzer evaluates internal access within your AWS environment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "internal_access_analysis_rule": {
                    "computed": true,
                    "description": "Contains information about analysis rules for the internal access analyzer. Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "inclusions": {
                          "computed": true,
                          "description": "A list of rules for the internal access analyzer containing criteria to include in analysis. Only resources that meet the rule criteria will generate findings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_ids": {
                                "computed": true,
                                "description": "A list of AWS account IDs to apply to the internal access analysis rule criteria. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers and cannot include the organization owner account.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "resource_arns": {
                                "computed": true,
                                "description": "A list of resource ARNs to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources that match these ARNs.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "resource_types": {
                                "computed": true,
                                "description": "A list of resource types to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources of these types.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
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
            "unused_access_configuration": {
              "computed": true,
              "description": "The Configuration for Unused Access Analyzer",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "analysis_rule": {
                    "computed": true,
                    "description": "Contains information about rules for the analyzer.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "exclusions": {
                          "computed": true,
                          "description": "A list of rules for the analyzer containing criteria to exclude from analysis. Entities that meet the rule criteria will not generate findings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_ids": {
                                "computed": true,
                                "description": "A list of AWS account IDs to apply to the analysis rule criteria. The accounts cannot include the organization analyzer owner account. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "resource_tags": {
                                "computed": true,
                                "description": "An array of key-value pairs to match for your resources. You can use the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.\n\nFor the tag key, you can specify a value that is 1 to 128 characters in length and cannot be prefixed with aws:.\n\nFor the tag value, you can specify a value that is 0 to 256 characters in length. If the specified tag value is 0 characters, the rule is applied to all principals with the specified tag key.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "unused_access_age": {
                    "computed": true,
                    "description": "The specified access age in days for which to generate findings for unused access. For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 365 days.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "analyzer_name": {
        "computed": true,
        "description": "Analyzer name",
        "description_kind": "plain",
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
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "eq": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exists": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "neq": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "property": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "rule_name": {
              "computed": true,
              "description": "The archive rule name",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the analyzer",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_INTERNAL_ACCESS, ORGANIZATION_INTERNAL_ACCESS, ACCOUNT_UNUSED_ACCESS and ORGANIZATION_UNUSED_ACCESS",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AccessAnalyzer::Analyzer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAccessanalyzerAnalyzerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccessanalyzerAnalyzer), &result)
	return &result
}
