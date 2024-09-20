package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderLifecyclePolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the lifecycle policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the lifecycle policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role": {
        "description": "The execution role of the lifecycle policy.",
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
      "name": {
        "description": "The name of the lifecycle policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_details": {
        "description": "The policy details of the lifecycle policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "description": "The action of the policy detail.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_resources": {
                    "computed": true,
                    "description": "The included resources of the policy detail.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "amis": {
                          "computed": true,
                          "description": "Use to configure lifecycle actions on AMIs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "containers": {
                          "computed": true,
                          "description": "Use to configure lifecycle actions on containers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "snapshots": {
                          "computed": true,
                          "description": "Use to configure lifecycle actions on snapshots.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "type": {
                    "description": "The action type of the policy detail.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "exclusion_rules": {
              "computed": true,
              "description": "The exclusion rules to apply of the policy detail.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amis": {
                    "computed": true,
                    "description": "The AMI exclusion rules for the policy detail.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_public": {
                          "computed": true,
                          "description": "Use to apply lifecycle policy actions on whether the AMI is public.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "last_launched": {
                          "computed": true,
                          "description": "Use to apply lifecycle policy actions on AMIs launched before a certain time.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
                                "computed": true,
                                "description": "The value's time unit.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The last launched value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "regions": {
                          "computed": true,
                          "description": "Use to apply lifecycle policy actions on AMIs distributed to a set of regions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "shared_accounts": {
                          "computed": true,
                          "description": "Use to apply lifecycle policy actions on AMIs shared with a set of regions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "tag_map": {
                          "computed": true,
                          "description": "The AMIs to select by tag.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "tag_map": {
                    "computed": true,
                    "description": "The Image Builder tags to filter on.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "filter": {
              "description": "The filters to apply of the policy detail.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "retain_at_least": {
                    "computed": true,
                    "description": "The minimum number of Image Builder resources to retain.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description": "The filter type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The value's time unit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The filter value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "resource_selection": {
        "description": "The resource selection of the lifecycle policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "recipes": {
              "computed": true,
              "description": "The recipes to select.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The recipe name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "semantic_version": {
                    "computed": true,
                    "description": "The recipe version.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "tag_map": {
              "computed": true,
              "description": "The Image Builder resources to select by tag.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "resource_type": {
        "description": "The resource type of the lifecycle policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the lifecycle policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the lifecycle policy.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::LifecyclePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderLifecyclePolicy), &result)
	return &result
}
