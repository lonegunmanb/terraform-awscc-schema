package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodebuildFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "base_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "compute_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disk": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "machine_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "memory": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "v_cpu": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "compute_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_proxy_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_behavior": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ordered_proxy_rules": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "effect": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "entities": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "type": {
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
      "fleet_service_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "overflow_behavior": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "scaling_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_tracking_scaling_configs": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_value": {
                    "computed": true,
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeBuild::Fleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodebuildFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodebuildFleet), &result)
	return &result
}
