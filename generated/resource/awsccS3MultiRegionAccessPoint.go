package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3MultiRegionAccessPoint = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of the when the Multi Region Access Point is created",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name you want to assign to this Multi Region Access Point.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_access_block_configuration": {
        "computed": true,
        "description": "The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "block_public_policy": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ignore_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "restrict_public_buckets": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "regions": {
        "description": "The list of buckets that you want to associate this Multi Region Access Point with.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bucket_account_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3MultiRegionAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3MultiRegionAccessPoint), &result)
	return &result
}
