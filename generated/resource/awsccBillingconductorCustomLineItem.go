package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBillingconductorCustomLineItem = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The account which this custom line item will be charged to",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "association_size": {
        "computed": true,
        "description": "Number of source values associated to this custom line item",
        "description_kind": "plain",
        "type": "number"
      },
      "billing_group_arn": {
        "description": "Billing Group ARN",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "billing_period_range": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclusive_end_billing_period": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inclusive_start_billing_period": {
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
      "computation_rule": {
        "computed": true,
        "description": "The display settings of the Custom Line Item.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "currency_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_line_item_charge_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "flat": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "charge_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "line_item_filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "attribute_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "match_option": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "percentage": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "child_associated_resources": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "percentage_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
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
      "last_modified_time": {
        "computed": true,
        "description": "Latest modified timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "presentation_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "service": {
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
      "product_code": {
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A custom line item is an one time charge that is applied to a specific billing group's bill.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBillingconductorCustomLineItemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorCustomLineItem), &result)
	return &result
}
