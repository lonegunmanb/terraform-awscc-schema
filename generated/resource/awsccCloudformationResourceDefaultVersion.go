package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationResourceDefaultVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "type_name": {
        "computed": true,
        "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_version_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "The ID of an existing version of the resource to set as the default.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The default version of a resource that has been registered in the CloudFormation Registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationResourceDefaultVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationResourceDefaultVersion), &result)
	return &result
}
