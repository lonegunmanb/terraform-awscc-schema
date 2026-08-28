package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDataexchangeEventAction = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "What occurs after a certain event.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "export_revision_to_s3": {
              "computed": true,
              "description": "Details of the operation to be performed by the job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption": {
                    "computed": true,
                    "description": "Encryption configuration of the export job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_key_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the AWS KMS key you want to use to encrypt the Amazon S3 objects.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The type of server side encryption used for encrypting the objects in Amazon S3.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "revision_destination": {
                    "computed": true,
                    "description": "A revision destination is the Amazon S3 bucket folder destination to where the export will be sent.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The Amazon S3 bucket that is the destination for the event action.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key_pattern": {
                          "computed": true,
                          "description": "A string representing the pattern for generated names of the individual assets in the revision.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "The ARN for the event action.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the event action was created, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "event": {
        "computed": true,
        "description": "What occurs to start an action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "revision_published": {
              "computed": true,
              "description": "Information about the published revision.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_set_id": {
                    "computed": true,
                    "description": "The data set ID of the published revision.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "event_action_id": {
        "computed": true,
        "description": "The unique identifier for the event action.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the event action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time that the event action was last updated, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataExchange::EventAction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDataexchangeEventActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDataexchangeEventAction), &result)
	return &result
}
