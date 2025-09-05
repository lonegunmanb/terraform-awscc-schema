package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockAutomatedReasoningPolicy = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Time this policy was created",
        "description_kind": "plain",
        "type": "string"
      },
      "definition_hash": {
        "computed": true,
        "description": "The hash for this version",
        "description_kind": "plain",
        "type": "string"
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "The rules definition block of an AutomatedReasoningPolicyDefinition.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alternate_expression": {
                    "computed": true,
                    "description": "An alternate expression for this rule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expression": {
                    "computed": true,
                    "description": "The SMT expression for this rule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description": "A unique id within the PolicyDefinition",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "types": {
              "computed": true,
              "description": "The types definition block of an AutomatedReasoningPolicyDefinition.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "A natural language description of this type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "A name for this type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "A list of valid values for this type.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "description": {
                          "computed": true,
                          "description": "A natural language description of the type's value.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value of the type value.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "variables": {
              "computed": true,
              "description": "The variables definition block of an AutomatedReasoningPolicyDefinition.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "A natural language description of this variable.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "A name from this variable.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "A type for this variable.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "version": {
              "computed": true,
              "description": "The policy format version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "policy_id": {
        "computed": true,
        "description": "The id of the policy",
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
              "description": "Tag Key",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "Time this policy was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Version of the policy that was created. This will always be ` + "`" + `DRAFT` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::AutomatedReasoningPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockAutomatedReasoningPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockAutomatedReasoningPolicy), &result)
	return &result
}
