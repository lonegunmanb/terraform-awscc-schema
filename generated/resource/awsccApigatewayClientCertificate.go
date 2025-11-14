package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayClientCertificate = `{
  "block": {
    "attributes": {
      "client_certificate_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::ClientCertificate` + "`" + `` + "`" + ` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayClientCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayClientCertificate), &result)
	return &result
}
