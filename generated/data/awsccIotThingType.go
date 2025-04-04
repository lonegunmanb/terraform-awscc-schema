package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotThingType = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deprecate_thing_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
              "description": "Tag key (1-128 chars). No 'aws:' prefix. Allows: [A-Za-z0-9 _.:/=+-]",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value (1-256 chars). No 'aws:' prefix. Allows: [A-Za-z0-9 _.:/=+-]",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "thing_type_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "thing_type_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "thing_type_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mqtt_5_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "propagating_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "connection_attribute": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "thing_attribute": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "user_property_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "searchable_attributes": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "thing_type_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::IoT::ThingType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotThingTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotThingType), &result)
	return &result
}
