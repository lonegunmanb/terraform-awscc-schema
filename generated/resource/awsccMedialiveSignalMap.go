package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveSignalMap = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "A signal map's ARN (Amazon Resource Name)",
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_alarm_template_group_identifiers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cloudwatch_alarm_template_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A resource's optional description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discovery_entry_point_arn": {
        "description": "A top-level supported AWS resource ARN to discovery a signal map from.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description": "Error message associated with a failed creation or failed update attempt of a signal map.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_bridge_rule_template_group_identifiers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "event_bridge_rule_template_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "failed_media_resource_map": {
        "computed": true,
        "description": "A map representing an incomplete AWS media workflow as a graph.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of a resource used in AWS media workflows.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The logical name of an AWS media resource.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The logical name of an AWS media resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "sources": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of a resource used in AWS media workflows.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The logical name of an AWS media resource.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "map"
        }
      },
      "force_rediscovery": {
        "computed": true,
        "description": "If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.",
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
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_discovered_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_successful_monitor_deployment": {
        "computed": true,
        "description": "Represents the latest successful monitor deployment of a signal map.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "details_uri": {
              "computed": true,
              "description": "URI associated with a signal map's monitor deployment.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "A signal map's monitor deployment status.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "media_resource_map": {
        "computed": true,
        "description": "A map representing an AWS media workflow as a graph.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of a resource used in AWS media workflows.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The logical name of an AWS media resource.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The logical name of an AWS media resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "sources": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of a resource used in AWS media workflows.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The logical name of an AWS media resource.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "map"
        }
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_changes_pending_deployment": {
        "computed": true,
        "description": "If true, there are pending monitor changes for this signal map that can be deployed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "monitor_deployment": {
        "computed": true,
        "description": "Represents the latest monitor deployment of a signal map.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "details_uri": {
              "computed": true,
              "description": "URI associated with a signal map's monitor deployment.",
              "description_kind": "plain",
              "type": "string"
            },
            "error_message": {
              "computed": true,
              "description": "Error message associated with a failed monitor deployment of a signal map.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "A signal map's monitor deployment status.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signal_map_id": {
        "computed": true,
        "description": "A signal map's id.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "A signal map's current status which is dependent on its lifecycle actions or associated jobs.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Represents the tags associated with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Definition of AWS::MediaLive::SignalMap Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveSignalMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveSignalMap), &result)
	return &result
}
