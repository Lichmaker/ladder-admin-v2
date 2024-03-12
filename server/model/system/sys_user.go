package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userext"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID       `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	Username    string          `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password    string          `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName    string          `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	SideMode    string          `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`     // 用户侧边主题
	HeaderImg   string          `json:"headerImg" gorm:"default:;comment:用户头像"`          // 用户头像
	BaseColor   string          `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`      // 基础颜色
	ActiveColor string          `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"` // 活跃颜色
	AuthorityId uint            `json:"authorityId" gorm:"default:888;comment:用户角色ID"`   // 用户角色ID
	Authority   SysAuthority    `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority  `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string          `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string          `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable      int             `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	UserExt     userext.UserExt `json:"userExt" gorm:"foreignKey:uuid;references:user_uuid;comment:用户信息扩展"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (u SysUser) BuildManagerUserInfo() *manager.UserInfo {
	tmp := manager.UserInfo{}
	tmp.ID = u.UUID.String()
	tmp.AlterID = global.GVA_CONFIG.V2RayManager.DefaultAlterId
	tmp.Email = u.Email
	tmp.Level = global.GVA_CONFIG.V2RayManager.DefaultLevel
	return &tmp
}
