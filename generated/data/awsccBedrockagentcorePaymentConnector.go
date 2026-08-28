package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcorePaymentConnector = `{
  "block": {
    "attributes": {
      "authorization_url": {
        "computed": true,
        "description": "The URL the user must open to complete OAuth consent. Only present when ConnectorStatus is PENDING_AUTHENTICATION.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_created_at": {
        "computed": true,
        "description": "The timestamp when the connector was created",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_last_updated_at": {
        "computed": true,
        "description": "The timestamp when the connector was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_name": {
        "computed": true,
        "description": "The name of the payment connector",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_configurations": {
        "computed": true,
        "description": "The credential provider configurations for the connector. Required when ProvisionMode is MANUAL or not specified. Empty for QUICK_CREATE until provisioning completes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "coinbase_cdp": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credential_provider_arn": {
                    "computed": true,
                    "description": "The ARN of the payment credential provider",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stripe_privy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credential_provider_arn": {
                    "computed": true,
                    "description": "The ARN of the payment credential provider",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the payment connector",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "payment_connector_arn": {
        "computed": true,
        "description": "Synthetic ARN for the payment connector (used for engine resolution)",
        "description_kind": "plain",
        "type": "string"
      },
      "payment_connector_id": {
        "computed": true,
        "description": "The unique identifier for the payment connector",
        "description_kind": "plain",
        "type": "string"
      },
      "payment_manager_id": {
        "computed": true,
        "description": "The identifier of the parent payment manager",
        "description_kind": "plain",
        "type": "string"
      },
      "provision_mode": {
        "computed": true,
        "description": "The provision mode for creating the connector. MANUAL requires CredentialProviderConfigurations; QUICK_CREATE orchestrates OAuth consent and credential provisioning.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::PaymentConnector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcorePaymentConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcorePaymentConnector), &result)
	return &result
}
