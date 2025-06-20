package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubAutomationRuleV2 = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "A list of actions to be performed when the rule criteria is met",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_integration_configuration": {
              "computed": true,
              "description": "The settings for integrating automation rule actions with external systems or service",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connector_arn": {
                    "computed": true,
                    "description": "The ARN of the connector that establishes the integration",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "finding_fields_update": {
              "computed": true,
              "description": "The changes to be applied to fields in a security finding when an automation rule is triggered",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comment": {
                    "computed": true,
                    "description": "Notes or contextual information for findings that are modified by the automation rule",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "severity_id": {
                    "computed": true,
                    "description": "The severity level to be assigned to findings that match the automation rule criteria",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status_id": {
                    "computed": true,
                    "description": "The status to be applied to findings that match automation rule criteria",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "description": "The category of action to be executed by the automation rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp formatted in ISO8601",
        "description_kind": "plain",
        "type": "string"
      },
      "criteria": {
        "description": "Defines the parameters and conditions used to evaluate and filter security findings",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ocsf_finding_criteria": {
              "computed": true,
              "description": "The filtering conditions that align with OCSF standards",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "composite_filters": {
                    "computed": true,
                    "description": "Enables the creation of complex filtering conditions by combining filter",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "boolean_filters": {
                          "computed": true,
                          "description": "Enables filtering based on boolean field values",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_name": {
                                "computed": true,
                                "description": "The name of the field",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter": {
                                "computed": true,
                                "description": "Boolean filter for querying findings",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "value": {
                                      "computed": true,
                                      "description": "The value of the boolean",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "date_filters": {
                          "computed": true,
                          "description": "Enables filtering based on date and timestamp fields",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_name": {
                                "computed": true,
                                "description": "The name of the field",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter": {
                                "computed": true,
                                "description": "A date filter for querying findings",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "date_range": {
                                      "computed": true,
                                      "description": "A date range for the date filter",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "unit": {
                                            "computed": true,
                                            "description": "A date range unit for the date filter",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "A date range value for the date filter",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "end": {
                                      "computed": true,
                                      "description": "The timestamp formatted in ISO8601",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "start": {
                                      "computed": true,
                                      "description": "The timestamp formatted in ISO8601",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "map_filters": {
                          "computed": true,
                          "description": "Enables filtering based on map field value",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_name": {
                                "computed": true,
                                "description": "The name of the field",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter": {
                                "computed": true,
                                "description": "A map filter for filtering findings",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "comparison": {
                                      "computed": true,
                                      "description": "The condition to apply to the key value when filtering findings with a map filter",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "computed": true,
                                      "description": "The key of the map filter",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The value for the key in the map filter",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "number_filters": {
                          "computed": true,
                          "description": "Enables filtering based on numerical field values",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_name": {
                                "computed": true,
                                "description": "The name of the field",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter": {
                                "computed": true,
                                "description": "A number filter for querying findings",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "eq": {
                                      "computed": true,
                                      "description": "The equal-to condition to be applied to a single field when querying for findings",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "gte": {
                                      "computed": true,
                                      "description": "The greater-than-equal condition to be applied to a single field when querying for findings",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "lte": {
                                      "computed": true,
                                      "description": "The less-than-equal condition to be applied to a single field when querying for findings",
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
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "operator": {
                          "computed": true,
                          "description": "The logical operator used to combine multiple conditions",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "string_filters": {
                          "computed": true,
                          "description": "Enables filtering based on string field values",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_name": {
                                "computed": true,
                                "description": "The name of the field",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter": {
                                "computed": true,
                                "description": "A string filter for filtering findings",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "comparison": {
                                      "computed": true,
                                      "description": "The condition to apply to a string value when filtering findings",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The string filter value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "composite_operator": {
                    "computed": true,
                    "description": "The logical operator used to combine multiple conditions",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      },
      "description": {
        "description": "A description of the automation rule",
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
      "rule_arn": {
        "computed": true,
        "description": "The ARN of the automation rule",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description": "The ID of the automation rule",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description": "The name of the automation rule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_order": {
        "description": "The value for the rule priority",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_status": {
        "computed": true,
        "description": "The status of the automation rule",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp formatted in ISO8601",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::SecurityHub::AutomationRuleV2",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubAutomationRuleV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubAutomationRuleV2), &result)
	return &result
}
