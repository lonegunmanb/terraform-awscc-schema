package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
      },
      "tls_inspection_configuration": {
        "computed": true,
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
                          "type": "string"
                        },
                        "unknown_status_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                                "type": "number"
                              },
                              "to_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "destinations": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "address_definition": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "protocols": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "number"
                              },
                              "to_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "sources": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "address_definition": {
                                "computed": true,
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
                  "server_certificates": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_arn": {
                          "computed": true,
                          "description": "A resource ARN.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::TLSInspectionConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallTlsInspectionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallTlsInspectionConfiguration), &result)
	return &result
}
