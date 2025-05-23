package karabinerconfig

// Default location is ${HOME}/.config/karabiner/karabiner.json
const Default = `
  {
    "complex_modifications": {
        "parameters": {
            "basic.simultaneous_threshold_milliseconds": 50,
            "basic.to_delayed_action_delay_milliseconds": 500,
            "basic.to_if_alone_timeout_milliseconds": 1000,
            "basic.to_if_held_down_threshold_milliseconds": 500,
            "mouse_motion_to_scroll.speed": 100
        },
        "rules": [
            {
                "description": "Change right_command+hjkl to arrow keys",
                "manipulators": [
                    {
                        "from": {
                            "key_code": "h",
                            "modifiers": {
                                "mandatory": [
                                    "right_command"
                                ],
                                "optional": [
                                    "any"
                                ]
                            }
                        },
                        "to": [
                            {
                                "key_code": "left_arrow"
                            }
                        ],
                        "type": "basic"
                    },
                    {
                        "from": {
                            "key_code": "j",
                            "modifiers": {
                                "mandatory": [
                                    "right_command"
                                ],
                                "optional": [
                                    "any"
                                ]
                            }
                        },
                        "to": [
                            {
                                "key_code": "down_arrow"
                            }
                        ],
                        "type": "basic"
                    },
                    {
                        "from": {
                            "key_code": "k",
                            "modifiers": {
                                "mandatory": [
                                    "right_command"
                                ],
                                "optional": [
                                    "any"
                                ]
                            }
                        },
                        "to": [
                            {
                                "key_code": "up_arrow"
                            }
                        ],
                        "type": "basic"
                    },
                    {
                        "from": {
                            "key_code": "l",
                            "modifiers": {
                                "mandatory": [
                                    "right_command"
                                ],
                                "optional": [
                                    "any"
                                ]
                            }
                        },
                        "to": [
                            {
                                "key_code": "right_arrow"
                            }
                        ],
                        "type": "basic"
                    }
                ]
            },
            {
                "description": "",
                "manipulators": [
                    {
                        "from": {
                            "key_code": "caps_lock",
                            "modifiers": {
                                "mandatory": null,
                                "optional": [
                                    "any"
                                ]
                            }
                        },
                        "to": [
                            {
                                "key_code": "left_shift"
                            }
                        ],
                        "type": "basic"
                    }
                ]
            }
        ]
    },
    "devices": [],
    "fn_function_keys": [
        {
            "from": {
                "key_code": "f1"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "display_brightness_decrement",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f2"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "display_brightness_increment",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f3"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "mission_control",
                    "consumer_key_code": "",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f4"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "spotlight",
                    "consumer_key_code": "",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f5"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "dictation",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f6"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "",
                    "key_code": "f6"
                }
            ]
        },
        {
            "from": {
                "key_code": "f7"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "rewind",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f8"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "play_or_pause",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f9"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "fast_forward",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f10"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "mute",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f11"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "volume_decrement",
                    "key_code": ""
                }
            ]
        },
        {
            "from": {
                "key_code": "f12"
            },
            "to": [
                {
                    "apple_vendor_keyboard_key_code": "",
                    "consumer_key_code": "volume_increment",
                    "key_code": ""
                }
            ]
        }
    ],
    "global": {
        "check_for_updates_on_startup": true,
        "show_in_menu_bar": true,
        "show_profile_name_in_menu_bar": false
    },
    "name": "Default profile",
    "parameters": {
        "delay_milliseconds_before_open_device": 1000
    },
    "profiles": [
        {
            "complex_modifications": {
                "parameters": {
                    "basic.simultaneous_threshold_milliseconds": 50,
                    "basic.to_delayed_action_delay_milliseconds": 500,
                    "basic.to_if_alone_timeout_milliseconds": 1000,
                    "basic.to_if_held_down_threshold_milliseconds": 500,
                    "mouse_motion_to_scroll.speed": 100
                },
                "rules": [
                    {
                        "manipulators": [
                            {
                                "description": "Change caps_lock to command+control+option+shift.",
                                "from": {
                                    "key_code": "caps_lock",
                                    "modifiers": {
                                        "optional": [
                                            "any"
                                        ]
                                    }
                                },
                                "to": [
                                    {
                                        "key_code": "left_shift",
                                        "modifiers": [
                                            "left_command",
                                            "left_control",
                                            "left_option"
                                        ]
                                    }
                                ],
                                "type": "basic"
                            }
                        ]
                    },
                    {
                        "description": "Change right_command+hjkl to arrow keys",
                        "manipulators": [
                            {
                                "from": {
                                    "key_code": "h",
                                    "modifiers": {
                                        "mandatory": [
                                            "right_command"
                                        ],
                                        "optional": [
                                            "any"
                                        ]
                                    }
                                },
                                "to": [
                                    {
                                        "key_code": "left_arrow"
                                    }
                                ],
                                "type": "basic"
                            },
                            {
                                "from": {
                                    "key_code": "j",
                                    "modifiers": {
                                        "mandatory": [
                                            "right_command"
                                        ],
                                        "optional": [
                                            "any"
                                        ]
                                    }
                                },
                                "to": [
                                    {
                                        "key_code": "down_arrow"
                                    }
                                ],
                                "type": "basic"
                            },
                            {
                                "from": {
                                    "key_code": "k",
                                    "modifiers": {
                                        "mandatory": [
                                            "right_command"
                                        ],
                                        "optional": [
                                            "any"
                                        ]
                                    }
                                },
                                "to": [
                                    {
                                        "key_code": "up_arrow"
                                    }
                                ],
                                "type": "basic"
                            },
                            {
                                "from": {
                                    "key_code": "l",
                                    "modifiers": {
                                        "mandatory": [
                                            "right_command"
                                        ],
                                        "optional": [
                                            "any"
                                        ]
                                    }
                                },
                                "to": [
                                    {
                                        "key_code": "right_arrow"
                                    }
                                ],
                                "type": "basic"
                            }
                        ]
                    }
                ]
            },
            "devices": [],
            "fn_function_keys": [
                {
                    "from": {
                        "key_code": "f1"
                    },
                    "to": [
                        {
                            "consumer_key_code": "display_brightness_decrement"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f2"
                    },
                    "to": [
                        {
                            "consumer_key_code": "display_brightness_increment"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f3"
                    },
                    "to": [
                        {
                            "apple_vendor_keyboard_key_code": "mission_control"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f4"
                    },
                    "to": [
                        {
                            "apple_vendor_keyboard_key_code": "spotlight"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f5"
                    },
                    "to": [
                        {
                            "consumer_key_code": "dictation"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f6"
                    },
                    "to": [
                        {
                            "key_code": "f6"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f7"
                    },
                    "to": [
                        {
                            "consumer_key_code": "rewind"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f8"
                    },
                    "to": [
                        {
                            "consumer_key_code": "play_or_pause"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f9"
                    },
                    "to": [
                        {
                            "consumer_key_code": "fast_forward"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f10"
                    },
                    "to": [
                        {
                            "consumer_key_code": "mute"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f11"
                    },
                    "to": [
                        {
                            "consumer_key_code": "volume_decrement"
                        }
                    ]
                },
                {
                    "from": {
                        "key_code": "f12"
                    },
                    "to": [
                        {
                            "consumer_key_code": "volume_increment"
                        }
                    ]
                }
            ],
            "name": "Default profile",
            "parameters": {
                "delay_milliseconds_before_open_device": 1000
            },
            "selected": true,
            "simple_modifications": [],
            "virtual_hid_keyboard": {
                "country_code": 0,
                "indicate_sticky_modifier_keys_state": true,
                "mouse_key_xy_scale": 100
            }
        }
    ],
    "selected": true,
    "simple_modifications": [],
    "virtual_hid_keyboard": {
        "country_code": 0,
        "indicate_sticky_modifier_keys_state": true,
        "mouse_key_xy_scale": 100
    }
}
`
