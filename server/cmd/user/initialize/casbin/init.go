package casbin

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func InitCasbin() {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.CasbinMysqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	a, _ := gormadapter.NewAdapter("mysql", dsn, true) // Your driver and data source.
	e, _ := casbin.NewEnforcer(consts.CasbinModelPath, a)
	e.AddPolicy("seller", "create", "create")
	config.Enforcer = e
}
