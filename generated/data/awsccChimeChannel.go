package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeChannel = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "computed": true,
        "description": "The ARN of the AppInstance that contains the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_flow_arn": {
        "computed": true,
        "description": "The ARN of the channel flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_id": {
        "computed": true,
        "description": "The ID of the channel. When omitted, the service generates a UUID.",
        "description_kind": "plain",
        "type": "string"
      },
      "chime_bearer": {
        "computed": true,
        "description": "The ARN of the AppInstanceUser or AppInstanceBot that performs every operation on this channel. Whichever of the two creates a channel automatically becomes one of its moderators, so the same ARN can subsequently read, update and delete the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The AppInstanceUser or AppInstanceBot that created the channel.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The ARN in an identity.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name in an identity.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_timestamp": {
        "computed": true,
        "description": "The time at which the channel was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_channel_configuration": {
        "computed": true,
        "description": "The attributes required to configure and create an elastic channel. An elastic channel must use RESTRICTED mode, cannot be created with MemberArns, and is available only in some regions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_sub_channels": {
              "computed": true,
              "description": "The maximum number of SubChannels allowed in the elastic channel.",
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_membership_percentage": {
              "computed": true,
              "description": "The minimum allowed percentage of TargetMembershipsPerSubChannel users, used to balance members across SubChannels.",
              "description_kind": "plain",
              "type": "number"
            },
            "target_memberships_per_sub_channel": {
              "computed": true,
              "description": "The maximum number of members allowed in a SubChannel.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "expiration_settings": {
        "computed": true,
        "description": "Settings that control the interval after which the channel is automatically deleted.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expiration_criterion": {
              "computed": true,
              "description": "The condition the expiration period is measured from.",
              "description_kind": "plain",
              "type": "string"
            },
            "expiration_days": {
              "computed": true,
              "description": "The period in days after which the system automatically deletes the channel.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_message_timestamp": {
        "computed": true,
        "description": "The time at which a member sent the last message in the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The time at which the channel was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "member_arns": {
        "computed": true,
        "description": "The ARNs of the AppInstanceUsers to add to the channel as members when it is created. Cannot be combined with ElasticChannelConfiguration.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "metadata": {
        "computed": true,
        "description": "The metadata of the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description": "The channel mode. In an UNRESTRICTED channel, members can add themselves and other members; in a RESTRICTED channel, only administrators and moderators can add members. An elastic channel must be RESTRICTED.",
        "description_kind": "plain",
        "type": "string"
      },
      "moderator_arns": {
        "computed": true,
        "description": "The ARNs of the AppInstanceUsers to add to the channel as moderators when it is created.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "privacy": {
        "computed": true,
        "description": "The channel's privacy level. A PUBLIC channel is discoverable by anyone in the AppInstance; a PRIVATE channel is not. Privacy cannot be changed after creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the channel.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in a tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in a tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Chime::Channel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChimeChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeChannel), &result)
	return &result
}
