package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsReplicationInstance = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The amount of storage (in gigabytes) to be initially allocated for the replication instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allow_major_version_upgrade": {
        "computed": true,
        "description": "Indicates that major version upgrades are allowed. Changing this parameter does not result in an outage, and the change is asynchronously applied as soon as possible.\n\nThis parameter must be set to ` + "`" + `true` + "`" + ` when specifying a value for the ` + "`" + `EngineVersion` + "`" + ` parameter that is a different major version than the replication instance's current version.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window. This parameter defaults to ` + "`" + `true` + "`" + `.\n\nDefault: ` + "`" + `true` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone that the replication instance will be created in.\n\nThe default value is a random, system-chosen Availability Zone in the endpoint's AWS Region, for example ` + "`" + `us-east-1d` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name_servers": {
        "computed": true,
        "description": "A list of custom DNS name servers supported for the replication instance to access your on-premise source or target database. This list overrides the default name servers supported by the replication instance. You can specify a comma-separated list of internet addresses for up to four on-premise DNS name servers. For example: ` + "`" + `\"1.1.1.1,2.2.2.2,3.3.3.3,4.4.4.4\"` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The engine version number of the replication instance.\n\nIf an engine version number is not specified when a replication instance is created, the default is the latest engine version available.",
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
      "kms_key_id": {
        "computed": true,
        "description": "An AWS KMS key identifier that is used to encrypt the data on the replication instance.\n\nIf you don't specify a value for the ` + "`" + `KmsKeyId` + "`" + ` parameter, AWS DMS uses your default encryption key.\n\nAWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS Region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description": "Specifies whether the replication instance is a Multi-AZ deployment. You can't set the ` + "`" + `AvailabilityZone` + "`" + ` parameter if the Multi-AZ parameter is set to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_type": {
        "computed": true,
        "description": "The type of IP address protocol used by a replication instance, such as IPv4 only or Dual-stack that supports both IPv4 and IPv6 addressing. IPv6 only is not yet supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range during which system maintenance can occur, in UTC.\n\nFormat: ` + "`" + `ddd:hh24:mi-ddd:hh24:mi` + "`" + `\n\nDefault: A 30-minute window selected at random from an 8-hour block of time per AWS Region, occurring on a random day of the week.\n\nValid days (` + "`" + `ddd` + "`" + `): ` + "`" + `Mon` + "`" + ` | ` + "`" + `Tue` + "`" + ` | ` + "`" + `Wed` + "`" + ` | ` + "`" + `Thu` + "`" + ` | ` + "`" + `Fri` + "`" + ` | ` + "`" + `Sat` + "`" + ` | ` + "`" + `Sun` + "`" + `\n\nConstraints: Minimum 30-minute window.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Specifies the accessibility options for the replication instance. A value of ` + "`" + `true` + "`" + ` represents an instance with a public IP address. A value of ` + "`" + `false` + "`" + ` represents an instance with a private IP address. The default value is ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replication_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the replication instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_instance_class": {
        "description": "The compute and memory capacity of the replication instance as defined for the specified replication instance class. For example, to specify the instance class dms.c4.large, set this parameter to ` + "`" + `\"dms.c4.large\"` + "`" + `. For more information on the settings and capacities for the available replication instance classes, see [Selecting the right AWS DMS replication instance for your migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth) in the AWS Database Migration Service User Guide.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_instance_identifier": {
        "computed": true,
        "description": "The replication instance identifier. This parameter is stored as a lowercase string.\n\nConstraints:\n- Must contain 1-63 alphanumeric characters or hyphens.\n- First character must be a letter.\n- Can't end with a hyphen or contain two consecutive hyphens.\n\nExample: ` + "`" + `myrepinstance` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_instance_private_ip_addresses": {
        "computed": true,
        "description": "One or more private IP addresses for the replication instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_instance_public_ip_addresses": {
        "computed": true,
        "description": "One or more public IP addresses for the replication instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_subnet_group_identifier": {
        "computed": true,
        "description": "A subnet group to associate with the replication instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_identifier": {
        "computed": true,
        "description": "A display name for the resource identifier at the end of the ` + "`" + `EndpointArn` + "`" + ` response parameter that is returned in the created ` + "`" + `Endpoint` + "`" + ` object. The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as ` + "`" + `Example-App-ARN1` + "`" + `. For example, this value might result in the ` + "`" + `EndpointArn` + "`" + ` value ` + "`" + `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` + "`" + `. If you don't specify a ` + "`" + `ResourceIdentifier` + "`" + ` value, AWS DMS generates a default identifier value for the end of ` + "`" + `EndpointArn` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags to be assigned to the replication instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "Specifies the virtual private cloud (VPC) security group to be used with the replication instance. The VPC security group must work with the VPC containing the replication instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::DMS::ReplicationInstance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsReplicationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsReplicationInstance), &result)
	return &result
}
