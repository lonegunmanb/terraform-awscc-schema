package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ObjectlambdaAccessPoint = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "The status of the Object Lambda alias.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the Object Lambda alias.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description": "The date and time when the Object lambda Access Point was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name you want to assign to this Object lambda Access Point.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_lambda_configuration": {
        "computed": true,
        "description": "The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_features": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "cloudwatch_metrics_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "supporting_access_point": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "transformation_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "actions": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "content_transformation": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "aws_lambda": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "function_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "function_payload": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "policy_status": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "is_public": {
              "computed": true,
              "description": "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "public_access_block_configuration": {
        "computed": true,
        "description": "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "block_public_policy": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ignore_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
              "description_kind": "plain",
              "type": "bool"
            },
            "restrict_public_buckets": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::S3ObjectLambda::AccessPoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3ObjectlambdaAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ObjectlambdaAccessPoint), &result)
	return &result
}
