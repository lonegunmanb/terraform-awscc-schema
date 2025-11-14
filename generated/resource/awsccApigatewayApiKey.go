package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayApiKey = `{
  "block": {
    "attributes": {
      "api_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "generate_distinct_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A name for the API key. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_keys": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rest_api_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stage_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
      },
      "value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::ApiKey` + "`" + `` + "`" + ` resource creates a unique key that you can distribute to clients who are executing API Gateway ` + "`" + `` + "`" + `Method` + "`" + `` + "`" + ` resources that require an API key. To specify which API key clients must use, map the API key with the ` + "`" + `` + "`" + `RestApi` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Stage` + "`" + `` + "`" + ` resources that include the methods that require a key.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayApiKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayApiKey), &result)
	return &result
}
