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
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the theme.\u003c/p\u003e",
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
        "description": "\u003cp\u003eThe theme configuration. This configuration contains all of the display properties for\n            a theme.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_color_palette": {
              "computed": true,
              "description": "\u003cp\u003eThe theme colors that are used for data colors in charts. The colors description is a\n            hexadecimal color code that consists of six alphanumerical characters, prefixed with\n                \u003ccode\u003e#\u003c/code\u003e, for example #37BFF5. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "colors": {
                    "computed": true,
                    "description": "\u003cp\u003eThe hexadecimal codes for the colors.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "empty_fill_color": {
                    "computed": true,
                    "description": "\u003cp\u003eThe hexadecimal code of a color that applies to charts where a lack of data is\n            highlighted.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_max_gradient": {
                    "computed": true,
                    "description": "\u003cp\u003eThe minimum and maximum hexadecimal codes that describe a color gradient. \u003c/p\u003e",
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
              "description": "\u003cp\u003eThe theme display options for sheets. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "background": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "color": {
                          "computed": true,
                          "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "gradient": {
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
                  "tile": {
                    "computed": true,
                    "description": "\u003cp\u003eDisplay options related to tiles on a sheet.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "background_color": {
                          "computed": true,
                          "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "border": {
                          "computed": true,
                          "description": "\u003cp\u003eThe display options for tile borders for visuals.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "color": {
                                "computed": true,
                                "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "show": {
                                "computed": true,
                                "description": "\u003cp\u003eThe option to enable display of borders for visuals.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "width": {
                                "computed": true,
                                "description": "String to encapsulate the most generic way Width can be formatted with whatever units (px, em etc)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "border_radius": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "padding": {
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
                  "tile_layout": {
                    "computed": true,
                    "description": "\u003cp\u003eThe display options for the layout of tiles on a sheet.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gutter": {
                          "computed": true,
                          "description": "\u003cp\u003eThe display options for gutter spacing between tiles on a sheet.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "show": {
                                "computed": true,
                                "description": "\u003cp\u003eThis Boolean value controls whether to display a gutter space between sheet tiles.\n        \u003c/p\u003e",
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
                          "description": "\u003cp\u003eThe display options for margins around the outside edge of sheets.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "show": {
                                "computed": true,
                                "description": "\u003cp\u003eThis Boolean value controls whether to display sheet margins.\u003c/p\u003e",
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
                  "axis_label_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_decoration": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "absolute": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "font_style": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                  },
                  "axis_title_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_decoration": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "absolute": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "font_style": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                  },
                  "data_label_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_decoration": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "absolute": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "font_style": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                  },
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
                  },
                  "legend_title_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_decoration": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "absolute": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "font_style": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                  },
                  "legend_value_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_decoration": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "absolute": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "font_style": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "font_weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                  },
                  "visual_subtitle_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        },
                        "text_alignment": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_transform": {
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
                  "visual_title_font_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        },
                        "text_alignment": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_transform": {
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
            },
            "ui_color_palette": {
              "computed": true,
              "description": "\u003cp\u003eThe theme colors that apply to UI and to charts, excluding data colors. The colors\n            description is a hexadecimal color code that consists of six alphanumerical characters,\n            prefixed with \u003ccode\u003e#\u003c/code\u003e, for example #37BFF5. For more information, see \u003ca href=\"https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html\"\u003eUsing Themes in Amazon Quick\u003c/a\u003e in the \u003ci\u003eAmazon Quick User\n                Guide.\u003c/i\u003e\n         \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accent": {
                    "computed": true,
                    "description": "\u003cp\u003eThis color is that applies to selected states and buttons.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "accent_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            accent color.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger": {
                    "computed": true,
                    "description": "\u003cp\u003eThe color that applies to error messages.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            error color.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension": {
                    "computed": true,
                    "description": "\u003cp\u003eThe color that applies to the names of fields that are identified as\n            dimensions.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            dimension color.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure": {
                    "computed": true,
                    "description": "\u003cp\u003eThe color that applies to the names of fields that are identified as measures.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            measure color.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_background": {
                    "computed": true,
                    "description": "\u003cp\u003eThe background color that applies to visuals and other high emphasis UI.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe color of text and other foreground elements that appear over the primary\n            background regions, such as grid lines, borders, table banding, icons, and so on.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_background": {
                    "computed": true,
                    "description": "\u003cp\u003eThe background color that applies to the sheet background and sheet controls.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any sheet title, sheet control text, or UI that\n            appears over the secondary background.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success": {
                    "computed": true,
                    "description": "\u003cp\u003eThe color that applies to success messages, for example the check mark for a\n            successful download.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            success color.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning": {
                    "computed": true,
                    "description": "\u003cp\u003eThis color that applies to warning and informational messages.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning_foreground": {
                    "computed": true,
                    "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            warning color.\u003c/p\u003e",
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
        "description": "\u003cp\u003eThe date and time that the theme was created.\u003c/p\u003e",
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
        "description": "\u003cp\u003eThe date and time that the theme was last updated.\u003c/p\u003e",
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
              "computed": true,
              "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n         \u003cul\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon Quick user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon Quick user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
              "description_kind": "plain",
              "optional": true,
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
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
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
        "description": "\u003cp\u003eA version of a theme.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "base_theme_id": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Quick-defined ID of the theme that a custom theme inherits from. All\n            themes initially inherit from a default Quick theme.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe theme configuration. This configuration contains all of the display properties for\n            a theme.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_color_palette": {
                    "computed": true,
                    "description": "\u003cp\u003eThe theme colors that are used for data colors in charts. The colors description is a\n            hexadecimal color code that consists of six alphanumerical characters, prefixed with\n                \u003ccode\u003e#\u003c/code\u003e, for example #37BFF5. \u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "colors": {
                          "computed": true,
                          "description": "\u003cp\u003eThe hexadecimal codes for the colors.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "empty_fill_color": {
                          "computed": true,
                          "description": "\u003cp\u003eThe hexadecimal code of a color that applies to charts where a lack of data is\n            highlighted.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "min_max_gradient": {
                          "computed": true,
                          "description": "\u003cp\u003eThe minimum and maximum hexadecimal codes that describe a color gradient. \u003c/p\u003e",
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
                    "description": "\u003cp\u003eThe theme display options for sheets. \u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "background": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "color": {
                                "computed": true,
                                "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "gradient": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "tile": {
                          "computed": true,
                          "description": "\u003cp\u003eDisplay options related to tiles on a sheet.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "background_color": {
                                "computed": true,
                                "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "border": {
                                "computed": true,
                                "description": "\u003cp\u003eThe display options for tile borders for visuals.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "color": {
                                      "computed": true,
                                      "description": "String to encapsulate the most generic way Color can be formatted (words, hexStrings etc)",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "show": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe option to enable display of borders for visuals.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "width": {
                                      "computed": true,
                                      "description": "String to encapsulate the most generic way Width can be formatted with whatever units (px, em etc)",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "border_radius": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "padding": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "tile_layout": {
                          "computed": true,
                          "description": "\u003cp\u003eThe display options for the layout of tiles on a sheet.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "gutter": {
                                "computed": true,
                                "description": "\u003cp\u003eThe display options for gutter spacing between tiles on a sheet.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "show": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThis Boolean value controls whether to display a gutter space between sheet tiles.\n        \u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "margin": {
                                "computed": true,
                                "description": "\u003cp\u003eThe display options for margins around the outside edge of sheets.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "show": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThis Boolean value controls whether to display sheet margins.\u003c/p\u003e",
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
                        "axis_label_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        "axis_title_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        "data_label_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        },
                        "legend_title_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        "legend_value_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_color": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_decoration": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_family": {
                                "computed": true,
                                "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "absolute": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "font_style": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "font_weight": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                        "visual_subtitle_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_configuration": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "font_color": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_decoration": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_family": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_size": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "absolute": {
                                            "computed": true,
                                            "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "font_style": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_weight": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
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
                              "text_alignment": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "text_transform": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "visual_title_font_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "font_configuration": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "font_color": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_decoration": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_family": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe font family that you want to use.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_size": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "absolute": {
                                            "computed": true,
                                            "description": "\u003cp\u003eThe font size that you want to use in px.\u003c/p\u003e",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "font_style": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "font_weight": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
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
                              "text_alignment": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "text_transform": {
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
                  "ui_color_palette": {
                    "computed": true,
                    "description": "\u003cp\u003eThe theme colors that apply to UI and to charts, excluding data colors. The colors\n            description is a hexadecimal color code that consists of six alphanumerical characters,\n            prefixed with \u003ccode\u003e#\u003c/code\u003e, for example #37BFF5. For more information, see \u003ca href=\"https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html\"\u003eUsing Themes in Amazon Quick\u003c/a\u003e in the \u003ci\u003eAmazon Quick User\n                Guide.\u003c/i\u003e\n         \u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accent": {
                          "computed": true,
                          "description": "\u003cp\u003eThis color is that applies to selected states and buttons.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "accent_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            accent color.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "danger": {
                          "computed": true,
                          "description": "\u003cp\u003eThe color that applies to error messages.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "danger_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            error color.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimension": {
                          "computed": true,
                          "description": "\u003cp\u003eThe color that applies to the names of fields that are identified as\n            dimensions.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimension_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            dimension color.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "measure": {
                          "computed": true,
                          "description": "\u003cp\u003eThe color that applies to the names of fields that are identified as measures.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "measure_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            measure color.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "primary_background": {
                          "computed": true,
                          "description": "\u003cp\u003eThe background color that applies to visuals and other high emphasis UI.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "primary_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe color of text and other foreground elements that appear over the primary\n            background regions, such as grid lines, borders, table banding, icons, and so on.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secondary_background": {
                          "computed": true,
                          "description": "\u003cp\u003eThe background color that applies to the sheet background and sheet controls.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secondary_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any sheet title, sheet control text, or UI that\n            appears over the secondary background.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "success": {
                          "computed": true,
                          "description": "\u003cp\u003eThe color that applies to success messages, for example the check mark for a\n            successful download.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "success_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            success color.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "warning": {
                          "computed": true,
                          "description": "\u003cp\u003eThis color that applies to warning and informational messages.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "warning_foreground": {
                          "computed": true,
                          "description": "\u003cp\u003eThe foreground color that applies to any text or other elements that appear over the\n            warning color.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe date and time that this theme version was created.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "\u003cp\u003eThe description of the theme.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "errors": {
              "computed": true,
              "description": "\u003cp\u003eErrors associated with the theme.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
                    "computed": true,
                    "description": "\u003cp\u003eThe error message.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe version number of the theme.\u003c/p\u003e",
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
