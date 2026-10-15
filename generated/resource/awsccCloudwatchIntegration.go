package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchIntegration = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time at which the integration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "credential": {
        "computed": true,
        "description": "Write-only API-key credential material for the integration. Stored encrypted and never returned on read.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_key_credential": {
              "computed": true,
              "description": "API-key credential material that can be rotated in place on update.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_value": {
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
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_arn": {
        "computed": true,
        "description": "The ARN of the integration. arn:aws:cloudwatch:{region}:{account}:integration/{integrationId}.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_attributes": {
        "computed": true,
        "description": "Type-specific configuration attributes (max 50 entries). EXTERNAL_AGENT requires 'catalogId'. On update, supplied attributes merge into the existing map; omitting a key does not remove it.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "integration_id": {
        "computed": true,
        "description": "The service-generated unique identifier (UUID) of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "description": "The type of integration. Immutable after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The customer-supplied name of the integration. Unique per account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "IAM role ARN the integration assumes. Required for AWS_INTEGRATION (a valid, service-assumable IAM role ARN). Not used by AWS_CONFIG_SLREC or EXTERNAL_AGENT.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Lifecycle status of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the integration.",
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
      "updated_at": {
        "computed": true,
        "description": "The time at which the integration was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::Integration. Manages a CloudWatch (Amazon CloudWatch Observability / Omni) account-scoped integration. V1 supports AWS_CONFIG_SLREC, AWS_INTEGRATION, and EXTERNAL_AGENT.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchIntegration), &result)
	return &result
}
