package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ExpressAccessPoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified accesspoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "computed": true,
        "description": "The name of the bucket that you want to associate this Access Point with.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_account_id": {
        "computed": true,
        "description": "The AWS account ID associated with the S3 bucket associated with this access point.",
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
        "description": "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. For directory buckets, the access point name must consist of a base name that you provide and su?x that includes the ZoneID (AWS Availability Zone or Local Zone) of your bucket location, followed by --xa-s3.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_origin": {
        "computed": true,
        "description": "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "The Access Point Policy you want to apply to this access point.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_access_block_configuration": {
        "computed": true,
        "description": "The PublicAccessBlock configuration that you want to apply to this Access Point.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
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
      },
      "scope": {
        "computed": true,
        "description": "For directory buckets, you can ?lter access control to speci?c pre?xes, API operations, or a combination of both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "permissions": {
              "computed": true,
              "description": "You can include one or more API operations as permissions",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "prefixes": {
              "computed": true,
              "description": "You can specify any amount of pre?xes, but the total length of characters of all pre?xes must be less than 256 bytes in size.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
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
          "nesting_mode": "set"
        }
      },
      "vpc_configuration": {
        "computed": true,
        "description": "If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_id": {
              "computed": true,
              "description": "If this field is specified, this access point will only allow connections from the specified VPC ID.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::S3Express::AccessPoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3ExpressAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ExpressAccessPoint), &result)
	return &result
}
