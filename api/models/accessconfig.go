package models

import (
	"encoding/json"
)

var AccessConfigList AccessList = nil

const access_config_list = `{

   "ws": {
   	"read": {
   		"actions": [
   			"/api/v1/ws", "/api/v1/ws/*"
   		],
   		"method": "get",
   		"description": "stream access"
   	}
   },
   "node": {
   	"read": {
   		"actions": [
   			"/api/v1/node", "/api/v1/node/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/node"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/node/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/node/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	}
   },
   "device": {
   	"read": {
   		"actions": [
   			"/api/v1/device", "/api/v1/device/[0-9]+",
   			"/api/v1/device/group", "/api/v1/device/[0-9]+/actions",
   			"/api/v1/device/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/device"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/device/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/device/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	}
   },
   "workflow": {
   	"read": {
   		"actions": [
   			"/api/v1/workflow", "/api/v1/workflow/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/workflow"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/workflow/[0-9]+", "/api/v1/workflow/[0-9]+/update_scenario"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/workflow/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	}
   },
   "flow": {
   	"read": {
   		"actions": [
   			"/api/v1/flow", "/api/v1/flow/[0-9]+", "/api/v1/flow/[0-9]+/flow",
   			"/api/v1/flow/[0-9]+/redactor", "/api/v1/flow/[0-9]+/workers",
   			"/api/v1/flow/[0-9]+/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/flow"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/flow/[0-9]+", "/api/v1/flow/[0-9]+/redactor"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/flow/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	}
   },
   "device_action": {
   	"read": {
   		"actions": [
   			"/api/v1/device_action", "/api/v1/device_action/[0-9]+",
   			"/api/v1/device_action/search", "/api/v1/device_action/get_by_device/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/device_action"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/device_action/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/device_action/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "worker": {
   	"read": {
   		"actions": [
   			"/api/v1/worker", "/api/v1/worker/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/worker"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/worker/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/worker/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "script": {
   	"read": {
   		"actions": [
   			"/api/v1/script", "/api/v1/script/[0-9]+",
   			"/api/v1/script/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/script"
   		],
   		"method": "post",
   		"description": ""
   	},
	"exec_script": {
   		"actions": [
   			"/api/v1/script/[0-9]+/exec"
   		],
   		"method": "post",
   		"description": "execute script"
   	},
   	"update": {
   		"actions": [
   			"/api/v1/script/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/script/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "log": {
   	"read": {
   		"actions": [
   			"/api/v1/log", "/api/v1/log/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	}
   },
   "notifr": {
   	"read_notifr_template": {
   		"actions": [
   			"/api/v1/email/template/[\\w]+", "/api/v1/email/template", "/api/v1/email/template/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_notifr_template": {
   		"actions": [
   			"/api/v1/email/template"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_notifr_template": {
   		"actions": [
   			"/api/v1/email/template/[\\w]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_notifr_template": {
   		"actions": [
   			"/api/v1/email/template/[\\w]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"preview_notifr": {
   		"actions": [
   			"/api/v1/email/preview"
   		],
   		"method": "post",
   		"description": ""
	},
	"read_notifr_item": {
   		"actions": [
   			"/api/v1/email/item/[\\w]+", "/api/v1/email/items",
   			"/api/v1/email/items/tree"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_notifr_item": {
   		"actions": [
   			"/api/v1/email/item", "/api/v1/email/items/tree"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_notifr_item": {
   		"actions": [
   			"/api/v1/email/item/[\\w]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_notifr_item": {
   		"actions": [
   			"/api/v1/email/item/[\\w]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"show_notify": {
		"actions": [
			"/api/v1/notifr/[0-9]+", "/api/v1/notifr"
		],
		"method": "get",
   		"description": ""
	},
	"create_notify": {
		"actions": [
			"/api/v1/notifr"
		],
		"method": "post",
   		"description": ""
	},
	"repeat_notify": {
		"actions": [
			"/api/v1/notifr/[0-9]+/repeat"
		],
		"method": "post",
   		"description": ""
	},
	"delete_notify": {
		"actions": [
			"/api/v1/notifr/[0-9]+"
		],
		"method": "delete",
   		"description": ""
	}
   },
   "map": {
   	"read_map": {
   		"actions": [
   			"/api/v1/map", "/api/v1/map/[0-9]+",
   			"/api/v1/map/[0-9]+/full"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_map": {
   		"actions": [
   			"/api/v1/map"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_map": {
   		"actions": [
   			"/api/v1/map/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_map": {
   		"actions": [
   			"/api/v1/map/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"read_map_layer": {
   		"actions": [
   			"/api/v1/map_layer", "/api/v1/map_layer/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_map_layer": {
   		"actions": [
   			"/api/v1/map_layer"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_map_layer": {
   		"actions": [
   			"/api/v1/map_layer/[0-9]+",
   			"/api/v1/map_layer/sort"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_map_layer": {
   		"actions": [
   			"/api/v1/map_layer/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"read_map_element": {
   		"actions": [
   			"/api/v1/map_element", "/api/v1/map_element/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_map_element": {
   		"actions": [
   			"/api/v1/map_element"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_map_element": {
   		"actions": [
   			"/api/v1/map_element/[0-9]+",
   			"/api/v1/map_element/[0-9]+/element_only",
   			"/api/v1/map_element/Sort"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_map_element": {
   		"actions": [
   			"/api/v1/map_element/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "device_state": {
   	"read": {
   		"actions": [
   			"/api/v1/device_state",
   			"/api/v1/device_state/[0-9]+",
   			"/api/v1/device_state/get_by_device/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/device_state"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/device_state/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/device_state/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "image": {
   	"read": {
   		"actions": [
   			"/api/v1/image", "/api/v1/image/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/image"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"upload": {
   		"actions": [
   			"/api/v1/image/upload"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/image/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/image/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "dashboard": {
   	"read": {
   		"actions": [
   			"/api/v1/dashboard", "/api/v1/dashboard/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/dashboard"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/dashboard/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/dashboard/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	}
   },
   "user": {
   	"read": {
   		"actions": [
   			"/api/v1/user",
   			"/api/v1/user/[0-9]+",
   			"/api/v1/user/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/user"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update": {
   		"actions": [
   			"/api/v1/user/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"update_status": {
   		"actions": [
   			"/api/v1/user/[0-9]+/update_status"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			"/api/v1/user/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"read_role": {
   		"actions": [
   			"/api/v1/role",
   			"/api/v1/role/[\\w]+",
   			"/api/v1/role/search"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_role": {
   		"actions": [
   			"/api/v1/role",
   			"/api/v1/role/[\\w]+"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_role": {
   		"actions": [
   			"/api/v1/user/[\\w]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_role": {
   		"actions": [
   			"/api/v1/role/[\\w]+"
   		],
   		"method": "delete",
   		"description": ""
	},
	"read_role_access_list": {
   		"actions": [
   			"/api/v1/role/[\\w]+/access_list",
   			"/api/v1/access_list"
   		],
   		"method": "get",
   		"description": "view role access list"
	},
	"update_role_access_list": {
   		"actions": [
   			"/api/v1/user/[\\w]+/access_list"
   		],
   		"method": "put",
   		"description": "update role access list info"
   	}
   },
   "scenarios": {
   	"read": {
   		"actions": [
   			"/api/v1/scenario", "/api/v1/scenario/[0-9]+"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create": {
   		"actions": [
   			"/api/v1/scenario"
   		],
   		"method": "post",
   		"description": ""
   	},
	"update": {
   		"actions": [
   			 "/api/v1/scenario/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete": {
   		"actions": [
   			 "/api/v1/scenario/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	},
   	"read_script": {
   		"actions": [
   			 "/api/v1/scenario_script/[0-9]+", "/api/v1/scenario_script"
   		],
   		"method": "get",
   		"description": ""
   	},
   	"create_script": {
   		"actions": [
   			 "/api/v1/scenario_script"
   		],
   		"method": "post",
   		"description": ""
   	},
   	"update_script": {
   		"actions": [
   			  "/api/v1/scenario_script/[0-9]+"
   		],
   		"method": "put",
   		"description": ""
   	},
   	"delete_script": {
   		"actions": [
   			  "/api/v1/scenario_script/[0-9]+"
   		],
   		"method": "delete",
   		"description": ""
   	}
	}
}
`

func init() {
	AccessConfigList = AccessList{}
	if err := json.Unmarshal([]byte(access_config_list), &AccessConfigList); err != nil {
		panic(err.Error())
	}

}