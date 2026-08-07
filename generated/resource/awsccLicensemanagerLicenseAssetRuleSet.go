package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerLicenseAssetRuleSet = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "License asset ruleset description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_asset_ruleset_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the license asset ruleset.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "License asset ruleset name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "description": "License asset rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rule_statement": {
              "description": "Rule statement. Specify exactly one of InstanceRuleStatement, LicenseRuleStatement, or LicenseConfigurationRuleStatement.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_rule_statement": {
                    "computed": true,
                    "description": "Instance rule statement.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and_rule_statement": {
                          "computed": true,
                          "description": "AND rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                        },
                        "matching_rule_statement": {
                          "computed": true,
                          "description": "Matching rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constraint": {
                                "computed": true,
                                "description": "Constraint (e.g. Equals, Not_Equals).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                        "or_rule_statement": {
                          "computed": true,
                          "description": "OR rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "license_configuration_rule_statement": {
                    "computed": true,
                    "description": "License configuration rule statement.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and_rule_statement": {
                          "computed": true,
                          "description": "AND rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                        },
                        "matching_rule_statement": {
                          "computed": true,
                          "description": "Matching rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constraint": {
                                "computed": true,
                                "description": "Constraint (e.g. Equals, Not_Equals).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                        "or_rule_statement": {
                          "computed": true,
                          "description": "OR rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "license_rule_statement": {
                    "computed": true,
                    "description": "License rule statement.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and_rule_statement": {
                          "computed": true,
                          "description": "AND rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                        },
                        "matching_rule_statement": {
                          "computed": true,
                          "description": "Matching rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constraint": {
                                "computed": true,
                                "description": "Constraint (e.g. Equals, Not_Equals).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                        "or_rule_statement": {
                          "computed": true,
                          "description": "OR rule statement.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "matching_rule_statements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constraint": {
                                      "computed": true,
                                      "description": "Constraint (e.g. Equals, Not_Equals).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                      "nesting_mode": "single"
                    },
                    "optional": true
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
      "tags": {
        "computed": true,
        "description": "Tags to add to the license asset ruleset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::LicenseManager::LicenseAssetRuleSet.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLicensemanagerLicenseAssetRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicenseAssetRuleSet), &result)
	return &result
}
