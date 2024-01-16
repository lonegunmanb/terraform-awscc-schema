package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessWirelessDevice = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Wireless device arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Wireless device description",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_name": {
        "computed": true,
        "description": "Wireless device destination name",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_uplink_received_at": {
        "computed": true,
        "description": "The date and time when the most recent uplink was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "lo_ra_wan": {
        "computed": true,
        "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "abp_v10_x": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dev_addr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_keys": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_s_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "nwk_s_key": {
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
            "abp_v11": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dev_addr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_keys": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_s_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "f_nwk_s_int_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "nwk_s_enc_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s_nwk_s_int_key": {
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
            "dev_eui": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "device_profile_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "otaa_v10_x": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_eui": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "app_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "otaa_v11": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "join_eui": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "nwk_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_profile_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "Wireless device name",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.",
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
      },
      "thing_arn": {
        "computed": true,
        "description": "Thing arn. Passed into update to associate Thing with Wireless device.",
        "description_kind": "plain",
        "type": "string"
      },
      "thing_name": {
        "computed": true,
        "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Wireless device type, currently only Sidewalk and LoRa",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTWireless::WirelessDevice",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessWirelessDeviceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessWirelessDevice), &result)
	return &result
}
