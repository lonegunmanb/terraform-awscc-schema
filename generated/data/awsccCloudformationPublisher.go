package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationPublisher = `{
  "block": {
    "attributes": {
      "accept_terms_and_conditions": {
        "computed": true,
        "description": "Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf",
        "description_kind": "plain",
        "type": "bool"
      },
      "connection_arn": {
        "computed": true,
        "description": "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider": {
        "computed": true,
        "description": "The type of account used as the identity provider when registering this publisher with CloudFormation.",
        "description_kind": "plain",
        "type": "string"
      },
      "publisher_id": {
        "computed": true,
        "description": "The publisher id assigned by CloudFormation for publishing in this region.",
        "description_kind": "plain",
        "type": "string"
      },
      "publisher_profile": {
        "computed": true,
        "description": "The URL to the publisher's profile with the identity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "publisher_status": {
        "computed": true,
        "description": "Whether the publisher is verified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::Publisher",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationPublisherSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationPublisher), &result)
	return &result
}
