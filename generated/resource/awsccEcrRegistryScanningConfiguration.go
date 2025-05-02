package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRegistryScanningConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "The registry id.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "description": "The scanning rules associated with the registry. A registry scanning configuration may contain a maximum of 2 rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_filters": {
              "description": "The repository filters associated with the scanning configuration for a private registry.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "filter": {
                    "description": "The filter to use when scanning.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter_type": {
                    "description": "The type associated with the filter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "scan_frequency": {
              "description": "The frequency that scans are performed.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "scan_type": {
        "description": "The type of scanning configured for the registry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::ECR::RegistryScanningConfiguration controls the scanning configuration for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrRegistryScanningConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRegistryScanningConfiguration), &result)
	return &result
}
