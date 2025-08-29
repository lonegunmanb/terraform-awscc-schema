package data

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
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_asset_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_domain_unit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_environment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_environment_from_blueprint": {
              "computed": true,
              "description_kind": "plain",
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
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_form_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_glossary": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_project": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "create_project_from_project_profile": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "project_profiles": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "delegate_create_environment_profile": {
              "computed": true,
              "description_kind": "plain",
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
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "override_project_owners": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "include_child_domain_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entity_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "grant_id": {
        "computed": true,
        "description": "The unique identifier of the policy grant returned by the AddPolicyGrant API",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description_kind": "plain",
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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "domain_unit_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "group": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "project": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "project_designation": {
                    "computed": true,
                    "description_kind": "plain",
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
                                "type": "string"
                              },
                              "include_child_domain_units": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "project_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "user": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "all_users_grant_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DataZone::PolicyGrant",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazonePolicyGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazonePolicyGrant), &result)
	return &result
}
