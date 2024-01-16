package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioStudioComponent = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe configuration of the studio component, based on component type.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_directory_configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe configuration for a Microsoft Active Directory (Microsoft AD) studio\n            resource.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "computer_attributes": {
                    "computed": true,
                    "description": "\u003cp\u003eA collection of custom attributes for an Active Directory computer.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name for the LDAP attribute.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "\u003cp\u003eThe value for the LDAP attribute.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "directory_id": {
                    "computed": true,
                    "description": "\u003cp\u003eThe directory ID of the Directory Service for Microsoft Active Directory to access\n            using this studio component.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_unit_distinguished_name": {
                    "computed": true,
                    "description": "\u003cp\u003eThe distinguished name (DN) and organizational unit (OU) of an Active Directory\n            computer.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "compute_farm_configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe configuration for a render farm that is associated with a studio resource.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "active_directory_user": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of an Active Directory user that is used on ComputeFarm worker\n            instances.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "\u003cp\u003eThe endpoint of the ComputeFarm that is accessed by the studio component\n            resource.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "license_service_configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe configuration for a license service that is associated with a studio\n            resource.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "endpoint": {
                    "computed": true,
                    "description": "\u003cp\u003eThe endpoint of the license service that is accessed by the studio component\n            resource.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "shared_file_system_configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe configuration for a shared file storage system that is associated with a studio\n            resource.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "endpoint": {
                    "computed": true,
                    "description": "\u003cp\u003eThe endpoint of the shared file system that is accessed by the studio component\n            resource.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_system_id": {
                    "computed": true,
                    "description": "\u003cp\u003eThe unique identifier for a file system.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "linux_mount_point": {
                    "computed": true,
                    "description": "\u003cp\u003eThe mount location for a shared file system on a Linux virtual workstation.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "share_name": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of the file share.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "windows_mount_drive": {
                    "computed": true,
                    "description": "\u003cp\u003eThe mount location for a shared file system on a Windows virtual workstation.\u003c/p\u003e",
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
      },
      "description": {
        "computed": true,
        "description": "\u003cp\u003eThe description.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ec_2_security_group_ids": {
        "computed": true,
        "description": "\u003cp\u003eThe EC2 security groups that control access to the studio component.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "initialization_scripts": {
        "computed": true,
        "description": "\u003cp\u003eInitialization scripts for studio components.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_profile_protocol_version": {
              "computed": true,
              "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "platform": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "run_context": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script": {
              "computed": true,
              "description": "\u003cp\u003eThe initialization script.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "name": {
        "description": "\u003cp\u003eThe name for the studio component.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "script_parameters": {
        "computed": true,
        "description": "\u003cp\u003eParameters for the studio component scripts.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eA script parameter key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eA script parameter value.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "secure_initialization_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "studio_component_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_id": {
        "description": "\u003cp\u003eThe studio ID. \u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subtype": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Represents a studio component that connects a non-Nimble Studio resource in your account to your studio",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNimblestudioStudioComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioStudioComponent), &result)
	return &result
}
