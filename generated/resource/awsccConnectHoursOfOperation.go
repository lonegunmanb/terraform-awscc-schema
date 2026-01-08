package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectHoursOfOperation = `{
  "block": {
    "attributes": {
      "child_hours_of_operations": {
        "computed": true,
        "description": "List of child hours of operations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The identifier for the hours of operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the hours of operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "config": {
        "description": "Configuration information for the hours of operation: day, start time, and end time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "day": {
              "description": "The day that the hours of operation applies to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "end_time": {
              "description": "The end time that your contact center closes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "description": "The hours.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "start_time": {
              "description": "The start time that your contact center opens.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "description": "The hours.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "The description of the hours of operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hours_of_operation_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "hours_of_operation_overrides": {
        "computed": true,
        "description": "One or more hours of operation overrides assigned to an hour of operation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "effective_from": {
              "computed": true,
              "description": "The date from which the hours of operation override would be effective.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "effective_till": {
              "computed": true,
              "description": "The date till which the hours of operation override would be effective.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hours_of_operation_override_id": {
              "computed": true,
              "description": "The Resource Identifier for the hours of operation override.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_config": {
              "computed": true,
              "description": "Configuration information for the hours of operation override: day, start time, and end time.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day": {
                    "computed": true,
                    "description": "The day that the hours of operation override applies to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "end_time": {
                    "computed": true,
                    "description": "The new end time that your contact center closes for the overriden days.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hours": {
                          "computed": true,
                          "description": "The hours.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "computed": true,
                          "description": "The minutes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "start_time": {
                    "computed": true,
                    "description": "The new start time that your contact center opens for the overriden days.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hours": {
                          "computed": true,
                          "description": "The hours.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "computed": true,
                          "description": "The minutes.",
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
                "nesting_mode": "set"
              },
              "optional": true
            },
            "override_description": {
              "computed": true,
              "description": "The description of the hours of operation override.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_name": {
              "computed": true,
              "description": "The name of the hours of operation override.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_type": {
              "computed": true,
              "description": "The type of hours of operation override.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recurrence_config": {
              "computed": true,
              "description": "Configuration for recurring hours of operation overrides.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "recurrence_pattern": {
                    "computed": true,
                    "description": "Pattern for recurring hours of operation overrides.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "by_month": {
                          "computed": true,
                          "description": "List of months (1-12) for recurrence pattern.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "by_month_day": {
                          "computed": true,
                          "description": "List of month days (-1 to 31) for recurrence pattern.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "by_weekday_occurrence": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "frequency": {
                          "computed": true,
                          "description": "The frequency of recurrence for hours of operation overrides.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "interval": {
                          "computed": true,
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
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the hours of operation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_hours_of_operations": {
        "computed": true,
        "description": "List of parent hours of operations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The identifier for the hours of operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the hours of operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "time_zone": {
        "description": "The time zone of the hours of operation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::HoursOfOperation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectHoursOfOperationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectHoursOfOperation), &result)
	return &result
}
