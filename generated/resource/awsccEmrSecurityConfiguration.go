package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrSecurityConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the security configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_configuration": {
        "description": "The security configuration details in JSON format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEmrSecurityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrSecurityConfiguration), &result)
	return &result
}
