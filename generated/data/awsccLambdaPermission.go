package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaPermission = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The action that the principal can use on the function. For example, ` + "`" + `` + "`" + `lambda:InvokeFunction` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `lambda:GetFunction` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_source_token": {
        "computed": true,
        "description": "For Alexa Smart Home functions, a token that the invoker must supply.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "computed": true,
        "description": "The name of the Lambda function, version, or alias.\n  **Name formats**\n +   *Function name* ? ` + "`" + `` + "`" + `my-function` + "`" + `` + "`" + ` (name-only), ` + "`" + `` + "`" + `my-function:v1` + "`" + `` + "`" + ` (with alias).\n  +   *Function ARN* ? ` + "`" + `` + "`" + `arn:aws:lambda:us-west-2:123456789012:function:my-function` + "`" + `` + "`" + `.\n  +   *Partial ARN* ? ` + "`" + `` + "`" + `123456789012:function:my-function` + "`" + `` + "`" + `.\n  \n You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_url_auth_type": {
        "computed": true,
        "description": "The type of authentication that your function URL uses. Set to ` + "`" + `` + "`" + `AWS_IAM` + "`" + `` + "`" + ` if you want to restrict access to authenticated users only. Set to ` + "`" + `` + "`" + `NONE` + "`" + `` + "`" + ` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The AWS-service or AWS-account that invokes the function. If you specify a service, use ` + "`" + `` + "`" + `SourceArn` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SourceAccount` + "`" + `` + "`" + ` to limit who can invoke the function through that service.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_org_id": {
        "computed": true,
        "description": "The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_account": {
        "computed": true,
        "description": "For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ` + "`" + `` + "`" + `SourceArn` + "`" + `` + "`" + ` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.\n Note that Lambda configures the comparison using the ` + "`" + `` + "`" + `StringLike` + "`" + `` + "`" + ` operator.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::Permission",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaPermission), &result)
	return &result
}
