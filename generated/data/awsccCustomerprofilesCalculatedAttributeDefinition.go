package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesCalculatedAttributeDefinition = `{
  "block": {
    "attributes": {
      "attribute_details": {
        "computed": true,
        "description": "Mathematical expression and a list of attribute items specified in that expression.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "A list of attribute items specified in the mathematical expression.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of an attribute defined in a profile object type.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "expression": {
              "computed": true,
              "description": "Mathematical expression that is performed on attribute items provided in the attribute list. Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "calculated_attribute_name": {
        "computed": true,
        "description": "The unique name of the calculated attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "conditions": {
        "computed": true,
        "description": "The conditions including range, object count, and threshold for the calculated attribute.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "object_count": {
              "computed": true,
              "description": "The number of profile objects used for the calculated attribute.",
              "description_kind": "plain",
              "type": "number"
            },
            "range": {
              "computed": true,
              "description": "The relative time period over which data is included in the aggregation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "timestamp_format": {
                    "computed": true,
                    "description": "The format the timestamp field in your JSON object is specified. This value should be one of EPOCHMILLI or ISO_8601. E.g. if your object type is MyType and source JSON is {\"generatedAt\": {\"timestamp\": \"2001-07-04T12:08:56.235Z\"}}, then TimestampFormat should be \"ISO_8601\".",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timestamp_source": {
                    "computed": true,
                    "description": "An expression specifying the field in your JSON object from which the date should be parsed. The expression should follow the structure of \\\"{ObjectTypeName.\u003cLocation of timestamp field in JSON pointer format\u003e}\\\". E.g. if your object type is MyType and source JSON is {\"generatedAt\": {\"timestamp\": \"1737587945945\"}}, then TimestampSource should be \"{MyType.generatedAt.timestamp}\".",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit of time.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The amount of time of the specified unit.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "value_range": {
                    "computed": true,
                    "description": "A structure specifying the endpoints of the relative time period over which data is included in the aggregation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end": {
                          "computed": true,
                          "description": "The ending point for this range. Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start": {
                          "computed": true,
                          "description": "The starting point for this range. Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.",
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
            "threshold": {
              "computed": true,
              "description": "The threshold for the calculated attribute.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "operator": {
                    "computed": true,
                    "description": "The operator of the threshold.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value of the threshold.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the calculated attribute definition was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the calculated attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the calculated attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the calculated attribute definition was most recently edited.",
        "description_kind": "plain",
        "type": "string"
      },
      "readiness": {
        "computed": true,
        "description": "The readiness status of the calculated attribute.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "message": {
              "computed": true,
              "description": "Any information pertaining to the status of the calculated attribute if required.",
              "description_kind": "plain",
              "type": "string"
            },
            "progress_percentage": {
              "computed": true,
              "description": "The progress percentage for including historical data in your calculated attribute.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "statistic": {
        "computed": true,
        "description": "The aggregation operation to perform for the calculated attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the calculated attribute definition.",
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
          "nesting_mode": "set"
        }
      },
      "use_historical_data": {
        "computed": true,
        "description": "Whether to use historical data for the calculated attribute.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::CalculatedAttributeDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesCalculatedAttributeDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesCalculatedAttributeDefinition), &result)
	return &result
}
