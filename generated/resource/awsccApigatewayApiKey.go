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
        "description": "An MKT customer identifier, when integrating with the AWS SaaS Marketplace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the ApiKey.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Specifies whether the ApiKey can be used by callers.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "generate_distinct_id": {
        "computed": true,
        "description": "Specifies whether (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `) or not (` + "`" + `` + "`" + `false` + "`" + `` + "`" + `) the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
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
        "description": "A name for the API key. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_keys": {
        "computed": true,
        "description": "DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rest_api_id": {
              "computed": true,
              "description": "The string identifier of the associated RestApi.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stage_name": {
              "computed": true,
              "description": "The stage name associated with the stage key.",
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
        "description": "The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `. The tag value can be up to 256 characters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
        "description": "Specifies a value of the API key.",
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
