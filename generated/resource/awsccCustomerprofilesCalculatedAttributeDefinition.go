package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesCalculatedAttributeDefinition = `{
  "block": {
    "attributes": {
      "attribute_details": {
        "description": "Mathematical expression and a list of attribute items specified in that expression.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "description": "A list of attribute items specified in the mathematical expression.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "description": "The name of an attribute defined in a profile object type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "required": true
            },
            "expression": {
              "description": "Mathematical expression that is performed on attribute items provided in the attribute list. Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "calculated_attribute_name": {
        "description": "The unique name of the calculated attribute.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "number"
            },
            "range": {
              "computed": true,
              "description": "The relative time period over which data is included in the aggregation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "unit": {
                    "description": "The unit of time.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The amount of time of the specified unit.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "threshold": {
              "computed": true,
              "description": "The threshold for the calculated attribute.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "operator": {
                    "description": "The operator of the threshold.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value of the threshold.",
                    "description_kind": "plain",
                    "required": true,
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
        "optional": true
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
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the calculated attribute.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "The unique name of the domain.",
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
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the calculated attribute definition was most recently edited.",
        "description_kind": "plain",
        "type": "string"
      },
      "statistic": {
        "description": "The aggregation operation to perform for the calculated attribute.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A calculated attribute definition for Customer Profiles",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCustomerprofilesCalculatedAttributeDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesCalculatedAttributeDefinition), &result)
	return &result
}
