package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationResource = `{
  "block": {
    "attributes": {
      "hybrid_access_enabled": {
        "computed": true,
        "description": "Indicates whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Data Source.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The CloudFormation primary identifier for the LakeFormation::Resource. Mirrors ResourceArn once the resource exists; populated by the handler. Customers should reference the LakeFormation::Resource by ResourceArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The IAM role that registered a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_service_linked_role": {
        "computed": true,
        "description": "Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.",
        "description_kind": "plain",
        "type": "bool"
      },
      "with_federation": {
        "computed": true,
        "description": "Allows Lake Formation to assume a role to access tables in a federated database.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::LakeFormation::Resource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLakeformationResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationResource), &result)
	return &result
}
