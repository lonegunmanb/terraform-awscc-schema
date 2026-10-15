package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightBrand = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the brand.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account that owns the brand.",
        "description_kind": "plain",
        "type": "string"
      },
      "brand_definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_theme": {
              "computed": true,
              "description": "The application theme.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "brand_color_palette": {
                    "computed": true,
                    "description": "The color palette.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accent": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "danger": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "dimension": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "info": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "measure": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "primary": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "secondary": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "success": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "warning": {
                          "computed": true,
                          "description": "The palette colors.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background": {
                                "computed": true,
                                "description": "The background color.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "foreground": {
                                "computed": true,
                                "description": "The foreground color.",
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
                  "brand_element_style": {
                    "computed": true,
                    "description": "The element style.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "navbar_style": {
                          "computed": true,
                          "description": "The navigation bar style.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "contextual_navbar": {
                                "computed": true,
                                "description": "The palette colors.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "background": {
                                      "computed": true,
                                      "description": "The background color.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "foreground": {
                                      "computed": true,
                                      "description": "The foreground color.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "global_navbar": {
                                "computed": true,
                                "description": "The palette colors.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "background": {
                                      "computed": true,
                                      "description": "The background color.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "foreground": {
                                      "computed": true,
                                      "description": "The foreground color.",
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
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "brand_name": {
              "computed": true,
              "description": "The name of the brand.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description of the brand.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "brand_id": {
        "computed": true,
        "description": "The ID of the QuickSight brand.",
        "description_kind": "plain",
        "type": "string"
      },
      "brand_status": {
        "computed": true,
        "description": "The status of the brand.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time that the brand was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The last time the brand was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the brand.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "version_id": {
        "computed": true,
        "description": "The ID of the version.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_status": {
        "computed": true,
        "description": "The status of the version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Brand",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightBrandSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightBrand), &result)
	return &result
}
