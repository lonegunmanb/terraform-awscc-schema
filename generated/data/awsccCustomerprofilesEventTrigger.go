package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesEventTrigger = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the event trigger was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the event trigger.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_trigger_conditions": {
        "computed": true,
        "description": "A list of conditions that determine when an event should trigger the destination.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_trigger_dimensions": {
              "computed": true,
              "description": "A list of dimensions to be evaluated for the event.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "object_attributes": {
                    "computed": true,
                    "description": "A list of object attributes to be evaluated.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison_operator": {
                          "computed": true,
                          "description": "The operator used to compare an attribute against a list of values.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "field_name": {
                          "computed": true,
                          "description": "A field defined within an object type.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "source": {
                          "computed": true,
                          "description": "An attribute contained within a source object.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "values": {
                          "computed": true,
                          "description": "A list of attribute values used for comparison.",
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
                "nesting_mode": "list"
              }
            },
            "logical_operator": {
              "computed": true,
              "description": "The operator used to combine multiple dimensions.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "event_trigger_limits": {
        "computed": true,
        "description": "Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_expiration": {
              "computed": true,
              "description": "Specifies that an event will only trigger the destination if it is processed within a certain latency period.",
              "description_kind": "plain",
              "type": "number"
            },
            "periods": {
              "computed": true,
              "description": "A list of time periods during which the limits apply.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_invocations_per_profile": {
                    "computed": true,
                    "description": "The maximum allowed number of destination invocations per profile.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit of time.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "unlimited": {
                    "computed": true,
                    "description": "If set to true, there is no limit on the number of destination invocations per profile. The default is false.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "value": {
                    "computed": true,
                    "description": "The amount of time of the specified unit.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "event_trigger_name": {
        "computed": true,
        "description": "The unique name of the event trigger.",
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
        "description": "The timestamp of when the event trigger was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_type_name": {
        "computed": true,
        "description": "The unique name of the object type.",
        "description_kind": "plain",
        "type": "string"
      },
      "segment_filter": {
        "computed": true,
        "description": "The destination is triggered only for profiles that meet the criteria of a segment definition.",
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
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::EventTrigger",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesEventTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesEventTrigger), &result)
	return &result
}
