package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmAssociation = `{
  "block": {
    "attributes": {
      "apply_only_at_cron_interval": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "association_id": {
        "computed": true,
        "description": "Unique identifier of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_name": {
        "computed": true,
        "description": "The name of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "automation_target_parameter_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "calendar_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "compliance_severity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "document_version": {
        "computed": true,
        "description": "The version of the SSM document to associate with the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance that the SSM document is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_concurrency": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_errors": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the SSM document.",
        "description_kind": "plain",
        "type": "string"
      },
      "output_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "output_s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_s3_key_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_s3_region": {
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
      "parameters": {
        "computed": true,
        "description": "Parameter values that the SSM document uses at runtime.",
        "description_kind": "plain",
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      },
      "schedule_expression": {
        "computed": true,
        "description": "A Cron or Rate expression that specifies when the association is applied to the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_offset": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "sync_compliance": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "The targets that the SSM document sends commands to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "values": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "wait_for_success_timeout_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::SSM::Association",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmAssociation), &result)
	return &result
}
