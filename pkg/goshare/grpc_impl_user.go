package goshare

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"encoding/hex"

	"github.com/golang/protobuf/jsonpb"
	"github.com/mineralres/goshare/pkg/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func makeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// UserLogin 用户登陆
func (gs *GrpcServer) UserLogin(ctx context.Context, req *pb.ReqUserLogin) (*pb.RspUserLogin, error) {
	var ret pb.RspUserLogin
	sessionID, _ := gs.cache.XDB().GetUniqueID()
	token := makeMd5(string(sessionID) + time.Now().String())
	header := metadata.Pairs("Grpc-Metadata-set-cookie", fmt.Sprintf("token=%s;Max-Age=886400;", token))
	grpc.SendHeader(ctx, header)
	ret.Message = "Ok"
	ret.Success = true
	return &ret, nil
}

// UserLogout 用户登出
func (gs *GrpcServer) UserLogout(ctx context.Context, req *pb.EmptyRequest) (*pb.EmptyResponse, error) {
	log.Println("this is UserLogout")
	var ret pb.EmptyResponse
	return &ret, nil
}

// Routes 路由表
func (gs *GrpcServer) Routes(ctx context.Context, req *pb.EmptyRequest) (*pb.RspGetRoutes, error) {
	session := getUserSession(ctx)
	if session == nil {
		return &pb.RspGetRoutes{}, ErrorNeedLogin
	}
	var ret pb.RspGetRoutes
	data, err := ioutil.ReadFile("routes.json")
	if err != nil {
		log.Println(err)
		return &ret, err
	}
	err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(bytes.NewReader(data), &ret)
	return &ret, err
}

// CurrentUser 当前用户
func (gs *GrpcServer) CurrentUser(ctx context.Context, req *pb.EmptyRequest) (*pb.RspCurrentUser, error) {
	var ret pb.RspCurrentUser
	ret.Success = true
	ret.User = &pb.XUser{Username: "admin"}
	ret.User.Avatar = "https://pbs.twimg.com/profile_images/835224725815246848/jdMBCxHS.jpg"
	ret.User.Permissions = &pb.XPermission{Role: "admin"}
	return &ret, nil
}

// Users Users
func (gs *GrpcServer) Users(ctx context.Context, req *pb.EmptyRequest) (*pb.XUsersItemList, error) {
	var ret pb.XUsersItemList
	return &ret, nil
}
