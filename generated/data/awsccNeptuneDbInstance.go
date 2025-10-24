package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneDbInstance = `{
  "block": {
    "attributes": {
      "allow_major_version_upgrade": {
        "computed": true,
        "description": "Indicates that major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible. This parameter must be set to true when specifying a value for the EngineVersion parameter that is a different major version than the DB instance's current version.",
        "description_kind": "plain",
        "type": "bool"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "Indicates that minor version patches are applied automatically.\n\nWhen updating this property, some interruptions may occur.",
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description": "Specifies the name of the Availability Zone the DB instance is located in.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description": "Contains the name of the compute and memory capacity class of the DB instance.\n\nIf you update this property, some interruptions may occur.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_identifier": {
        "computed": true,
        "description": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_parameter_group_name": {
        "computed": true,
        "description": "The name of an existing DB parameter group or a reference to an AWS::Neptune::DBParameterGroup resource created in the template. If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_snapshot_identifier": {
        "computed": true,
        "description": "This parameter is not supported.\n\n` + "`" + `AWS::Neptune::DBInstance` + "`" + ` does not support restoring from snapshots.\n\n` + "`" + `AWS::Neptune::DBCluster` + "`" + ` does support restoring from snapshots.\n\n",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "A DB subnet group to associate with the DB instance. If you update this value, the new subnet group must be a subnet group in a new virtual private cloud (VPC).",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the database. For example: ` + "`" + `mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com` + "`" + `.",
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
        "description": "The port number on which the database accepts connections. For example: ` + "`" + `8182` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Indicates that public accessibility is enabled. This should be enabled in combination with IAM Auth enabled on the DBCluster",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this DB instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Neptune::DBInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNeptuneDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneDbInstance), &result)
	return &result
}
