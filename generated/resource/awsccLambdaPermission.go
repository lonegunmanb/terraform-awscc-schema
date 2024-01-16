package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaPermission = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The action that the principal can use on the function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_source_token": {
        "computed": true,
        "description": "For Alexa Smart Home functions, a token that must be supplied by the invoker.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_name": {
        "description": "The name of the Lambda function, version, or alias.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_url_auth_type": {
        "computed": true,
        "description": "The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "A statement identifier that differentiates the statement from others in the same policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "description": "The AWS service or account that invokes the function. If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_org_id": {
        "computed": true,
        "description": "The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_account": {
        "computed": true,
        "description": "For Amazon S3, the ID of the account that owns the resource. Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "For AWS services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::Permission",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaPermission), &result)
	return &result
}
