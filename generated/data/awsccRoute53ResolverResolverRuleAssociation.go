package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverRuleAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of an association between a Resolver rule and a VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_rule_association_id": {
        "computed": true,
        "description": "Primary Identifier for Resolver Rule Association",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_rule_id": {
        "computed": true,
        "description": "The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC that you associated the Resolver rule with.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::ResolverRuleAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverResolverRuleAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverRuleAssociation), &result)
	return &result
}
