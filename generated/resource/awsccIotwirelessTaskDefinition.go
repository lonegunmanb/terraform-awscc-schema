package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessTaskDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "TaskDefinition arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_create_tasks": {
        "description": "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "The ID of the new wireless gateway task definition",
        "description_kind": "plain",
        "type": "string"
      },
      "lo_ra_wan_update_gateway_task_entry": {
        "computed": true,
        "description": "The list of task definitions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "current_version": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "model": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "package_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "station": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "update_version": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "model": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "package_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "station": {
                    "computed": true,
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
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "The name of the new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the destination.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "task_definition_type": {
        "computed": true,
        "description": "A filter to list only the wireless gateway task definitions that use this task definition type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update": {
        "computed": true,
        "description": "Information about the gateways to update.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lo_ra_wan": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "current_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "model": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "package_version": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "station": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sig_key_crc": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "update_signature": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "update_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "model": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "package_version": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "station": {
                          "computed": true,
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
              "optional": true
            },
            "update_data_role": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_data_source": {
              "computed": true,
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
    "description": "Creates a gateway task definition.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessTaskDefinition), &result)
	return &result
}
