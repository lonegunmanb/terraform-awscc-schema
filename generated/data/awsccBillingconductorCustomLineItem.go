package data

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
        "computed": true,
        "description": "Billing Group ARN",
        "description_kind": "plain",
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
              "type": "string"
            },
            "inclusive_start_billing_period": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "computation_rule": {
        "computed": true,
        "description": "The display settings of the Custom Line Item.",
        "description_kind": "plain",
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
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "line_item_filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "match_option": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "set"
              }
            },
            "percentage": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "child_associated_resources": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "percentage_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
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
      "last_modified_time": {
        "computed": true,
        "description": "Latest modified timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
      }
    },
    "description": "Data Source schema for AWS::BillingConductor::CustomLineItem",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBillingconductorCustomLineItemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorCustomLineItem), &result)
	return &result
}
