package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2EnclaveCertificateIamRoleAssociation = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_s3_bucket_name": {
        "computed": true,
        "description": "The name of the Amazon S3 bucket to which the certificate was uploaded.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_s3_object_key": {
        "computed": true,
        "description": "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_kms_key_id": {
        "computed": true,
        "description": "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::EnclaveCertificateIamRoleAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2EnclaveCertificateIamRoleAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2EnclaveCertificateIamRoleAssociation), &result)
	return &result
}
