package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverRuleAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of an association between a Resolver rule and a VPC.\n The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and spaces. The name cannot consist of only numbers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolver_rule_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_rule_id": {
        "description": "The ID of the Resolver rule that you associated with the VPC that is specified by ` + "`" + `` + "`" + `VPCId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC that you associated the Resolver rule with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "In the response to an [AssociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html), [DisassociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html), or [ListResolverRuleAssociations](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html) request, provides information about an association between a resolver rule and a VPC. The association determines which DNS queries that originate in the VPC are forwarded to your network.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverResolverRuleAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverRuleAssociation), &result)
	return &result
}
