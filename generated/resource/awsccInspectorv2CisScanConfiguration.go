package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorv2CisScanConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "CIS Scan configuration unique identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_name": {
        "description": "Name of the scan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description": "Choose a Schedule cadence",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "daily": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "start_time": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_zone": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "monthly": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_zone": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "one_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weekly": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_zone": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
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
      "security_level": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "targets": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "target_resource_tags": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                [
                  "list",
                  "string"
                ]
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "CIS Scan Configuration resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccInspectorv2CisScanConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2CisScanConfiguration), &result)
	return &result
}
