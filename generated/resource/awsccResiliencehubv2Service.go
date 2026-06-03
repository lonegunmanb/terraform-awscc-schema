package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubv2Service = `{
  "block": {
    "attributes": {
      "assertions": {
        "computed": true,
        "description": "Assertions associated with this service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "text": {
              "computed": true,
              "description": "The text of the assertion.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "associated_systems": {
        "computed": true,
        "description": "Systems associated with this service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "system_arn": {
              "computed": true,
              "description": "The system ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_journey_ids": {
              "computed": true,
              "description": "User journey IDs associated with this system.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the service was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dependency_discovery": {
        "computed": true,
        "description": "Dependency discovery state.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_policy_values": {
        "computed": true,
        "description": "Effective policy values computed from the associated policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_slo": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "multi_az_dr_approach": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
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
                "nesting_mode": "single"
              }
            },
            "multi_az_rpo": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "multi_az_rto": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "multi_region_dr_approach": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
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
                "nesting_mode": "single"
              }
            },
            "multi_region_rpo": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "multi_region_rto": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_sources": {
        "computed": true,
        "description": "Input sources for this service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_configuration": {
              "computed": true,
              "description": "Resource configuration for an input source. Provide exactly one field.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cfn_stack_arn": {
                    "computed": true,
                    "description": "ARN of a CloudFormation stack.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "design_file_s3_url": {
                    "computed": true,
                    "description": "S3 URL of a design file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "eks": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_arn": {
                          "computed": true,
                          "description": "ARN of the EKS cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "namespaces": {
                          "computed": true,
                          "description": "EKS namespaces.",
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
                  "resource_tags": {
                    "computed": true,
                    "description": "Resource tags to discover resources.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "Tag key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "values": {
                          "computed": true,
                          "description": "Tag values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "tf_state_file_url": {
                    "computed": true,
                    "description": "URL of a Terraform state file.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "kms_key_id": {
        "computed": true,
        "description": "The KMS key ID for encrypting service data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission_model": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cross_account_role_arns": {
              "computed": true,
              "description": "Cross-account role ARNs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cross_account_role_arn": {
                    "computed": true,
                    "description": "ARN of the cross-account IAM role.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "external_id": {
                    "computed": true,
                    "description": "External ID for cross-account access.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "invoker_role_name": {
              "computed": true,
              "description": "Name of the invoker IAM role.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "policy_arn": {
        "computed": true,
        "description": "The ARN of the resilience policy to associate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regions": {
        "description": "AWS regions for the service.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "report_configuration": {
        "computed": true,
        "description": "Configuration for automatic report generation on a Service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "report_output": {
              "computed": true,
              "description": "Output destinations for generated reports.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3": {
                    "computed": true,
                    "description": "S3 configuration for report output.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_owner": {
                          "computed": true,
                          "description": "Account ID of the bucket owner.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_path": {
                          "computed": true,
                          "description": "S3 bucket path where reports will be written.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "service_arn": {
        "computed": true,
        "description": "The ARN of the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the service was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a resilience-managed service with associated systems, input sources, assertions, and service functions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubv2ServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2Service), &result)
	return &result
}
