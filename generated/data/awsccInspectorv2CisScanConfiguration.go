package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scan_name": {
        "computed": true,
        "description": "Name of the scan",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
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
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "time_zone": {
                          "computed": true,
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
            "monthly": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "time_zone": {
                          "computed": true,
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
            "one_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "weekly": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "days": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "start_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "time_of_day": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "time_zone": {
                          "computed": true,
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "security_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "targets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "target_resource_tags": {
              "computed": true,
              "description_kind": "plain",
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
        }
      }
    },
    "description": "Data Source schema for AWS::InspectorV2::CisScanConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorv2CisScanConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2CisScanConfiguration), &result)
	return &result
}
