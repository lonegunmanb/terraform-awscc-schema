package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeSipMediaApplication = `{
  "block": {
    "attributes": {
      "aws_region": {
        "computed": true,
        "description": "The AWS Region in which the SIP media application is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The SIP media application creation timestamp, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description": "List of endpoints (Lambda ARNs) specified for the SIP media application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lambda_arn": {
              "computed": true,
              "description": "Valid Amazon Resource Name (ARN) of the Lambda function.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
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
        "description": "The name of the SIP media application.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The time at which the SIP media application was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Chime::SipMediaApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChimeSipMediaApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeSipMediaApplication), &result)
	return &result
}
