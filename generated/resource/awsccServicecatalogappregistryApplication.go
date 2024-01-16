package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogappregistryApplication = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "The name of the application. ",
        "description_kind": "plain",
        "type": "string"
      },
      "application_tag_key": {
        "computed": true,
        "description": "The key of the AWS application tag, which is awsApplication. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_tag_value": {
        "computed": true,
        "description": "The value of the AWS application tag, which is the identifier of an associated resource. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value. ",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the application. ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the application. ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Schema for AWS::ServiceCatalogAppRegistry::Application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogappregistryApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogappregistryApplication), &result)
	return &result
}
