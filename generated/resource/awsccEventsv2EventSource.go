package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsv2EventSource = `{
  "block": {
    "attributes": {
      "configuration": {
        "description": "The event source configuration. Specify exactly one of AwsServiceEventsConfiguration or PartnerEventsConfiguration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_service_events_configuration": {
              "computed": true,
              "description": "Configuration for forwarding a single AWS service's events.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_service": {
                    "computed": true,
                    "description": "A single AWS service source identifier, for example aws.s3. Wildcards and lists are not allowed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_failure_configuration": {
                    "computed": true,
                    "description": "The destination for events that could not be forwarded.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description": "The ARN of the Amazon SQS standard queue that receives events that could not be forwarded. FIFO queues are not supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "pattern": {
                    "computed": true,
                    "description": "A filter pattern, as a JSON string, that defines which events from the specified AWS service are forwarded to the event bus. Do not include source, account, or region as top-level fields. If you do not specify a pattern, all events from the service are forwarded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "partner_events_configuration": {
              "computed": true,
              "description": "Configuration for forwarding a partner event source's events.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_failure_configuration": {
                    "computed": true,
                    "description": "The destination for events that could not be forwarded, covering both the forwarding target and the managed partner event bus.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description": "The ARN of the Amazon SQS standard queue that receives events that could not be forwarded. FIFO queues are not supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "partner_bus_kms_key_identifier": {
                    "computed": true,
                    "description": "The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on the managed partner event bus. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN. If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt events on the event bus.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partner_event_source_arn": {
                    "computed": true,
                    "description": "The ARN of the partner event source to forward. The partner owns the event source, so the ARN's account segment is empty. Changing this property replaces the event source. Because Name and EventBusArn together identify an event source, and the replacement is created before the old resource is deleted, change Name in the same update.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pattern": {
                    "computed": true,
                    "description": "A filter pattern, as a JSON string, that defines which events from the specified partner event source are forwarded to the event bus. If you do not specify a pattern, all events from the partner event source are forwarded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "creation_time": {
        "computed": true,
        "description": "The time the event source was created, as an ISO 8601 timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the event source. Control characters and Unicode line separators are not allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_bus_arn": {
        "description": "The ARN of the custom event bus the event source forwards onto.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_source_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event source. Its resource segment has the form event-sourcev2/\u003ctype\u003e/\u003cname\u003e/\u003cid\u003e. The type segment is set by the service (aws.service for AWS service events, aws.partner for partner events) and is not part of the event source name. The id segment is a 25-character identifier generated by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time the event source was last modified, as an ISO 8601 timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the event source. The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'. Names cannot begin with the reserved aws. prefix.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revoked": {
        "computed": true,
        "description": "Whether the event bus owner has revoked this event source. Revocation is permanent: a revoked event source cannot be updated, but it can be deleted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the event source: CREATING, ACTIVE, UPDATING, CREATE_FAILED, UPDATE_FAILED, DELETING, or DELETE_FAILED. Revocation by the event bus owner is reported by the Revoked property, not by the state.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Unique per resource; keys are case sensitive. No leading or trailing whitespace (interior whitespace is allowed).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value. May be empty. No leading or trailing whitespace (interior whitespace is allowed).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::EventsV2::EventSource. A managed event source that forwards AWS service events or partner events onto a custom event bus.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsv2EventSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsv2EventSource), &result)
	return &result
}
