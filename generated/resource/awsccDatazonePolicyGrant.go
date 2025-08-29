package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazonePolicyGrant = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Specifies the timestamp at which policy grant member was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "Specifies the user who created the policy grant member.",
        "description_kind": "plain",
        "type": "string"
      },
      "detail": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_to_project_member_pool": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_asset_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_domain_unit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_environment": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_environment_from_blueprint": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_environment_profile": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_unit_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_form_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_glossary": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_project": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "create_project_from_project_profile": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "project_profiles": {
                    "computed": true,
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
              "optional": true
            },
            "delegate_create_environment_profile": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_domain_unit_owners": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "override_project_owners": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "domain_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "grant_id": {
        "computed": true,
        "description": "The unique identifier of the policy grant returned by the AddPolicyGrant API",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain_unit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_unit_designation": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "domain_unit_grant_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "all_domain_units_grant_filter": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "domain_unit_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "group": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "project": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "project_designation": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_grant_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "domain_unit_filter": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "domain_unit": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "include_child_domain_units": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "project_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "user": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "all_users_grant_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Policy Grant in AWS DataZone is an explicit authorization assignment that allows a specific principal (user, group, or project) to perform particular actions (such as creating glossary terms, managing projects, or accessing resources) on governed resources within a certain scope (like a Domain Unit or Project). Policy Grants are essentially the mechanism by which DataZone enforces fine-grained, role-based access control beyond what is possible through AWS IAM alone.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazonePolicyGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazonePolicyGrant), &result)
	return &result
}
