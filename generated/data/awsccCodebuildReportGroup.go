package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodebuildReportGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_reports": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "export_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "export_config_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_destination": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "encryption_disabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "encryption_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "packaging": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CodeBuild::ReportGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodebuildReportGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodebuildReportGroup), &result)
	return &result
}
