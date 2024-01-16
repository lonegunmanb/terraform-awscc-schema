package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDatabase = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_retention": {
        "computed": true,
        "description": "When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description": "Indicates the certificate that needs to be associated with the database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_arn": {
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
      "master_database_name": {
        "description": "The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "master_user_password": {
        "computed": true,
        "description": "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\". It cannot contain spaces.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_username": {
        "description": "The name for the master user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "The daily time range during which automated backups are created for your new database if automated backups are enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range during which system maintenance can occur on your new database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "relational_database_blueprint_id": {
        "description": "The blueprint ID for your new database. A blueprint describes the major engine version of a database.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relational_database_bundle_id": {
        "description": "The bundle ID for your new database. A bundle describes the performance specifications for your database.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relational_database_name": {
        "description": "The name to use for your new Lightsail database resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relational_database_parameters": {
        "computed": true,
        "description": "Update one or more parameters of the relational database.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_values": {
              "computed": true,
              "description": "Specifies the valid range of values for the parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "apply_method": {
              "computed": true,
              "description": "Indicates when parameter updates are applied. Can be immediate or pending-reboot.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "apply_type": {
              "computed": true,
              "description": "Specifies the engine-specific parameter type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_type": {
              "computed": true,
              "description": "Specifies the valid data type for the parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "Provides a description of the parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_modifiable": {
              "computed": true,
              "description": "A Boolean value indicating whether the parameter can be modified.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "parameter_name": {
              "computed": true,
              "description": "Specifies the name of the parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_value": {
              "computed": true,
              "description": "Specifies the value of the parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "rotate_master_user_password": {
        "computed": true,
        "description": "When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::Database",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDatabase), &result)
	return &result
}
