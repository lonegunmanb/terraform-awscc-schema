package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseGateway = `{
  "block": {
    "attributes": {
      "gateway_capability_summaries": {
        "computed": true,
        "description": "A list of gateway capability summaries that each contain a namespace and status.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capability_configuration": {
              "computed": true,
              "description": "The JSON document that defines the gateway capability's configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "capability_namespace": {
              "computed": true,
              "description": "The namespace of the capability configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "gateway_id": {
        "computed": true,
        "description": "The ID of the gateway device.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_name": {
        "description": "A unique, friendly name for the gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_platform": {
        "description": "The gateway's platform. You can only specify one platform in a gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "greengrass_v2": {
              "computed": true,
              "description": "A gateway that runs on AWS IoT Greengrass V2.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "core_device_operating_system": {
                    "computed": true,
                    "description": "The operating system of the core device in AWS IoT Greengrass V2.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "core_device_thing_name": {
                    "computed": true,
                    "description": "The name of the CoreDevice in GreenGrass V2.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "siemens_ie": {
              "computed": true,
              "description": "A gateway that runs on Siemens Industrial Edge.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iot_core_thing_name": {
                    "computed": true,
                    "description": "The name of the IoT Core Thing.",
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
        "required": true
      },
      "gateway_version": {
        "computed": true,
        "description": "The version of the gateway you want to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the gateway.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::IoTSiteWise::Gateway",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseGateway), &result)
	return &result
}
