package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDomainNameAccessAssociation = `{
  "block": {
    "attributes": {
      "access_association_source": {
        "description": "The source of the domain name access association resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_association_source_type": {
        "description": "The source type of the domain name access association resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_name_access_association_arn": {
        "computed": true,
        "description": "The amazon resource name (ARN) of the domain name access association resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_arn": {
        "description": "The amazon resource name (ARN) of the domain name resource.",
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
      "tags": {
        "computed": true,
        "description": "An array of arbitrary tags (key-value pairs) to associate with the domainname access association.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::ApiGateway::DomainNameAccessAssociation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDomainNameAccessAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDomainNameAccessAssociation), &result)
	return &result
}
