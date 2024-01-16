package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewisePortal = `{
  "block": {
    "attributes": {
      "alarms": {
        "computed": true,
        "description": "Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarm_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.",
              "description_kind": "plain",
              "type": "string"
            },
            "notification_lambda_arn": {
              "computed": true,
              "description": "The ARN of the AWS Lambda function that manages alarm notifications. For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.",
              "description_kind": "plain",
              "type": "string"
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
      "notification_sender_email": {
        "computed": true,
        "description": "The email address that sends alarm notifications.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description": "The ARN of the portal, which has the following format.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_auth_mode": {
        "computed": true,
        "description": "The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_client_id": {
        "computed": true,
        "description": "The AWS SSO application generated client ID (used with AWS SSO APIs).",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_contact_email": {
        "computed": true,
        "description": "The AWS administrator's contact email address.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_description": {
        "computed": true,
        "description": "A description for the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_id": {
        "computed": true,
        "description": "The ID of the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_name": {
        "computed": true,
        "description": "A friendly name for the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_start_url": {
        "computed": true,
        "description": "The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the portal.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTSiteWise::Portal",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotsitewisePortalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewisePortal), &result)
	return &result
}
