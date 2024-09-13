package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectorscepConnector = `{
  "block": {
    "attributes": {
      "certificate_authority_arn": {
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "domain": {
                    "description_kind": "plain",
                    "required": true,
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
        "optional": true,
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
    "description": "Represents a Connector that allows certificate issuance through Simple Certificate Enrollment Protocol (SCEP)",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectorscepConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectorscepConnector), &result)
	return &result
}
