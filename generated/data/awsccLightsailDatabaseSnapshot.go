package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDatabaseSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the database snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the database snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description": "The software of the database snapshot (for example, MySQL).",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The database engine version for the database snapshot (for example, 5.7.23).",
        "description_kind": "plain",
        "type": "string"
      },
      "from_relational_database_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the database from which the database snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_relational_database_blueprint_id": {
        "computed": true,
        "description": "The blueprint ID of the database from which the database snapshot was created. A blueprint describes the major engine version of a database.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_relational_database_bundle_id": {
        "computed": true,
        "description": "The bundle ID of the database from which the database snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_relational_database_name": {
        "computed": true,
        "description": "The name of the source database from which the database snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The Region name and Availability Zone where the database snapshot is located.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone where the database snapshot was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The AWS Region where the database snapshot was created.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the database snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "relational_database_name": {
        "computed": true,
        "description": "The name of the database on which to base your new snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "relational_database_snapshot_name": {
        "computed": true,
        "description": "The name for your new database snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The Lightsail resource type.",
        "description_kind": "plain",
        "type": "string"
      },
      "size_in_gb": {
        "computed": true,
        "description": "The size of the disk in GB (for example, 32) for the database snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The state of the database snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "The support code for the database snapshot. Include this code in your email to support when you have questions about a database snapshot in Lightsail. This code enables our support team to look up your Lightsail information more easily.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Lightsail::DatabaseSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailDatabaseSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDatabaseSnapshot), &result)
	return &result
}
