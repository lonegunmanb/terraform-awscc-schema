package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationPublicTypeVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_delivery_bucket": {
        "computed": true,
        "description": "A url to the S3 bucket where logs for the testType run will be available",
        "description_kind": "plain",
        "type": "string"
      },
      "public_type_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
        "description_kind": "plain",
        "type": "string"
      },
      "public_version_number": {
        "computed": true,
        "description": "The version number of a public third-party extension",
        "description_kind": "plain",
        "type": "string"
      },
      "publisher_id": {
        "computed": true,
        "description": "The publisher id assigned by CloudFormation for publishing in this region.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The kind of extension",
        "description_kind": "plain",
        "type": "string"
      },
      "type_name": {
        "computed": true,
        "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
        "description_kind": "plain",
        "type": "string"
      },
      "type_version_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the extension with the versionId.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::PublicTypeVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationPublicTypeVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationPublicTypeVersion), &result)
	return &result
}
