package entity

// 子菜单
type MenuSvo struct {
	MenuName string `json:"menuName"` // 菜单名称
	Icon     string `json:"icon"`     // 图标
	Url      string `json:"url"`      // url
}

// 左侧菜单
type LeftMenuVo struct {
	Id          int       `json:"id"`          // ID
	MenuName    string    `json:"menuName"`    // 菜单名称
	Icon        string    `json:"icon"`        // 图标
	Url         string    `json:"url"`         // url
	MenuSvoList []MenuSvo `json:"menuSvoList"` // 菜单列表
}

// 用户权限
type PerssionVo struct {
	Value string `json:"value"` // 权限
}
