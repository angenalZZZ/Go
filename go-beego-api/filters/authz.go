package filters

// 权限访问控制
func AddPluginOfAuthZ() {
	//ap := conf.GetAppPath()
	//model, policy := filepath.Join(ap, "conf", "authz_model.conf"), filepath.Join(ap, "conf", "authz_policy.csv")

	// Simple Usage
	//beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(casbin.NewEnforcer(model, policy)))

	// Advanced Usage
	//e := casbin.NewEnforcer("authz_model.conf", "")
	//e.AddRoleForUser("alice", "admin")
	//e.AddPolicy(...)
	//beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(e))
}
