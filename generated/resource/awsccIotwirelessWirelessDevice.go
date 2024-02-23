package resource

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
        "optional": true,
        "type": "string"
      },
      "destination_name": {
        "description": "Wireless device destination name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Wireless device Id. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_uplink_received_at": {
        "computed": true,
        "description": "The date and time when the most recent uplink was received.",
        "description_kind": "plain",
        "optional": true,
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
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "session_keys": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_s_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "nwk_s_key": {
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
            "abp_v11": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dev_addr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "session_keys": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_s_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "f_nwk_s_int_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "nwk_s_enc_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s_nwk_s_int_key": {
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
            "dev_eui": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_profile_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "f_ports": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "applications": {
                    "computed": true,
                    "description": "A list of optional LoRaWAN application information, which can be used for geolocation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination_name": {
                          "computed": true,
                          "description": "The name of the position data destination that describes the AWS IoT rule that processes the device's position data for use by AWS IoT Core for LoRaWAN.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "f_port": {
                          "computed": true,
                          "description": "The Fport value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "type": {
                          "computed": true,
                          "description": "Application type, which can be specified to obtain real-time position information of your LoRaWAN device.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "otaa_v10_x": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_eui": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "app_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "otaa_v11": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "join_eui": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "nwk_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_profile_id": {
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
      "name": {
        "computed": true,
        "description": "Wireless device name",
        "description_kind": "plain",
        "optional": true,
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
      "thing_arn": {
        "computed": true,
        "description": "Thing arn. Passed into update to associate Thing with Wireless device.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "thing_name": {
        "computed": true,
        "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "Wireless device type, currently only Sidewalk and LoRa",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Create and manage wireless gateways, including LoRa gateways.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessWirelessDeviceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessWirelessDevice), &result)
	return &result
}
