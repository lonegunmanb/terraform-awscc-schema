package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorv2Connector = `{
  "block": {
    "attributes": {
      "connector_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the connector was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Optional description of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "enablement_status": {
        "computed": true,
        "description": "The enablement status of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "enablement_status_reason": {
        "computed": true,
        "description": "Reason for the current enablement status, if applicable.",
        "description_kind": "plain",
        "type": "string"
      },
      "health": {
        "computed": true,
        "description": "Health status of the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connector_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "last_checked_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
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
      "last_updated_at": {
        "computed": true,
        "description": "Timestamp when the connector was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Display name for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_configuration": {
        "computed": true,
        "description": "Provider-specific configuration including regions and scope.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auto_install_vm_scanner": {
                    "computed": true,
                    "description": "Whether to automatically install the VM scanner. Defaults to true.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "aws_config_connector_arn": {
                    "computed": true,
                    "description": "The ARN of the AWS Config connector used for Azure resource discovery.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "azure_regions": {
                    "computed": true,
                    "description": "List of Azure regions to scan.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "scope_configuration": {
                    "computed": true,
                    "description": "Defines which resource types to scan and at what scope level.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_image_scanning": {
                          "computed": true,
                          "description": "Defines the scope of Azure resources to monitor for a specific resource type.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "scope_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "scope_values": {
                                "computed": true,
                                "description": "List of subscription IDs. Empty for TENANT scope.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "state_reason": {
                                "computed": true,
                                "description": "Reason for the current scope state.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "serverless_scanning": {
                          "computed": true,
                          "description": "Defines the scope of Azure resources to monitor for a specific resource type.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "scope_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "scope_values": {
                                "computed": true,
                                "description": "List of subscription IDs. Empty for TENANT scope.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "state_reason": {
                                "computed": true,
                                "description": "Reason for the current scope state.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "vm_scanning": {
                          "computed": true,
                          "description": "Defines the scope of Azure resources to monitor for a specific resource type.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "scope_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "scope_values": {
                                "computed": true,
                                "description": "List of subscription IDs. Empty for TENANT scope.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "state_reason": {
                                "computed": true,
                                "description": "Reason for the current scope state.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "provider_name": {
        "computed": true,
        "description": "The cloud provider for this connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the connector.",
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
      }
    },
    "description": "Data Source schema for AWS::InspectorV2::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorv2ConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2Connector), &result)
	return &result
}
