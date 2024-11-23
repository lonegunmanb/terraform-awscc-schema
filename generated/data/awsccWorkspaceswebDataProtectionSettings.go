package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebDataProtectionSettings = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "associated_portal_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_protection_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
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
      "inline_redaction_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "global_confidence_level": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "global_enforced_urls": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "global_exempt_urls": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "inline_redaction_patterns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "built_in_pattern_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "confidence_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "custom_pattern": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "keyword_regex": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "pattern_description": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "pattern_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "pattern_regex": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "enforced_urls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exempt_urls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "redaction_place_holder": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "redaction_place_holder_text": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "redaction_place_holder_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
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
    "description": "Data Source schema for AWS::WorkSpacesWeb::DataProtectionSettings",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspaceswebDataProtectionSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebDataProtectionSettings), &result)
	return &result
}
