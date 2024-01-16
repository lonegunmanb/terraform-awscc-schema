package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightTheme = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "base_theme_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_color_palette": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "colors": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "empty_fill_color": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_max_gradient": {
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
            "sheet": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "tile": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "border": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "show": {
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
                  "tile_layout": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gutter": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "show": {
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
                        "margin": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "show": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "typography": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "font_families": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_family": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
              "optional": true
            },
            "ui_color_palette": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accent": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "accent_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_background": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_background": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success_foreground": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning_foreground": {
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
        "required": true
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "theme_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "base_theme_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_color_palette": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "colors": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "empty_fill_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "min_max_gradient": {
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
                  "sheet": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tile": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "border": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "show": {
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
                        "tile_layout": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "gutter": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "show": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "margin": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "show": {
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "typography": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_families": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_family": {
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
                  "ui_color_palette": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accent": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "accent_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "danger": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "danger_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimension": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimension_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "measure": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "measure_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "primary_background": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "primary_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secondary_background": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secondary_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "success": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "success_foreground": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "warning": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "warning_foreground": {
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
            "created_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "errors": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version_number": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::Theme Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightThemeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightTheme), &result)
	return &result
}
