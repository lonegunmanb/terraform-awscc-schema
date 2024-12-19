package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotCommand = `{
  "block": {
    "attributes": {
      "command_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the command.",
        "description_kind": "plain",
        "type": "string"
      },
      "command_id": {
        "computed": true,
        "description": "The unique identifier for the command.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the command was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deprecated": {
        "computed": true,
        "description": "A flag indicating whether the command is deprecated.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "The description of the command.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the command.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The date and time when the command was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "mandatory_parameters": {
        "computed": true,
        "description": "The list of mandatory parameters for the command.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "b": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "bin": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "d": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "i": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "l": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ul": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "b": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "bin": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "d": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "i": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "l": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ul": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "namespace": {
        "computed": true,
        "description": "The namespace to which the command belongs.",
        "description_kind": "plain",
        "type": "string"
      },
      "payload": {
        "computed": true,
        "description": "The payload associated with the command.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "content_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "pending_deletion": {
        "computed": true,
        "description": "A flag indicating whether the command is pending deletion.",
        "description_kind": "plain",
        "type": "bool"
      },
      "role_arn": {
        "computed": true,
        "description": "The customer role associated with the command.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be associated with the command.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoT::Command",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotCommandSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotCommand), &result)
	return &result
}
