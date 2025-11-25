package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerRotationSchedule = `{
  "block": {
    "attributes": {
      "external_secret_rotation_metadata": {
        "computed": true,
        "description": "The list of metadata needed to successfully rotate a managed external secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the metadata item. You can specify a value that's 1 to 256 characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the metadata item. You can specify a value that's 1 to 2048 characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "external_secret_rotation_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that is used by Secrets Manager to rotate a managed external secret.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_rotation_lambda": {
        "computed": true,
        "description": "Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclude_characters": {
              "computed": true,
              "description": "A string of the characters that you don't want in the password.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key that Secrets Manager uses to encrypt the secret. If you don't specify this value, then Secrets Manager uses the key aws/secretsmanager. If aws/secretsmanager doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "master_secret_arn": {
              "computed": true,
              "description": "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "master_secret_kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rotation_lambda_name": {
              "computed": true,
              "description": "The name of the Lambda rotation function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rotation_type": {
              "computed": true,
              "description": "The type of rotation template to use",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime": {
              "computed": true,
              "description": "The python runtime associated with the Lambda function",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "superuser_secret_arn": {
              "computed": true,
              "description": "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "superuser_secret_kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_security_group_ids": {
              "computed": true,
              "description": "A comma-separated list of security group IDs applied to the target database.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_subnet_ids": {
              "computed": true,
              "description": "A comma separated list of VPC subnet IDs of the target database network. The Lambda rotation function is in the same subnet group.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rotate_immediately_on_update": {
        "computed": true,
        "description": "Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "rotation_lambda_arn": {
        "computed": true,
        "description": "The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rotation_rules": {
        "computed": true,
        "description": "A structure that defines the rotation configuration for this secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "automatically_after_days": {
              "computed": true,
              "description": "The number of days between automatic scheduled rotations of the secret. You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "duration": {
              "computed": true,
              "description": "The length of the rotation window in hours, for example 3h for a three hour window. Secrets Manager rotates your secret at any time during this window. The window must not extend into the next rotation window or the next UTC day. The window starts according to the ScheduleExpression. If you don't specify a Duration, for a ScheduleExpression in hours, the window automatically closes after one hour. For a ScheduleExpression in days, the window automatically closes at the end of the UTC day.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression": {
              "computed": true,
              "description": "A cron() or rate() expression that defines the schedule for rotating your secret. Secrets Manager rotation schedules use UTC time zone.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "rotation_schedule_id": {
        "computed": true,
        "description": "The ARN of the secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
        "description": "The ARN or name of the secret to rotate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SecretsManager::RotationSchedule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecretsmanagerRotationScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerRotationSchedule), &result)
	return &result
}
