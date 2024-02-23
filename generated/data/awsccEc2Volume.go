package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Volume = `{
  "block": {
    "attributes": {
      "auto_enable_io": {
        "computed": true,
        "description": "Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.",
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description": "The ID of the Availability Zone in which to create the volume. For example, ` + "`" + `` + "`" + `us-east-1a` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "Indicates whether the volume should be encrypted. The effect of setting the encryption state to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default) in the *Amazon Elastic Compute Cloud User Guide*.\n Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "iops": {
        "computed": true,
        "description": "The number of I/O operations per second (IOPS). For ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes, this represents the number of IOPS that are provisioned for the volume. For ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.\n The following are the supported values for each volume type:\n  +   ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 3,000 - 16,000 IOPS\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `: 100 - 64,000 IOPS\n  +   ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 100 - 256,000 IOPS\n  \n For ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.\n This parameter is required for ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes. The default for ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volumes is 3,000 IOPS. This parameter is not supported for ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` volumes.",
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The identifier of the kms-key-long to use for Amazon EBS encryption. If ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` is specified, the encrypted state must be ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.\n Alternatively, if you want to specify a different key, you can specify one of the following:\n  +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Key alias. Specify the alias for the key, prefixed with ` + "`" + `` + "`" + `alias/` + "`" + `` + "`" + `. For example, for a key with the alias ` + "`" + `` + "`" + `my_cmk` + "`" + `` + "`" + `, use ` + "`" + `` + "`" + `alias/my_cmk` + "`" + `` + "`" + `. Or to specify the aws-managed-key, use ` + "`" + `` + "`" + `alias/aws/ebs` + "`" + `` + "`" + `.\n  +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_attach_enabled": {
        "computed": true,
        "description": "Indicates whether Amazon EBS Multi-Attach is enabled.\n CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.",
        "description_kind": "plain",
        "type": "bool"
      },
      "outpost_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Outpost.",
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.\n The following are the supported volumes sizes for each volume type:\n  +   ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 1 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `: 4 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 4 - 65,536 GiB\n  +   ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + `: 125 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + `: 1 - 1024 GiB",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_id": {
        "computed": true,
        "description": "The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the volume during creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "throughput": {
        "computed": true,
        "description": "The throughput to provision for a volume, with a maximum of 1,000 MiB/s.\n This parameter is valid only for ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volumes. The default value is 125.\n Valid Range: Minimum value of 125. Maximum value of 1000.",
        "description_kind": "plain",
        "type": "number"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description": "The volume type. This parameter can be one of the following values:\n  +  General Purpose SSD: ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` \n  +  Provisioned IOPS SSD: ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` \n  +  Throughput Optimized HDD: ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + ` \n  +  Cold HDD: ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + ` \n  +  Magnetic: ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` \n  \n For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon Elastic Compute Cloud User Guide*.\n Default: ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::Volume",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Volume), &result)
	return &result
}
