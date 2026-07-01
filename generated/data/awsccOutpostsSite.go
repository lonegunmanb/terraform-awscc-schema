package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOutpostsSite = `{
  "block": {
    "attributes": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "notes": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "operating_address": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "address_line_3": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "contact_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "contact_phone_number": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "district_or_county": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "municipality": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "postal_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "state_or_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "rack_physical_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fiber_optic_cable_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "maximum_supported_weight_lbs": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "optical_standard": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "power_connector": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "power_draw_kva": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "power_feed_drop": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "power_phase": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "uplink_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "uplink_gbps": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "shipping_address": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "address_line_3": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "contact_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "contact_phone_number": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "district_or_county": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "municipality": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "postal_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "state_or_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "site_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "site_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Outposts::Site",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOutpostsSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOutpostsSite), &result)
	return &result
}
