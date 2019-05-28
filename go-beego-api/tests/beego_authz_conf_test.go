package test

import (
	"github.com/angenalZZZ/Go/go-beego-api/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/auth"
	"github.com/casbin/beego-authz/authz"
	"github.com/casbin/casbin"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// 检查访问返回HTTP状态码
func testAuthZRequest(t *testing.T, handler *beego.ControllerRegister, user string, path string, method string, code int) {
	r, _ := http.NewRequest(method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}

// 初始化访问控制的HTTP接口
func testAuthZRequestBasicAuth(username, password string) (e *casbin.Enforcer, handler *beego.ControllerRegister) {
	// 接口
	ap := conf.GetAppPath()
	handler = beego.NewControllerRegister()
	// 接口认证
	_ = handler.InsertFilter("*", beego.BeforeRouter, auth.Basic(username, password))
	// 访问授权
	e = casbin.NewEnforcer(filepath.Join(ap, "conf", "authz_model.conf"), filepath.Join(ap, "conf", "authz_policy.csv"))
	_ = handler.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(e))
	// 访问成功
	handler.Any("*", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})
	return
}

// 测试用户访问1
func TestAuthZ_ACL1(t *testing.T) {
	_, handler := testAuthZRequestBasicAuth("alice", "123")

	testAuthZRequest(t, handler, "alice", "/dataset1/resource1", "GET", 200)
	testAuthZRequest(t, handler, "alice", "/dataset1/resource1", "POST", 200)
	testAuthZRequest(t, handler, "alice", "/dataset1/resource2", "GET", 200)
	testAuthZRequest(t, handler, "alice", "/dataset1/resource2", "POST", 403)
}

// 测试用户访问2
func TestAuthZ_ACL2(t *testing.T) {
	_, handler := testAuthZRequestBasicAuth("bob", "123")

	testAuthZRequest(t, handler, "bob", "/dataset2/resource1", "GET", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/resource1", "POST", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/resource1", "DELETE", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/resource2", "GET", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/resource2", "POST", 403)
	testAuthZRequest(t, handler, "bob", "/dataset2/resource2", "DELETE", 403)

	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item1", "GET", 403)
	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item1", "POST", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item1", "DELETE", 403)
	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item2", "GET", 403)
	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item2", "POST", 200)
	testAuthZRequest(t, handler, "bob", "/dataset2/folder1/item2", "DELETE", 403)
}

// 测试角色访问
func TestAuthZ_RBAC(t *testing.T) {
	e, handler := testAuthZRequestBasicAuth("cathy", "123")

	// cathy can access all /dataset1/* resources via all methods because it has the dataset1_admin role.
	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "GET", 200)
	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "POST", 200)
	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "DELETE", 200)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "GET", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "POST", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "DELETE", 403)

	// delete all roles on user cathy, so cathy cannot access any resources now.
	e.DeleteRolesForUser("cathy")

	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "GET", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "POST", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset1/item", "DELETE", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "GET", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "POST", 403)
	testAuthZRequest(t, handler, "cathy", "/dataset2/item", "DELETE", 403)
}
