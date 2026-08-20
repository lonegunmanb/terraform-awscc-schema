package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightDlpSetting = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dlp_setting_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "provider_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "microsoft_purview": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "label_action_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "unmapped_action": {
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
      "provider_outage_action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::DLPSetting",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightDlpSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightDlpSetting), &result)
	return &result
}
