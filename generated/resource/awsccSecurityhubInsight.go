package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubInsight = `{
  "block": {
    "attributes": {
      "filters": {
        "description": "One or more attributes used to filter the findings included in the insight",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_account_id": {
              "computed": true,
              "description": "The AWS account ID in which a finding is generated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "aws_account_name": {
              "computed": true,
              "description": "The name of the AWS account in which a finding is generated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "company_name": {
              "computed": true,
              "description": "The name of the findings provider (company) that owns the solution (product) that generates findings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "compliance_associated_standards_id": {
              "computed": true,
              "description": "The unique identifier of a standard in which a control is enabled.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "compliance_security_control_id": {
              "computed": true,
              "description": "The unique identifier of a control across standards.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "compliance_security_control_parameters_name": {
              "computed": true,
              "description": "The name of a security control parameter.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "compliance_security_control_parameters_value": {
              "computed": true,
              "description": "The current value of a security control parameter.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "compliance_status": {
              "computed": true,
              "description": "Exclusive to findings that are generated as the result of a check run against a specific rule in a supported standard.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "confidence": {
              "computed": true,
              "description": "A finding's confidence.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "created_at": {
              "computed": true,
              "description": "An ISO8601-formatted timestamp that indicates when the security findings provider captured the potential security issue that a finding captured.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "criticality": {
              "computed": true,
              "description": "The level of importance assigned to the resources associated with the finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "description": {
              "computed": true,
              "description": "A finding's description.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_confidence": {
              "computed": true,
              "description": "The finding provider value for the finding confidence.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_criticality": {
              "computed": true,
              "description": "The finding provider value for the level of importance assigned to the resources associated with the findings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_related_findings_id": {
              "computed": true,
              "description": "The finding identifier of a related finding that is identified by the finding provider.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_related_findings_product_arn": {
              "computed": true,
              "description": "The ARN of the solution that generated a related finding that is identified by the finding provider.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_severity_label": {
              "computed": true,
              "description": "The finding provider value for the severity label.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_severity_original": {
              "computed": true,
              "description": "The finding provider's original value for the severity.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_provider_fields_types": {
              "computed": true,
              "description": "One or more finding types that the finding provider assigned to the finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "first_observed_at": {
              "computed": true,
              "description": "An ISO8601-formatted timestamp that indicates when the security findings provider first observed the potential security issue that a finding captured.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "generator_id": {
              "computed": true,
              "description": "The identifier for the solution-specific component (a discrete unit of logic) that generated a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "id": {
              "computed": true,
              "description": "The security findings provider-specific identifier for a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "keyword": {
              "computed": true,
              "description": "A keyword for a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "description": "A value for the keyword.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "last_observed_at": {
              "computed": true,
              "description": "An ISO8601-formatted timestamp that indicates when the security findings provider most recently observed the potential security issue that a finding captured.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "malware_name": {
              "computed": true,
              "description": "The name of the malware that was observed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "malware_path": {
              "computed": true,
              "description": "The filesystem path of the malware that was observed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "malware_state": {
              "computed": true,
              "description": "The state of the malware that was observed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "malware_type": {
              "computed": true,
              "description": "The type of the malware that was observed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_destination_domain": {
              "computed": true,
              "description": "The destination domain of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_destination_ip_v4": {
              "computed": true,
              "description": "The destination IPv4 address of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_destination_ip_v6": {
              "computed": true,
              "description": "The destination IPv6 address of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_destination_port": {
              "computed": true,
              "description": "The destination port of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_direction": {
              "computed": true,
              "description": "Indicates the direction of network traffic associated with a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_protocol": {
              "computed": true,
              "description": "The protocol of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_source_domain": {
              "computed": true,
              "description": "The source domain of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_source_ip_v4": {
              "computed": true,
              "description": "The source IPv4 address of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_source_ip_v6": {
              "computed": true,
              "description": "The source IPv6 address of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_source_mac": {
              "computed": true,
              "description": "The source media access control (MAC) address of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_source_port": {
              "computed": true,
              "description": "The source port of network-related information about a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "note_text": {
              "computed": true,
              "description": "The text of a note.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "note_updated_at": {
              "computed": true,
              "description": "The timestamp of when the note was updated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "note_updated_by": {
              "computed": true,
              "description": "The principal that created a note.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_launched_at": {
              "computed": true,
              "description": "A timestamp that identifies when the process was launched.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_name": {
              "computed": true,
              "description": "The name of the process.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_parent_pid": {
              "computed": true,
              "description": "The parent process ID.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_path": {
              "computed": true,
              "description": "The path to the process executable.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_pid": {
              "computed": true,
              "description": "The process ID.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "process_terminated_at": {
              "computed": true,
              "description": "A timestamp that identifies when the process was terminated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "product_arn": {
              "computed": true,
              "description": "The ARN generated by Security Hub that uniquely identifies a third-party company (security findings provider) after this provider's product (solution that generates findings) is registered with Security Hub.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "product_fields": {
              "computed": true,
              "description": "A data type where security findings providers can include additional solution-specific details that aren't part of the defined AwsSecurityFinding format.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to the key value when filtering Security Hub findings with a map filter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "product_name": {
              "computed": true,
              "description": "The name of the solution (product) that generates findings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "recommendation_text": {
              "computed": true,
              "description": "The recommendation of what to do about the issue described in a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "record_state": {
              "computed": true,
              "description": "The updated record state for the finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "region": {
              "computed": true,
              "description": "The Region from which the finding was generated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "related_findings_id": {
              "computed": true,
              "description": "The solution-generated identifier for a related finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "related_findings_product_arn": {
              "computed": true,
              "description": "The ARN of the solution that generated a related finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_application_arn": {
              "computed": true,
              "description": "The ARN of the application that is related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_application_name": {
              "computed": true,
              "description": "The name of the application that is related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_iam_instance_profile_arn": {
              "computed": true,
              "description": "The IAM profile ARN of the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_image_id": {
              "computed": true,
              "description": "The Amazon Machine Image (AMI) ID of the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_ip_v4_addresses": {
              "computed": true,
              "description": "The IPv4 addresses associated with the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_ip_v6_addresses": {
              "computed": true,
              "description": "The IPv6 addresses associated with the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A finding's CIDR value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_key_name": {
              "computed": true,
              "description": "The key name associated with the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_launched_at": {
              "computed": true,
              "description": "The date and time the instance was launched.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_subnet_id": {
              "computed": true,
              "description": "The identifier of the subnet that the instance was launched in.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_type": {
              "computed": true,
              "description": "The instance type of the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_ec_2_instance_vpc_id": {
              "computed": true,
              "description": "The identifier of the VPC that the instance was launched in.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_iam_access_key_created_at": {
              "computed": true,
              "description": "The creation date/time of the IAM access key related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_iam_access_key_principal_name": {
              "computed": true,
              "description": "The name of the principal that is associated with an IAM access key.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_iam_access_key_status": {
              "computed": true,
              "description": "The status of the IAM access key related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_iam_access_key_user_name": {
              "computed": true,
              "description": "The user associated with the IAM access key related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_iam_user_user_name": {
              "computed": true,
              "description": "The name of an IAM user.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_s3_bucket_owner_id": {
              "computed": true,
              "description": "The canonical user ID of the owner of the S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_aws_s3_bucket_owner_name": {
              "computed": true,
              "description": "The display name of the owner of the S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_container_image_id": {
              "computed": true,
              "description": "The identifier of the image related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_container_image_name": {
              "computed": true,
              "description": "The name of the image related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_container_launched_at": {
              "computed": true,
              "description": "A timestamp that identifies when the container was started.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_container_name": {
              "computed": true,
              "description": "The name of the container related to a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_details_other": {
              "computed": true,
              "description": "The details of a resource that doesn't have a specific subfield for the resource type defined.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to the key value when filtering Security Hub findings with a map filter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_id": {
              "computed": true,
              "description": "The canonical identifier for the given resource type.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_partition": {
              "computed": true,
              "description": "The canonical AWS partition name that the Region is assigned to.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_region": {
              "computed": true,
              "description": "The canonical AWS external Region name where this resource is located.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_tags": {
              "computed": true,
              "description": "A list of AWS tags associated with a resource at the time the finding was processed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to the key value when filtering Security Hub findings with a map filter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_type": {
              "computed": true,
              "description": "Specifies the type of the resource that details are provided for.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "sample": {
              "computed": true,
              "description": "Indicates whether or not sample findings are included in the filter results.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "description": "The value of the boolean.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "severity_label": {
              "computed": true,
              "description": "The label of a finding's severity.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "severity_normalized": {
              "computed": true,
              "description": "The normalized severity of a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "severity_product": {
              "computed": true,
              "description": "The native severity as defined by the security findings provider's solution that generated the finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description": "The equal-to condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description": "The greater-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description": "The less-than-equal condition to be applied to a single field when querying for findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "source_url": {
              "computed": true,
              "description": "A URL that links to a page about the current finding in the security findings provider's solution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_category": {
              "computed": true,
              "description": "The category of a threat intelligence indicator.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_last_observed_at": {
              "computed": true,
              "description": "A timestamp that identifies the last observation of a threat intelligence indicator.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_source": {
              "computed": true,
              "description": "The source of the threat intelligence.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_source_url": {
              "computed": true,
              "description": "The URL for more details from the source of the threat intelligence.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_type": {
              "computed": true,
              "description": "The type of a threat intelligence indicator.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threat_intel_indicator_value": {
              "computed": true,
              "description": "The value of a threat intelligence indicator.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "title": {
              "computed": true,
              "description": "A finding's title.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description": "A finding type in the format of namespace/category/classifier that classifies a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "updated_at": {
              "computed": true,
              "description": "An ISO8601-formatted timestamp that indicates when the security findings provider last updated the finding record.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "date_range": {
                    "computed": true,
                    "description": "A date range for the date filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "unit": {
                          "description": "A date range unit for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A date range value for the date filter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "end": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "The date and time, in UTC and ISO 8601 format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "user_defined_fields": {
              "computed": true,
              "description": "A list of name/value string pairs associated with the finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to the key value when filtering Security Hub findings with a map filter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "verification_state": {
              "computed": true,
              "description": "The veracity of a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "vulnerabilities_exploit_available": {
              "computed": true,
              "description": "Indicates whether a software vulnerability in your environment has a known exploit.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "vulnerabilities_fix_available": {
              "computed": true,
              "description": "Indicates whether a vulnerability is fixed in a newer version of the affected software packages.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "workflow_state": {
              "computed": true,
              "description": "The workflow state of a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "workflow_status": {
              "computed": true,
              "description": "The status of the investigation into a finding.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description": "The condition to apply to a string value when filtering Security Hub findings.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Non-empty string definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "group_by_attribute": {
        "description": "The grouping attribute for the insight's findings",
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
      "insight_arn": {
        "computed": true,
        "description": "The ARN of a Security Hub insight",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of a Security Hub insight",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubInsightSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubInsight), &result)
	return &result
}
