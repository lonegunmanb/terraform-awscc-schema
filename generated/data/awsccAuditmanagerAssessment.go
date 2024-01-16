package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAuditmanagerAssessment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the assessment.",
        "description_kind": "plain",
        "type": "string"
      },
      "assessment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assessment_reports_destination": {
        "computed": true,
        "description": "The destination in which evidence reports are stored for the specified assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "computed": true,
              "description": "The URL of the specified Amazon S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "destination_type": {
              "computed": true,
              "description": "The destination type, such as Amazon S3.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "aws_account": {
        "computed": true,
        "description": "The AWS account associated with the assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email_address": {
              "computed": true,
              "description": "The unique identifier for the email account.",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "The identifier for the specified AWS account.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the specified AWS account.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "creation_time": {
        "computed": true,
        "description": "The sequence of characters that identifies when the event occurred.",
        "description_kind": "plain",
        "type": "number"
      },
      "delegations": {
        "computed": true,
        "description": "The list of delegations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "assessment_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "assessment_name": {
              "computed": true,
              "description": "The name of the related assessment.",
              "description_kind": "plain",
              "type": "string"
            },
            "comment": {
              "computed": true,
              "description": "The comment related to the delegation.",
              "description_kind": "plain",
              "type": "string"
            },
            "control_set_id": {
              "computed": true,
              "description": "The identifier for the specified control set.",
              "description_kind": "plain",
              "type": "string"
            },
            "created_by": {
              "computed": true,
              "description": "The IAM user or role that performed the action.",
              "description_kind": "plain",
              "type": "string"
            },
            "creation_time": {
              "computed": true,
              "description": "The sequence of characters that identifies when the event occurred.",
              "description_kind": "plain",
              "type": "number"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "last_updated": {
              "computed": true,
              "description": "The sequence of characters that identifies when the event occurred.",
              "description_kind": "plain",
              "type": "number"
            },
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_type": {
              "computed": true,
              "description": " The IAM role type.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "The status of the delegation.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the specified assessment.",
        "description_kind": "plain",
        "type": "string"
      },
      "framework_id": {
        "computed": true,
        "description": "The identifier for the specified framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the related assessment.",
        "description_kind": "plain",
        "type": "string"
      },
      "roles": {
        "computed": true,
        "description": "The list of roles for the specified assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_type": {
              "computed": true,
              "description": " The IAM role type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "scope": {
        "computed": true,
        "description": "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_accounts": {
              "computed": true,
              "description": "The AWS accounts included in scope.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "email_address": {
                    "computed": true,
                    "description": "The unique identifier for the email account.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description": "The identifier for the specified AWS account.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the specified AWS account.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "aws_services": {
              "computed": true,
              "description": "The AWS services included in scope.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "service_name": {
                    "computed": true,
                    "description": "The name of the AWS service.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the specified assessment. ",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::AuditManager::Assessment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAuditmanagerAssessmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAuditmanagerAssessment), &result)
	return &result
}
