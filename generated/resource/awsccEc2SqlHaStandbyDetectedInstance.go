package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SqlHaStandbyDetectedInstance = `{
  "block": {
    "attributes": {
      "ha_status": {
        "computed": true,
        "description": "The SQL Server high availability status of the EC2 instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The ID of the EC2 instance to enable for SQL Server high availability standby detection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the EC2 instance's SQL Server high availability status was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "sql_server_credentials": {
        "computed": true,
        "description": "The ARN of the AWS Secrets Manager secret containing SQL Server access credentials to the EC2 instance. If not specified, AWS Systems Manager agent will use default local user credentials.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_server_license_usage": {
        "computed": true,
        "description": "The SQL Server license type of the EC2 instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SqlHaStandbyDetectedInstance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SqlHaStandbyDetectedInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SqlHaStandbyDetectedInstance), &result)
	return &result
}
