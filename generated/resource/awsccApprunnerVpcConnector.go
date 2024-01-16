package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApprunnerVpcConnector = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "subnets": {
        "description": "A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.",
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
      },
      "vpc_connector_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of this VPC connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_connector_name": {
        "computed": true,
        "description": "A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_connector_revision": {
        "computed": true,
        "description": "The revision of this VPC connector. It's unique among all the active connectors (\"Status\": \"ACTIVE\") that share the same Name.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "The AWS::AppRunner::VpcConnector resource specifies an App Runner VpcConnector.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApprunnerVpcConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApprunnerVpcConnector), &result)
	return &result
}
