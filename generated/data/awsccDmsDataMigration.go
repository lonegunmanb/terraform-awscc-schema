package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsDataMigration = `{
  "block": {
    "attributes": {
      "data_migration_arn": {
        "computed": true,
        "description": "The property describes an ARN of the data migration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_migration_create_time": {
        "computed": true,
        "description": "The property describes the create time of the data migration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_migration_identifier": {
        "computed": true,
        "description": "The property describes an ARN of the data migration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_migration_name": {
        "computed": true,
        "description": "The property describes a name to identify the data migration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_migration_settings": {
        "computed": true,
        "description": "The property describes the settings for the data migration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_enabled": {
              "computed": true,
              "description": "The property specifies whether to enable the CloudWatch log.",
              "description_kind": "plain",
              "type": "bool"
            },
            "number_of_jobs": {
              "computed": true,
              "description": "The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.",
              "description_kind": "plain",
              "type": "number"
            },
            "selection_rules": {
              "computed": true,
              "description": "The property specifies the rules of selecting objects for data migration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_migration_type": {
        "computed": true,
        "description": "The property describes the type of migration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "migration_project_identifier": {
        "computed": true,
        "description": "The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn",
        "description_kind": "plain",
        "type": "string"
      },
      "service_access_role_arn": {
        "computed": true,
        "description": "The property describes Amazon Resource Name (ARN) of the service access role.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_data_settings": {
        "computed": true,
        "description": "The property describes the settings for the data migration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cdc_start_position": {
              "computed": true,
              "description": "The property is a point in the database engine's log that defines a time where you can begin CDC.",
              "description_kind": "plain",
              "type": "string"
            },
            "cdc_start_time": {
              "computed": true,
              "description": "The property indicates the start time for a change data capture (CDC) operation. The value is server time in UTC format.",
              "description_kind": "plain",
              "type": "string"
            },
            "cdc_stop_time": {
              "computed": true,
              "description": "The property indicates the stop time for a change data capture (CDC) operation. The value is server time in UTC format.",
              "description_kind": "plain",
              "type": "string"
            },
            "slot_name": {
              "computed": true,
              "description": "The property sets the name of a previously created logical replication slot for a change data capture (CDC) load of the source instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
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
    "description": "Data Source schema for AWS::DMS::DataMigration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsDataMigrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsDataMigration), &result)
	return &result
}
