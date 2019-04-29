package user

import "github.com/mineralres/goshare/pkg/pb"

// XUserManager 用户管理系统
type XUserManager interface {
	Login(user, password string) error    // 登陆
	Logout(user string) error             // 登出
	CreateUser() error                    // 新建用户
	DeleteUser(user string) error         // 删除用户
	CreateBranch(branch pb.Branch) error  // 创建部门
	DeleteBranch(id string) error         // 删除部门
	MoveUserToBranch(user, branch string) // 转移用户到新的部门
}
