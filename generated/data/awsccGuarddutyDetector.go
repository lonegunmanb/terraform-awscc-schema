package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyDetector = `{
  "block": {
    "attributes": {
      "data_sources": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kubernetes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "audit_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable": {
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
            "malware_protection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scan_ec_2_instance_with_findings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ebs_volumes": {
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
            "s3_logs": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable": {
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
      "enable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "features": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "finding_publishing_frequency": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
    "description": "Data Source schema for AWS::GuardDuty::Detector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyDetector), &result)
	return &result
}
