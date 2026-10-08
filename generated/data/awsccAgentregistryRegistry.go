package data

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
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "authorizer_type": {
        "computed": true,
        "description": "The type of authorizer that controls how consumers access the registry's search and MCP invoke operations.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_detection_enabled": {
        "computed": true,
        "description": "Specifies whether auto-detection is requested for the registry. Must be specified together with AutoDetectionScope. Setting this to true is necessary but not sufficient for auto-detection to become active; the preconditions of the configured scope must also be met. To turn auto-detection off, explicitly set this to false - removing AutoDetectionEnabled and AutoDetectionScope from the template is a no-op and leaves the existing auto-detection settings unchanged. A registry cannot be deleted while auto-detection is enabled: set this to false and update the stack before deleting the registry.",
        "description_kind": "plain",
        "type": "bool"
      },
      "auto_detection_scope": {
        "computed": true,
        "description": "The source from which resources are detected. ORGANIZATION sources resources from all member accounts of an AWS Organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_detection_status": {
        "computed": true,
        "description": "The current auto-detection status. ACTIVE indicates that the registry is actively being populated with detected resources. INACTIVE indicates that the preconditions required at the configured scope are not currently met.",
        "description_kind": "plain",
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
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allowed_clients": {
                          "computed": true,
                          "description": "The client identifiers accepted during JWT validation.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allowed_scopes": {
                          "computed": true,
                          "description": "The scopes accepted during JWT validation.",
                          "description_kind": "plain",
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
                                            "type": "string"
                                          },
                                          "match_value_string_list": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "inbound_token_claim_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "inbound_token_claim_value_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "discovery_url": {
                          "computed": true,
                          "description": "The OpenID Connect discovery URL used to retrieve the identity provider's metadata and signing keys.",
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
      "encryption_configuration": {
        "computed": true,
        "description": "The server-side encryption configuration for a registry. Specifies a customer managed key used to encrypt the registry's content. When omitted, the registry's content is encrypted with an AWS owned key. You cannot change the encryption configuration after registry creation. Specifying a different KMS key, adding this property to an existing registry, or removing it replaces the registry: CloudFormation creates a new registry with a new Amazon Resource Name (ARN) and then deletes the original, including all registry records it contains. Registry records that are not managed by the stack are not re-created in the new registry, and if any remain in the original registry its deletion fails and it is left behind.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the customer-managed AWS KMS key used to encrypt the registry's content. The key must be a symmetric encryption key in the same AWS account and Region as the registry. Multi-Region keys are not supported.",
              "description_kind": "plain",
              "type": "string"
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
      "name": {
        "computed": true,
        "description": "The name of the registry.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the registry was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AgentRegistry::Registry",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAgentregistryRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAgentregistryRegistry), &result)
	return &result
}
