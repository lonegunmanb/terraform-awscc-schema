package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeSipMediaApplication = `{
  "block": {
    "attributes": {
      "aws_region": {
        "description": "The AWS Region in which the SIP media application is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The SIP media application creation timestamp, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "description": "List of endpoints (Lambda ARNs) specified for the SIP media application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lambda_arn": {
              "description": "Valid Amazon Resource Name (ARN) of the Lambda function.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the SIP media application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sip_media_application_arn": {
        "computed": true,
        "description": "The ARN of the SIP media application.",
        "description_kind": "plain",
        "type": "string"
      },
      "sip_media_application_id": {
        "computed": true,
        "description": "The SIP media application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the SIP media application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The time at which the SIP media application was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Chime::SipMediaApplication",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeSipMediaApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeSipMediaApplication), &result)
	return &result
}
