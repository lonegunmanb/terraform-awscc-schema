package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesCaseRule = `{
  "block": {
    "attributes": {
      "case_rule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of a case rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "case_rule_id": {
        "computed": true,
        "description": "The unique identifier of a case rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time at which the case rule was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description explaining the purpose and behavior of this case rule. Helps administrators understand when and why this rule applies to case fields.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique identifier of the Cases domain.",
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
        "description": "The time at which the case rule was created or last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A descriptive name for the case rule. Must be unique within the domain and should clearly indicate the rule's purpose (e.g., 'Priority Field Required for Urgent Cases').",
        "description_kind": "plain",
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description": "Defines the rule behavior and conditions. Specifies the rule type and the conditions under which it applies. In the Amazon Connect admin website, this corresponds to case field conditions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hidden": {
              "computed": true,
              "description": "Hidden rule type, used to indicate whether a field is hidden",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "conditions": {
                    "computed": true,
                    "description": "List of conditions for the hidden rule; the first condition to evaluate to true dictates the value of the rule",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "equal_to": {
                          "computed": true,
                          "description": "Boolean operands for a condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "operand_one": {
                                "computed": true,
                                "description": "The left hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "field_id": {
                                      "computed": true,
                                      "description": "The field ID this operand should take the value of.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "operand_two": {
                                "computed": true,
                                "description": "The right hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "boolean_value": {
                                      "computed": true,
                                      "description": "A boolean value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "double_value": {
                                      "computed": true,
                                      "description": "A numeric value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "empty_value": {
                                      "computed": true,
                                      "description": "An empty operand value.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "string_value": {
                                      "computed": true,
                                      "description": "A string value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "result": {
                                "computed": true,
                                "description": "The value of the outer rule if the condition evaluates to true.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "not_equal_to": {
                          "computed": true,
                          "description": "Boolean operands for a condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "operand_one": {
                                "computed": true,
                                "description": "The left hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "field_id": {
                                      "computed": true,
                                      "description": "The field ID this operand should take the value of.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "operand_two": {
                                "computed": true,
                                "description": "The right hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "boolean_value": {
                                      "computed": true,
                                      "description": "A boolean value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "double_value": {
                                      "computed": true,
                                      "description": "A numeric value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "empty_value": {
                                      "computed": true,
                                      "description": "An empty operand value.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "string_value": {
                                      "computed": true,
                                      "description": "A string value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "result": {
                                "computed": true,
                                "description": "The value of the outer rule if the condition evaluates to true.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "default_value": {
                    "computed": true,
                    "description": "The value of the rule (i.e. whether the field is hidden) should none of the conditions evaluate to true",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "required": {
              "computed": true,
              "description": "A required rule type, used to indicate whether a field is required.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "conditions": {
                    "computed": true,
                    "description": "An ordered list of boolean conditions that determine when the field should be required. Conditions are evaluated in order, and the first condition that evaluates to true determines whether the field is required, overriding the default value.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "equal_to": {
                          "computed": true,
                          "description": "Boolean operands for a condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "operand_one": {
                                "computed": true,
                                "description": "The left hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "field_id": {
                                      "computed": true,
                                      "description": "The field ID this operand should take the value of.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "operand_two": {
                                "computed": true,
                                "description": "The right hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "boolean_value": {
                                      "computed": true,
                                      "description": "A boolean value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "double_value": {
                                      "computed": true,
                                      "description": "A numeric value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "empty_value": {
                                      "computed": true,
                                      "description": "An empty operand value.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "string_value": {
                                      "computed": true,
                                      "description": "A string value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "result": {
                                "computed": true,
                                "description": "The value of the outer rule if the condition evaluates to true.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "not_equal_to": {
                          "computed": true,
                          "description": "Boolean operands for a condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "operand_one": {
                                "computed": true,
                                "description": "The left hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "field_id": {
                                      "computed": true,
                                      "description": "The field ID this operand should take the value of.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "operand_two": {
                                "computed": true,
                                "description": "The right hand operand in the condition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "boolean_value": {
                                      "computed": true,
                                      "description": "A boolean value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "double_value": {
                                      "computed": true,
                                      "description": "A numeric value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "empty_value": {
                                      "computed": true,
                                      "description": "An empty operand value.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "string_value": {
                                      "computed": true,
                                      "description": "A string value to compare against the field value in the condition evaluation.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "result": {
                                "computed": true,
                                "description": "The value of the outer rule if the condition evaluates to true.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "default_value": {
                    "computed": true,
                    "description": "The default required state for the field when none of the specified conditions are met. If true, the field is required by default; if false, the field is optional by default.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this case rule.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Cases::CaseRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCasesCaseRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesCaseRule), &result)
	return &result
}
