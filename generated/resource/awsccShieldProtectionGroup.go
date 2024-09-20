package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccShieldProtectionGroup = `{
  "block": {
    "attributes": {
      "aggregation": {
        "description": "Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.\n* Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.\n* Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.\n* Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront and origin resources for CloudFront distributions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "members": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the resources to include in the protection group. You must set this when you set ` + "`" + `Pattern` + "`" + ` to ` + "`" + `ARBITRARY` + "`" + ` and you must not set it for any other ` + "`" + `Pattern` + "`" + ` setting.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "pattern": {
        "description": "The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protection_group_arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the protection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "protection_group_id": {
        "description": "The name of the protection group. You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type to include in the protection group. All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set ` + "`" + `Pattern` + "`" + ` to ` + "`" + `BY_RESOURCE_TYPE` + "`" + ` and you must not set it for any other ` + "`" + `Pattern` + "`" + ` setting.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tag key-value pairs for the Protection object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
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
    "description": "A grouping of protected resources so they can be handled as a collective. This resource grouping improves the accuracy of detection and reduces false positives.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccShieldProtectionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccShieldProtectionGroup), &result)
	return &result
}
