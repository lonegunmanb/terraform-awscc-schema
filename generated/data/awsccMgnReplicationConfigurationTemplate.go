package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMgnReplicationConfigurationTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Replication Configuration Template ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "associate_default_security_group": {
        "computed": true,
        "description": "Whether to associate the default Application Migration Service Security Group with the Replication Configuration Template.",
        "description_kind": "plain",
        "type": "bool"
      },
      "bandwidth_throttling": {
        "computed": true,
        "description": "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.",
        "description_kind": "plain",
        "type": "number"
      },
      "create_public_ip": {
        "computed": true,
        "description": "Whether to create a Public IP for the Recovery Instance by default.",
        "description_kind": "plain",
        "type": "bool"
      },
      "data_plane_routing": {
        "computed": true,
        "description": "The data plane routing mechanism that will be used for replication.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_large_staging_disk_type": {
        "computed": true,
        "description": "The Staging Disk EBS volume type to be used during replication.",
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_encryption": {
        "computed": true,
        "description": "The type of EBS encryption to be used during replication.",
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_encryption_key_arn": {
        "computed": true,
        "description": "The ARN of the EBS encryption key to be used during replication.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "internet_protocol": {
        "computed": true,
        "description": "Which internet protocol to use for replication of data.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration_template_id": {
        "computed": true,
        "description": "The Replication Configuration Template ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_server_instance_type": {
        "computed": true,
        "description": "The instance type to be used for the replication server.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_servers_security_groups_i_ds": {
        "computed": true,
        "description": "The security group IDs that will be used by the replication server.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "staging_area_subnet_id": {
        "computed": true,
        "description": "The subnet to be used by the replication staging area.",
        "description_kind": "plain",
        "type": "string"
      },
      "staging_area_tags": {
        "computed": true,
        "description": "A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "store_snapshot_on_local_zone": {
        "computed": true,
        "description": "Whether to store the Snapshot in the Local Zone of the replication staging area subnet.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags to be associated with the Replication Configuration Template resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "use_dedicated_replication_server": {
        "computed": true,
        "description": "Whether to use a dedicated Replication Server in the replication staging area.",
        "description_kind": "plain",
        "type": "bool"
      },
      "use_fips_endpoint": {
        "computed": true,
        "description": "Whether to use a FIPS endpoint when initiating the agent connection to the Application Migration Service service.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::MGN::ReplicationConfigurationTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMgnReplicationConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnReplicationConfigurationTemplate), &result)
	return &result
}
