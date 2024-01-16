package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubApp = `{
  "block": {
    "attributes": {
      "app_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the App.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_assessment_schedule": {
        "computed": true,
        "description": "Assessment execution schedule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "app_template_body": {
        "description": "A string containing full ResilienceHub app template body.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "App description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "drift_status": {
        "computed": true,
        "description": "Indicates if compliance drifts (deviations) were detected while running an assessment for your application.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_subscriptions": {
        "computed": true,
        "description": "The list of events you would like to subscribe and get notification for.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_type": {
              "description": "The type of event you would like to subscribe and get notification for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "Unique name to identify an event subscription.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sns_topic_arn": {
              "computed": true,
              "description": "Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission_model": {
        "computed": true,
        "description": "Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cross_account_role_arns": {
              "computed": true,
              "description": "Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts. These ARNs are used for querying purposes while importing resources and assessing your application.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "invoker_role_name": {
              "computed": true,
              "description": "Existing AWS IAM role name in the primary AWS account that will be assumed by AWS Resilience Hub Service Principle to obtain a read-only access to your application resources while running an assessment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Defines how AWS Resilience Hub scans your resources. It can scan for the resources by using a pre-existing role in your AWS account, or by using the credentials of the current IAM user.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resiliency_policy_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_mappings": {
        "description": "An array of ResourceMapping objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "eks_source_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logical_stack_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mapping_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "physical_resource_id": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_account_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "aws_region": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "resource_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "terraform_source_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type Definition for AWS::ResilienceHub::App.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubApp), &result)
	return &result
}
