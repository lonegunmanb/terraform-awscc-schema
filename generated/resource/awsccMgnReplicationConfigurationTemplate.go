package resource

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
        "description": "Whether to associate the default Application Migration Service Security Group with the Replication Configuration Template.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "bandwidth_throttling": {
        "description": "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "create_public_ip": {
        "description": "Whether to create a Public IP for the Recovery Instance by default.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "data_plane_routing": {
        "description": "The data plane routing mechanism that will be used for replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_large_staging_disk_type": {
        "description": "The Staging Disk EBS volume type to be used during replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ebs_encryption": {
        "description": "The type of EBS encryption to be used during replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ebs_encryption_key_arn": {
        "computed": true,
        "description": "The ARN of the EBS encryption key to be used during replication.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_protocol": {
        "computed": true,
        "description": "Which internet protocol to use for replication of data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_configuration_template_id": {
        "computed": true,
        "description": "The Replication Configuration Template ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_server_instance_type": {
        "description": "The instance type to be used for the replication server.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_servers_security_groups_i_ds": {
        "description": "The security group IDs that will be used by the replication server.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "staging_area_subnet_id": {
        "description": "The subnet to be used by the replication staging area.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "staging_area_tags": {
        "description": "A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "store_snapshot_on_local_zone": {
        "computed": true,
        "description": "Whether to store the Snapshot in the Local Zone of the replication staging area subnet.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "use_dedicated_replication_server": {
        "description": "Whether to use a dedicated Replication Server in the replication staging area.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "use_fips_endpoint": {
        "computed": true,
        "description": "Whether to use a FIPS endpoint when initiating the agent connection to the Application Migration Service service.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Account level Replication Configuration Template for AWS Application Migration Service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMgnReplicationConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnReplicationConfigurationTemplate), &result)
	return &result
}
