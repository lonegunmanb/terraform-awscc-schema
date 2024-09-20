package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotMitigationAction = `{
  "block": {
    "attributes": {
      "action_name": {
        "computed": true,
        "description": "A unique identifier for the mitigation action.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "action_params": {
        "description": "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_things_to_thing_group_params": {
              "computed": true,
              "description": "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "override_dynamic_groups": {
                    "computed": true,
                    "description": "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "thing_group_names": {
                    "computed": true,
                    "description": "The list of groups to which you want to add the things that triggered the mitigation action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enable_io_t_logging_params": {
              "computed": true,
              "description": "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_level": {
                    "computed": true,
                    "description": " Specifies which types of information are logged.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn_for_logging": {
                    "computed": true,
                    "description": " The ARN of the IAM role used for logging.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "publish_finding_to_sns_params": {
              "computed": true,
              "description": "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "topic_arn": {
                    "computed": true,
                    "description": "The ARN of the topic to which you want to publish the findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "replace_default_policy_version_params": {
              "computed": true,
              "description": "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "template_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "update_ca_certificate_params": {
              "computed": true,
              "description": "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "update_device_certificate_params": {
              "computed": true,
              "description": "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mitigation_action_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mitigation_action_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
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
    "description": "Mitigation actions can be used to take actions to mitigate issues that were found in an Audit finding or Detect violation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotMitigationActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotMitigationAction), &result)
	return &result
}
