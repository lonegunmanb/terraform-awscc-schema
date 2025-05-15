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
        "description": "The scanning rules associated with the registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_filters": {
              "computed": true,
              "description": "The details of a scanning repository filter. For more information on how to use filters, see [Using filters](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html#image-scanning-filters) in the *Amazon Elastic Container Registry User Guide*.",
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
              "description": "The frequency that scans are performed at for a private registry. When the ` + "`" + `` + "`" + `ENHANCED` + "`" + `` + "`" + ` scan type is specified, the supported scan frequencies are ` + "`" + `` + "`" + `CONTINUOUS_SCAN` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SCAN_ON_PUSH` + "`" + `` + "`" + `. When the ` + "`" + `` + "`" + `BASIC` + "`" + `` + "`" + ` scan type is specified, the ` + "`" + `` + "`" + `SCAN_ON_PUSH` + "`" + `` + "`" + ` scan frequency is supported. If scan on push is not specified, then the ` + "`" + `` + "`" + `MANUAL` + "`" + `` + "`" + ` scan frequency is set by default.",
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
