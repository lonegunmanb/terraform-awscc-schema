package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallTlsInspectionConfiguration = `{
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
      },
      "tls_inspection_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_certificate_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_authority_arn": {
                    "computed": true,
                    "description": "A resource ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "check_certificate_revocation_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "revoked_status_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "unknown_status_action": {
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
                  "scopes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination_ports": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "from_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "to_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "destinations": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "address_definition": {
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
                        "protocols": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "source_ports": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "from_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "to_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "sources": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "address_definition": {
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "server_certificates": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_arn": {
                          "computed": true,
                          "description": "A resource ARN.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tls_inspection_configuration_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "tls_inspection_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tls_inspection_configuration_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallTlsInspectionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallTlsInspectionConfiguration), &result)
	return &result
}
