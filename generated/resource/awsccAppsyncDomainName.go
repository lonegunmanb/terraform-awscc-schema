package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncDomainName = `{
  "block": {
    "attributes": {
      "app_sync_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppSync::DomainName",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncDomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncDomainName), &result)
	return &result
}
