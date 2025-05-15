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
        "description": "The scanning rules associated with the registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_filters": {
              "description": "The details of a scanning repository filter. For more information on how to use filters, see [Using filters](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html#image-scanning-filters) in the *Amazon Elastic Container Registry User Guide*.",
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
              "description": "The frequency that scans are performed at for a private registry. When the ` + "`" + `` + "`" + `ENHANCED` + "`" + `` + "`" + ` scan type is specified, the supported scan frequencies are ` + "`" + `` + "`" + `CONTINUOUS_SCAN` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SCAN_ON_PUSH` + "`" + `` + "`" + `. When the ` + "`" + `` + "`" + `BASIC` + "`" + `` + "`" + ` scan type is specified, the ` + "`" + `` + "`" + `SCAN_ON_PUSH` + "`" + `` + "`" + ` scan frequency is supported. If scan on push is not specified, then the ` + "`" + `` + "`" + `MANUAL` + "`" + `` + "`" + ` scan frequency is set by default.",
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
    "description": "The scanning configuration for a private registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrRegistryScanningConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRegistryScanningConfiguration), &result)
	return &result
}
