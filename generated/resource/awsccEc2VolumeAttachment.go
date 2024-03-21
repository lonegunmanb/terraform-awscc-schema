package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The ID of the instance to which the volume attaches. This value can be a reference to an [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource, or it can be the physical ID of an existing EC2 instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_id": {
        "description": "The ID of the Amazon EBS volume. The volume and instance must be within the same Availability Zone. This value can be a reference to an [AWS::EC2::Volume](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html) resource, or it can be the volume ID of an existing Amazon EBS volume.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Attaches an Amazon EBS volume to a running instance and exposes it to the instance with the specified device name.\n Before this resource can be deleted (and therefore the volume detached), you must first unmount the volume in the instance. Failure to do so results in the volume being stuck in the busy state while it is trying to detach, which could possibly damage the file system or the data it contains.\n If an Amazon EBS volume is the root device of an instance, it cannot be detached while the instance is in the \"running\" state. To detach the root volume, stop the instance first.\n If the root volume is detached from an instance with an MKT product code, then the product codes from that volume are no longer associated with the instance.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VolumeAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VolumeAttachment), &result)
	return &result
}
