package ucenter

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/mineralres/goshare/pkg/pb"
)

// DBEngine 数据库插件
type DBEngine interface {
}

// RPCHandler RPCHandler
type RPCHandler struct {
}

// MakeRPCHandler MakeRPCHandler
func MakeRPCHandler() *RPCHandler {
	var ret RPCHandler
	return &ret
}

func makeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// UserLogin 用户登陆
func (gs *RPCHandler) UserLogin(ctx context.Context, req *pb.ReqUserLogin) (*pb.RspUserLogin, error) {
	ret := &pb.RspUserLogin{}
	// token := makeMd5(string(fmt.Sprintf("%d", time.Now().Unix())) + time.Now().String())
	ret.Message = "Ok"
	ret.Success = true
	log.Println("UserLogin")
	return ret, nil
}

// UserLogout 用户登出
func (gs *RPCHandler) UserLogout(ctx context.Context, req *pb.EmptyRequest) (*pb.EmptyResponse, error) {
	ret := &pb.EmptyResponse{}
	log.Println("this is UserLogout")
	return ret, nil
}

// CheckAPIPermission CheckAPIPermission
func (gs *RPCHandler) CheckAPIPermission(ctx context.Context, req *pb.ReqCheckAPIPermission) (*pb.RspCheckAPIPermission, error) {
	ret := &pb.RspCheckAPIPermission{}
	ret.Passed = true
	return ret, nil
}

// CheckResourcePermission CheckResourcePermission
func (gs *RPCHandler) CheckResourcePermission(ctx context.Context, req *pb.ReqCheckResourcePermission) (*pb.RspCheckResourcePermission, error) {
	ret := &pb.RspCheckResourcePermission{}
	ret.Passed = true
	return ret, nil
}

// CreateBranch CreateBranch
func (gs *RPCHandler) CreateBranch(ctx context.Context, req *pb.ReqCreateBranch) (*pb.CommonResponse, error) {
	ret := &pb.CommonResponse{}
	return ret, nil
}

// DeleteBranch DeleteBranch
func (gs *RPCHandler) DeleteBranch(ctx context.Context, req *pb.ReqDeleteBranch) (*pb.CommonResponse, error) {
	ret := &pb.CommonResponse{}
	return ret, nil
}

// CreateUser CreateUser
func (gs *RPCHandler) CreateUser(ctx context.Context, req *pb.ReqCreateUser) (*pb.CommonResponse, error) {
	ret := &pb.CommonResponse{}
	return ret, nil
}

// DeleteUser DeleteUser
func (gs *RPCHandler) DeleteUser(ctx context.Context, req *pb.ReqDeleteUser) (*pb.CommonResponse, error) {
	ret := &pb.CommonResponse{}
	return ret, nil
}
