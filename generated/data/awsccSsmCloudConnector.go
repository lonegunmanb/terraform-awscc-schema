package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmCloudConnector = `{
  "block": {
    "attributes": {
      "cloud_connector_arn": {
        "computed": true,
        "description": "The ARN of the cloud connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_connector_id": {
        "computed": true,
        "description": "The unique identifier of the cloud connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "config_connector_arn": {
        "computed": true,
        "description": "The ARN of the AWS Config connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The configuration for the cloud connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure_configuration": {
              "computed": true,
              "description": "Configuration for connecting to Azure.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_display_name": {
                    "computed": true,
                    "description": "The display name of the Azure AD application.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "application_id": {
                    "computed": true,
                    "description": "The Azure AD application ID.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "targets": {
                    "computed": true,
                    "description": "The targets for the cloud connector. If omitted, the entire tenant is targeted.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "subscriptions": {
                          "computed": true,
                          "description": "List of Azure subscriptions.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "display_name": {
                                "computed": true,
                                "description": "The display name of the Azure subscription.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "The Azure subscription ID.",
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
                  "tenant_display_name": {
                    "computed": true,
                    "description": "The display name of the Azure AD tenant.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "The Azure AD tenant ID. Cannot be changed after creation.",
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
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the cloud connector was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the cloud connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the cloud connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The IAM role ARN used by the cloud connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the cloud connector.",
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
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the cloud connector was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::CloudConnector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmCloudConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmCloudConnector), &result)
	return &result
}
