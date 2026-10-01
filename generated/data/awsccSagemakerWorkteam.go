package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerWorkteam = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the work team.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "member_definitions": {
        "computed": true,
        "description": "A list of MemberDefinition objects that contains objects that identify the workers that make up the work team.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cognito_member_definition": {
              "computed": true,
              "description": "The Amazon Cognito user group that is part of the work team",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cognito_client_id": {
                    "computed": true,
                    "description": "An identifier for an application client. You must create the app client ID using Amazon Cognito.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cognito_user_group": {
                    "computed": true,
                    "description": "An identifier for a user group.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cognito_user_pool": {
                    "computed": true,
                    "description": "An identifier for a user pool. The user pool must be in the same region as the service that you are calling.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "oidc_member_definition": {
              "computed": true,
              "description": "A list user groups that exist in your OIDC Identity Provider (IdP).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "oidc_groups": {
                    "computed": true,
                    "description": "A list of OIDC group names whose members will be part of this workteam",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "notification_configuration": {
        "computed": true,
        "description": "Configures SNS notifications of available or expiring work items for work teams.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_topic_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications should be published.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "workforce_name": {
        "computed": true,
        "description": "The name of the Workforce",
        "description_kind": "plain",
        "type": "string"
      },
      "workteam_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that identifies the work team.",
        "description_kind": "plain",
        "type": "string"
      },
      "workteam_name": {
        "computed": true,
        "description": "The name of the work team.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Workteam",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerWorkteamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerWorkteam), &result)
	return &result
}
