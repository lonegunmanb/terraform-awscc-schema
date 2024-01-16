package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDocumentationPart = `{
  "block": {
    "attributes": {
      "documentation_part_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the targeted API entity of the to-be-created documentation part.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "method": {
              "computed": true,
              "description": "The HTTP verb of a method. It is a valid field for the API entity types of ` + "`" + `` + "`" + `METHOD` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `PATH_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `QUERY_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_HEADER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_BODY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE_HEADER` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `RESPONSE_BODY` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ` + "`" + `` + "`" + `location` + "`" + `` + "`" + ` attributes, the child entity's ` + "`" + `` + "`" + `method` + "`" + `` + "`" + ` attribute must match that of the parent entity exactly.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the targeted API entity. It is a valid and required field for the API entity types of ` + "`" + `` + "`" + `AUTHORIZER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MODEL` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `PATH_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `QUERY_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_HEADER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_BODY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `RESPONSE_HEADER` + "`" + `` + "`" + `. It is an invalid field for any other entity type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "computed": true,
              "description": "The URL path of the target. It is a valid field for the API entity types of ` + "`" + `` + "`" + `RESOURCE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `METHOD` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `PATH_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `QUERY_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_HEADER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_BODY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE_HEADER` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `RESPONSE_BODY` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `/` + "`" + `` + "`" + ` for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other ` + "`" + `` + "`" + `location` + "`" + `` + "`" + ` attributes, the child entity's ` + "`" + `` + "`" + `path` + "`" + `` + "`" + ` attribute must match that of the parent entity as a prefix.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status_code": {
              "computed": true,
              "description": "The HTTP status code of a response. It is a valid field for the API entity types of ` + "`" + `` + "`" + `RESPONSE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE_HEADER` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `RESPONSE_BODY` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ` + "`" + `` + "`" + `location` + "`" + `` + "`" + ` attributes, the child entity's ` + "`" + `` + "`" + `statusCode` + "`" + `` + "`" + ` attribute must match that of the parent entity exactly.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of API entity to which the documentation content applies. Valid values are ` + "`" + `` + "`" + `API` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AUTHORIZER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MODEL` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESOURCE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `METHOD` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `PATH_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `QUERY_PARAMETER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_HEADER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_BODY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RESPONSE_HEADER` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `RESPONSE_BODY` + "`" + `` + "`" + `. Content inheritance does not apply to any entity of the ` + "`" + `` + "`" + `API` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AUTHORIZER` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `METHOD` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MODEL` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `REQUEST_BODY` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `RESOURCE` + "`" + `` + "`" + ` type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "properties": {
        "description": "The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::DocumentationPart` + "`" + `` + "`" + ` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDocumentationPartSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDocumentationPart), &result)
	return &result
}
