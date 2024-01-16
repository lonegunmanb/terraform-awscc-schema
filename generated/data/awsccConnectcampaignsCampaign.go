package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectcampaignsCampaign = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Connect Campaign Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "connect_instance_arn": {
        "computed": true,
        "description": "Amazon Connect Instance Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "dialer_config": {
        "computed": true,
        "description": "The possible types of dialer config parameters",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agentless_dialer_config": {
              "computed": true,
              "description": "Agentless Dialer config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dialing_capacity": {
                    "computed": true,
                    "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "predictive_dialer_config": {
              "computed": true,
              "description": "Predictive Dialer config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bandwidth_allocation": {
                    "computed": true,
                    "description": "The bandwidth allocation of a queue resource.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "dialing_capacity": {
                    "computed": true,
                    "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "progressive_dialer_config": {
              "computed": true,
              "description": "Progressive Dialer config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bandwidth_allocation": {
                    "computed": true,
                    "description": "The bandwidth allocation of a queue resource.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "dialing_capacity": {
                    "computed": true,
                    "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
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
      "name": {
        "computed": true,
        "description": "Amazon Connect Campaign Name",
        "description_kind": "plain",
        "type": "string"
      },
      "outbound_call_config": {
        "computed": true,
        "description": "The configuration used for outbound calls.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "answer_machine_detection_config": {
              "computed": true,
              "description": "The configuration used for answering machine detection during outbound calls",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_answer_machine_detection": {
                    "computed": true,
                    "description": "Flag to decided whether outbound calls should have answering machine detection enabled or not",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "connect_contact_flow_arn": {
              "computed": true,
              "description": "The identifier of the contact flow for the outbound call.",
              "description_kind": "plain",
              "type": "string"
            },
            "connect_queue_arn": {
              "computed": true,
              "description": "The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.",
              "description_kind": "plain",
              "type": "string"
            },
            "connect_source_phone_number": {
              "computed": true,
              "description": "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::ConnectCampaigns::Campaign",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectcampaignsCampaignSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectcampaignsCampaign), &result)
	return &result
}
