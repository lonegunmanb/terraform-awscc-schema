package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportauthzSupportPermit = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the support permit.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time at which the support permit was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of the support permit.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the support permit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permit": {
        "description": "The grant definition: which actions on which resources, optionally constrained by time conditions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "description": "The set of actions a support permit grants. Exactly one of AllActions or Actions must be provided.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "actions": {
                    "computed": true,
                    "description": "An explicit list of actions to grant.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "all_actions": {
                    "computed": true,
                    "description": "Grants all actions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "conditions": {
              "computed": true,
              "description": "Optional time-bound conditions (at most two).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allow_after": {
                    "computed": true,
                    "description": "The permit is active only after this time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "allow_before": {
                    "computed": true,
                    "description": "The permit is active only before this time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resources": {
              "description": "The set of resources a support permit applies to. Exactly one of AllResourcesInRegion or Resources must be provided.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "all_resources_in_region": {
                    "computed": true,
                    "description": "Applies to all resources in the region.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resources": {
                    "computed": true,
                    "description": "An explicit list of resource ARNs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "permit_id": {
        "computed": true,
        "description": "The service-generated identifier of the support permit (the resource segment of the ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "signing_key_info": {
        "description": "The signing key used by the permit. Exactly one key type must be provided.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key": {
              "description": "The ARN of the KMS key used to sign permit grants.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The current status of the support permit.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_case_display_id": {
        "computed": true,
        "description": "The support case display identifier associated with the permit. When provided, the permit is linked to the specified AWS Support case.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SupportAuthZ::SupportPermit. Represents a support permit that grants AWS support time-bounded access to one or more resources for a set of actions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSupportauthzSupportPermitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportauthzSupportPermit), &result)
	return &result
}
