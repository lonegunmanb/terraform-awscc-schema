package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmPatchBaseline = `{
  "block": {
    "attributes": {
      "approval_rules": {
        "computed": true,
        "description": "A set of rules defining the approval rules for a patch baseline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "patch_rules": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "approve_after_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "approve_until_date": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "compliance_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enable_non_security": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "patch_filter_group": {
                    "computed": true,
                    "description": "The patch filter group that defines the criteria for the rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "patch_filters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "values": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "approved_patches": {
        "computed": true,
        "description": "A list of explicitly approved patches for the baseline.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "approved_patches_compliance_level": {
        "computed": true,
        "description": "Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.",
        "description_kind": "plain",
        "type": "string"
      },
      "approved_patches_enable_non_security": {
        "computed": true,
        "description": "Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.",
        "description_kind": "plain",
        "type": "bool"
      },
      "available_security_updates_compliance_status": {
        "computed": true,
        "description": "The compliance status for vendor recommended security updates that are not approved by this patch baseline.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_baseline": {
        "computed": true,
        "description": "Set the baseline as default baseline. Only registering to default patch baseline is allowed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "The description of the patch baseline.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_filters": {
        "computed": true,
        "description": "A set of global filters used to include patches in the baseline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "patch_filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the patch baseline.",
        "description_kind": "plain",
        "type": "string"
      },
      "operating_system": {
        "computed": true,
        "description": "Defines the operating system the patch baseline applies to. The Default value is WINDOWS.",
        "description_kind": "plain",
        "type": "string"
      },
      "patch_baseline_id": {
        "computed": true,
        "description": "The ID of the patch baseline.",
        "description_kind": "plain",
        "type": "string"
      },
      "patch_groups": {
        "computed": true,
        "description": "PatchGroups is used to associate instances with a specific patch baseline",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "rejected_patches": {
        "computed": true,
        "description": "A list of explicitly rejected patches for the baseline.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "rejected_patches_action": {
        "computed": true,
        "description": "The action for Patch Manager to take on patches included in the RejectedPackages list.",
        "description_kind": "plain",
        "type": "string"
      },
      "sources": {
        "computed": true,
        "description": "Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "products": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SSM::PatchBaseline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmPatchBaselineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmPatchBaseline), &result)
	return &result
}
