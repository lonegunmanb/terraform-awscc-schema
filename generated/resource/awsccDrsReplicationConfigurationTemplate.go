package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDrsReplicationConfigurationTemplate = `{
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
        "description": "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_replicate_new_disks": {
        "computed": true,
        "description": "Whether to allow the AWS replication agent to automatically replicate newly added disks.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth_throttling": {
        "description": "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "create_public_ip": {
        "computed": true,
        "description": "Whether to create a Public IP for the Recovery Instance by default.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_plane_routing": {
        "computed": true,
        "description": "The data plane routing mechanism that will be used for replication.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_large_staging_disk_type": {
        "computed": true,
        "description": "The Staging Disk EBS volume type to be used during replication.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Which version of the Internet Protocol to use for replication of data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pit_policy": {
        "description": "The Point in time (PIT) policy to manage snapshots taken during replication.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Whether this rule is enabled or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "interval": {
              "description": "How often, in the chosen units, a snapshot should be taken.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "retention_duration": {
              "description": "The duration to retain a snapshot for, in the chosen units.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "rule_id": {
              "computed": true,
              "description": "The ID of the rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "units": {
              "description": "The units used to measure the interval and retentionDuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
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
        "optional": true,
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
      "tags": {
        "computed": true,
        "description": "A set of tags to be associated with the Replication Configuration Template resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "use_dedicated_replication_server": {
        "computed": true,
        "description": "Whether to use a dedicated Replication Server in the replication staging area.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "A replication configuration template for AWS Elastic Disaster Recovery Service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDrsReplicationConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDrsReplicationConfigurationTemplate), &result)
	return &result
}
