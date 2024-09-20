package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppflowConnector = `{
  "block": {
    "attributes": {
      "connector_arn": {
        "computed": true,
        "description": " The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_label": {
        "computed": true,
        "description": " The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connector_provisioning_config": {
        "description": "Contains information about the configuration of the connector being registered.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lambda": {
              "computed": true,
              "description": "Contains information about the configuration of the lambda which is being registered as the connector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description": "Lambda ARN of the connector being registered.",
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
        "required": true
      },
      "connector_provisioning_type": {
        "description": "The provisioning type of the connector. Currently the only supported value is LAMBDA. ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description about the connector that's being registered.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::AppFlow::Connector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppflowConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppflowConnector), &result)
	return &result
}
