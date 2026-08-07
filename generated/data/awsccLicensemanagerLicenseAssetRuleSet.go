package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "license_asset_ruleset_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the license asset ruleset.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "License asset ruleset name.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "License asset rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rule_statement": {
              "computed": true,
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                                "type": "string"
                              },
                              "key_to_match": {
                                "computed": true,
                                "description": "Key to match.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value_to_match": {
                                "computed": true,
                                "description": "Values to match.",
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
                                      "type": "string"
                                    },
                                    "key_to_match": {
                                      "computed": true,
                                      "description": "Key to match.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_to_match": {
                                      "computed": true,
                                      "description": "Values to match.",
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
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::LicenseManager::LicenseAssetRuleSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerLicenseAssetRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicenseAssetRuleSet), &result)
	return &result
}
