package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigOrganizationConformancePack = `{
  "block": {
    "attributes": {
      "conformance_pack_input_parameters": {
        "computed": true,
        "description": "A list of ConformancePackInputParameter objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "delivery_s3_bucket": {
        "computed": true,
        "description": "AWS Config stores intermediate files while processing conformance pack template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_s3_key_prefix": {
        "computed": true,
        "description": "The prefix for the delivery S3 bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "excluded_accounts": {
        "computed": true,
        "description": "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_conformance_pack_name": {
        "description": "The name of the organization conformance pack.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_body": {
        "computed": true,
        "description": "A string containing full conformance pack template body.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_s3_uri": {
        "computed": true,
        "description": "Location of file containing the template body.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Config::OrganizationConformancePack.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConfigOrganizationConformancePackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigOrganizationConformancePack), &result)
	return &result
}
