package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsPolicy = `{
  "block": {
    "attributes": {
      "definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "static": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "statement": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "template_linked": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_template_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "principal": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "entity_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "entity_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "resource": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "entity_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "entity_type": {
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::VerifiedPermissions::Policy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVerifiedpermissionsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicy), &result)
	return &result
}
