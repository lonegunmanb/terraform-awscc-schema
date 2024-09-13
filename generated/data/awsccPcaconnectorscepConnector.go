package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectorscepConnector = `{
  "block": {
    "attributes": {
      "certificate_authority_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
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
      "mobile_device_management": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "intune": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "azure_application_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "domain": {
                    "computed": true,
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
      "open_id_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audience": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "issuer": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subject": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorSCEP::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectorscepConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectorscepConnector), &result)
	return &result
}
