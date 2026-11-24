package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbDbInstance = `{
  "block": {
    "attributes": {
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_rotation_restart": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_performance_insights": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
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
      }
    },
    "description": "Data Source schema for AWS::DocDB::DBInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDocdbDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbDbInstance), &result)
	return &result
}
