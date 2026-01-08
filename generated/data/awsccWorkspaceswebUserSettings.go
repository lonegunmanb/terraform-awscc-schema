package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebUserSettings = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "associated_portal_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "branding_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "color_theme": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "favicon": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "favicon_metadata": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "file_extension": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "last_upload_timestamp": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mime_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "localized_strings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "browser_tab_title": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "contact_button_text": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "contact_link": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "loading_text": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "login_button_text": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "login_description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "login_title": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "welcome_text": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "map"
              }
            },
            "logo": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "logo_metadata": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "file_extension": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "last_upload_timestamp": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mime_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "terms_of_service": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "wallpaper": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "wallpaper_metadata": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "file_extension": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "last_upload_timestamp": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mime_type": {
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
      },
      "cookie_synchronization_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowlist": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "blocklist": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "copy_allowed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deep_link_allowed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disconnect_timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "download_allowed": {
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
      "idle_disconnect_timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "paste_allowed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "print_allowed": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "toolbar_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hidden_toolbar_items": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "max_display_resolution": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "toolbar_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "visual_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "upload_allowed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "web_authn_allowed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkSpacesWeb::UserSettings",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspaceswebUserSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebUserSettings), &result)
	return &result
}
