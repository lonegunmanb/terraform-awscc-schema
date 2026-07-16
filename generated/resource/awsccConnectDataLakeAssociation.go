package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectDataLakeAssociation = `{
  "block": {
    "attributes": {
      "data_set_id": {
        "description": "The identifier of the analytics data set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The identifier of the Amazon Connect instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_share_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS Resource Access Manager share",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_share_id": {
        "computed": true,
        "description": "The AWS Resource Access Manager share ID",
        "description_kind": "plain",
        "type": "string"
      },
      "target_account_id": {
        "computed": true,
        "description": "The identifier of the target account",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Connect::DataLakeAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectDataLakeAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectDataLakeAssociation), &result)
	return &result
}
