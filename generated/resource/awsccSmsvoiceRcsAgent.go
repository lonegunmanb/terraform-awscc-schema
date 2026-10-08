package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceRcsAgent = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The time when the RCS agent was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "When set to true the RCS agent can't be deleted. By default this is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "opt_out_list_name": {
        "computed": true,
        "description": "The name of the opt-out list associated with the RCS agent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The unique identifier of the pool associated with the RCS agent. This is populated only while the agent is associated with a pool, which is done through the pool APIs rather than through this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rcs_agent_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the RCS agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "rcs_agent_id": {
        "computed": true,
        "description": "The unique identifier for the RCS agent, in the format rcs- followed by a hexadecimal string.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_managed_opt_outs_enabled": {
        "computed": true,
        "description": "When set to true you're responsible for responding to HELP and STOP requests, and for tracking and honoring opt-out requests. By default this is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description": "The current status of the RCS agent. CREATED means the agent exists but no registration has been submitted. PENDING means a registration is awaiting processing. TESTING means the testing registration is approved and the agent can message registered test devices. PARTIAL means at least one country launch is active. ACTIVE means all submitted country launches are active. DELETED means the agent has been deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the RCS agent.",
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
      "testing_agent": {
        "computed": true,
        "description": "Information about the testing agent (RCS for Business ID) associated with the RCS agent. A testing agent is created by submitting a testing registration, and allows sending to registered test devices without carrier approval.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "registration_id": {
              "computed": true,
              "description": "The unique identifier of the testing registration that created the testing agent.",
              "description_kind": "plain",
              "type": "string"
            },
            "testing_agent_id": {
              "computed": true,
              "description": "The identifier of the testing agent assigned by the RCS infrastructure provider.",
              "description_kind": "plain",
              "type": "string"
            },
            "testing_agent_status": {
              "computed": true,
              "description": "The status of the testing agent. Named distinctly from the resource-level Status property, which has its own larger set of values.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "two_way_channel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the two way channel where inbound messages are delivered.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_channel_role": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role for the service to assume in order to post inbound messages to the two way channel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_enabled": {
        "computed": true,
        "description": "When set to true two-way messaging is enabled for the RCS agent.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "two_way_media_s3_bucket_name": {
        "computed": true,
        "description": "The name of the Amazon S3 bucket where inbound RCS media objects are written.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_media_s3_key_prefix": {
        "computed": true,
        "description": "The key prefix used for inbound RCS media objects in the Amazon S3 bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_media_s3_role": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role used to write inbound RCS media files to the Amazon S3 bucket. The role must have s3:PutObject permission on the bucket and a trust policy allowing sms-voice.amazonaws.com to assume it.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_rcs_events_enabled": {
        "computed": true,
        "description": "The list of RCS event types enabled for two-way messaging. An empty list disables all event types. The special value ALL enables all current and future event types and must be the only element if used. Requires TwoWayEnabled to be true.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "An AWS RCS Agent in AWS End User Messaging SMS. The RCS agent is the top-level resource representing your brand for RCS messaging, and serves as an origination identity for sending RCS messages. A newly created agent is in CREATED status; it advances through TESTING, PARTIAL and ACTIVE only as testing and country launch registrations are submitted and approved by carriers.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSmsvoiceRcsAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceRcsAgent), &result)
	return &result
}
