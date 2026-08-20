package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAgentregistryRegistry = `{
  "block": {
    "attributes": {
      "approval_configuration": {
        "computed": true,
        "description": "Configuration for the registry's record approval workflow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_approval_rules": {
              "computed": true,
              "description": "The rules that determine which registry records are automatically approved on submission. When omitted or empty, submitted records require manual review.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "authorizer_type": {
        "computed": true,
        "description": "The type of authorizer that controls how consumers access the registry's search and MCP invoke operations.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the registry was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the registry.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discovery_configuration": {
        "computed": true,
        "description": "Discovery configuration for the registry. Controls how consumers are authorized to search the registry and invoke its MCP endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorizer_configuration": {
              "computed": true,
              "description": "The authorizer configuration for the registry. This is a union - specify exactly one member.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_jwt_authorizer": {
                    "computed": true,
                    "description": "Configuration for a custom JWT authorizer that validates inbound bearer tokens against an OpenID Connect identity provider.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "allowed_audience": {
                          "computed": true,
                          "description": "The audience values accepted during JWT validation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allowed_clients": {
                          "computed": true,
                          "description": "The client identifiers accepted during JWT validation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allowed_scopes": {
                          "computed": true,
                          "description": "The scopes accepted during JWT validation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "custom_claims": {
                          "computed": true,
                          "description": "Additional custom claim validations applied to the inbound JWT.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authorizing_claim_match_value": {
                                "computed": true,
                                "description": "The value and match operator used to authorize a claim during JWT validation.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "claim_match_operator": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "claim_match_value": {
                                      "computed": true,
                                      "description": "The expected value used to match a claim. Exactly one member is set.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "match_value_string": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "match_value_string_list": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
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
                              },
                              "inbound_token_claim_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "inbound_token_claim_value_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "discovery_url": {
                          "computed": true,
                          "description": "The OpenID Connect discovery URL used to retrieve the identity provider's metadata and signing keys.",
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
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the registry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "The unique identifier of the registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the registry was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::AgentRegistry::Registry Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAgentregistryRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAgentregistryRegistry), &result)
	return &result
}
