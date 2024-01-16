package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallLoggingConfiguration = `{
  "block": {
    "attributes": {
      "firewall_arn": {
        "description": "A resource ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_name": {
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
      "logging_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_destination_configs": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_destination": {
                    "description": "A key-value pair to configure the logDestinations.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "log_destination_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "log_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource type definition for AWS::NetworkFirewall::LoggingConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallLoggingConfiguration), &result)
	return &result
}
