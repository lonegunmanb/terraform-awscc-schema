package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRegistryScanningConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "The registry id.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "The scanning rules associated with the registry. A registry scanning configuration may contain a maximum of 2 rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_filters": {
              "computed": true,
              "description": "The repository filters associated with the scanning configuration for a private registry.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "filter": {
                    "computed": true,
                    "description": "The filter to use when scanning.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter_type": {
                    "computed": true,
                    "description": "The type associated with the filter.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "scan_frequency": {
              "computed": true,
              "description": "The frequency that scans are performed.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "scan_type": {
        "computed": true,
        "description": "The type of scanning configured for the registry.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECR::RegistryScanningConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrRegistryScanningConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRegistryScanningConfiguration), &result)
	return &result
}
