package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotfleetwiseSignalCatalog = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modification_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_counts": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "total_actuators": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "total_attributes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "total_branches": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "total_nodes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "total_sensors": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "nodes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actuator": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "assigned_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fully_qualified_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "attribute": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "assigned_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "default_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fully_qualified_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "branch": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fully_qualified_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sensor": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fully_qualified_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "tags": {
        "computed": true,
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
    "description": "Data Source schema for AWS::IoTFleetWise::SignalCatalog",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotfleetwiseSignalCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotfleetwiseSignalCatalog), &result)
	return &result
}
