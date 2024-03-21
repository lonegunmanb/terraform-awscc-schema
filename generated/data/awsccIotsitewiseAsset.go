package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseAsset = `{
  "block": {
    "attributes": {
      "asset_arn": {
        "computed": true,
        "description": "The ARN of the asset",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_description": {
        "computed": true,
        "description": "A description for the asset",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_external_id": {
        "computed": true,
        "description": "The External ID of the asset",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_hierarchies": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "child_asset_id": {
              "computed": true,
              "description": "The ID of the child asset to be associated.",
              "description_kind": "plain",
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description": "String-friendly customer provided external ID",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "Customer provided actual UUID for property",
              "description_kind": "plain",
              "type": "string"
            },
            "logical_id": {
              "computed": true,
              "description": "The LogicalID of a hierarchy in the parent asset's model.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "asset_id": {
        "computed": true,
        "description": "The ID of the asset",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_model_id": {
        "computed": true,
        "description": "The ID of the asset model from which to create the asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_name": {
        "computed": true,
        "description": "A unique, friendly name for the asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alias": {
              "computed": true,
              "description": "The property alias that identifies the property.",
              "description_kind": "plain",
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description": "String-friendly customer provided external ID",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "Customer provided actual UUID for property",
              "description_kind": "plain",
              "type": "string"
            },
            "logical_id": {
              "computed": true,
              "description": "Customer provided ID for property.",
              "description_kind": "plain",
              "type": "string"
            },
            "notification_state": {
              "computed": true,
              "description": "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
              "description_kind": "plain",
              "type": "string"
            },
            "unit": {
              "computed": true,
              "description": "The unit of measure (such as Newtons or RPM) of the asset property. If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTSiteWise::Asset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotsitewiseAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseAsset), &result)
	return &result
}
