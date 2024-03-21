package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VolumeAttachment = `{
  "block": {
    "attributes": {
      "device": {
        "computed": true,
        "description": "The device name (for example, ` + "`" + `` + "`" + `/dev/sdh` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `xvdh` + "`" + `` + "`" + `).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance to which the volume attaches. This value can be a reference to an [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource, or it can be the physical ID of an existing EC2 instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description": "The ID of the Amazon EBS volume. The volume and instance must be within the same Availability Zone. This value can be a reference to an [AWS::EC2::Volume](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html) resource, or it can be the volume ID of an existing Amazon EBS volume.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VolumeAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VolumeAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VolumeAttachment), &result)
	return &result
}
