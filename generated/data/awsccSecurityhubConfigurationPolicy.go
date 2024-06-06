package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubConfigurationPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the configuration policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_policy": {
        "computed": true,
        "description": "An object that defines how Security Hub is configured.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_hub": {
              "computed": true,
              "description": "An object that defines how AWS Security Hub is configured.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled_standard_identifiers": {
                    "computed": true,
                    "description": "A list that defines which security standards are enabled in the configuration policy.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "security_controls_configuration": {
                    "computed": true,
                    "description": "An object that defines which security controls are enabled in an AWS Security Hub configuration policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "disabled_security_control_identifiers": {
                          "computed": true,
                          "description": "A list of security controls that are disabled in the configuration policy",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "enabled_security_control_identifiers": {
                          "computed": true,
                          "description": "A list of security controls that are enabled in the configuration policy.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "security_control_custom_parameters": {
                          "computed": true,
                          "description": "A list of security controls and control parameter values that are included in a configuration policy.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "parameters": {
                                "computed": true,
                                "description": "An object that specifies parameter values for a control in a configuration policy.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "value": {
                                      "computed": true,
                                      "description": "An object that includes the data type of a security control parameter and its current value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "boolean": {
                                            "computed": true,
                                            "description": "A control parameter that is a boolean.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "double": {
                                            "computed": true,
                                            "description": "A control parameter that is a double.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "enum": {
                                            "computed": true,
                                            "description": "A control parameter that is an enum.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "enum_list": {
                                            "computed": true,
                                            "description": "A control parameter that is a list of enums.",
                                            "description_kind": "plain",
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "integer": {
                                            "computed": true,
                                            "description": "A control parameter that is an integer.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "integer_list": {
                                            "computed": true,
                                            "description": "A control parameter that is a list of integers.",
                                            "description_kind": "plain",
                                            "type": [
                                              "list",
                                              "number"
                                            ]
                                          },
                                          "string": {
                                            "computed": true,
                                            "description": "A control parameter that is a string.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "string_list": {
                                            "computed": true,
                                            "description": "A control parameter that is a list of strings.",
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
                                    "value_type": {
                                      "computed": true,
                                      "description": "Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "map"
                                }
                              },
                              "security_control_id": {
                                "computed": true,
                                "description": "The ID of the security control.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "service_enabled": {
                    "computed": true,
                    "description": "Indicates whether Security Hub is enabled in the policy.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "configuration_policy_id": {
        "computed": true,
        "description": "The universally unique identifier (UUID) of the configuration policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time, in UTC and ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the configuration policy.",
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
        "description": "The name of the configuration policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_enabled": {
        "computed": true,
        "description": "Indicates whether the service that the configuration policy applies to is enabled in the policy.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time, in UTC and ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::ConfigurationPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubConfigurationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubConfigurationPolicy), &result)
	return &result
}
