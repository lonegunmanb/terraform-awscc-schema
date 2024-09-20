package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsPolicy = `{
  "block": {
    "attributes": {
      "definition": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "statement": {
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
            "template_linked": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_template_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "entity_type": {
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
                  "resource": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "entity_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "entity_type": {
                          "computed": true,
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
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::VerifiedPermissions::Policy Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVerifiedpermissionsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicy), &result)
	return &result
}
