package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrSecurityConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the security configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_configuration": {
        "computed": true,
        "description": "The security configuration details in JSON format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EMR::SecurityConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEmrSecurityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrSecurityConfiguration), &result)
	return &result
}
