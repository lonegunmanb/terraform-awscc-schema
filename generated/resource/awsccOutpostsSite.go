package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_3": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_phone_number": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "district_or_county": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "municipality": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state_or_region": {
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
      "rack_physical_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fiber_optic_cable_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maximum_supported_weight_lbs": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "optical_standard": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "power_connector": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "power_draw_kva": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "power_feed_drop": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "power_phase": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uplink_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uplink_gbps": {
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
      "shipping_address": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_3": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_phone_number": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "district_or_county": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "municipality": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state_or_region": {
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
    "description": "Definition of AWS::Outposts::Site Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOutpostsSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOutpostsSite), &result)
	return &result
}
