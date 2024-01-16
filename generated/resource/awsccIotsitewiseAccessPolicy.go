package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseAccessPolicy = `{
  "block": {
    "attributes": {
      "access_policy_arn": {
        "computed": true,
        "description": "The ARN of the access policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_policy_id": {
        "computed": true,
        "description": "The ID of the access policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_policy_identity": {
        "description": "The identity for this access policy. Choose either a user or a group but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam_role": {
              "computed": true,
              "description": "Contains information for an IAM role identity in an access policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of the IAM role.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "iam_user": {
              "computed": true,
              "description": "Contains information for an IAM user identity in an access policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of the IAM user.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "user": {
              "computed": true,
              "description": "Contains information for a user identity in an access policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The AWS SSO ID of the user.",
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
      "access_policy_permission": {
        "description": "The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_policy_resource": {
        "description": "The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "portal": {
              "computed": true,
              "description": "A portal resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The ID of the portal.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "project": {
              "computed": true,
              "description": "A project resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The ID of the project.",
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
      }
    },
    "description": "Resource schema for AWS::IoTSiteWise::AccessPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseAccessPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseAccessPolicy), &result)
	return &result
}
