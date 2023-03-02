package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"mxshop-api/user-api/proto"
	"mxshop-api/user-api/response"
	"net/http"
	"strconv"
)

func HandlerGrpcError(err error, ctx *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"message": e.Message()})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "内部错误"})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "服务端错误"})
			}
		}
	}
}

func GetUserList(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Error("connect grpc server failed")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	rpcClient := proto.NewUserClient(conn)

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	pageSize := ctx.DefaultQuery("page_size", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	userList, err := rpcClient.GetUserList(context.Background(), &proto.PageInfoRequest{
		Page:     uint32(pageInt),
		PageSize: uint32(pageSizeInt),
	})
	if err != nil {
		zap.S().Error("grpc request failed")
		HandlerGrpcError(err, ctx)
	}
	result := make([]*response.UserResponse, 0)

	for _, value := range userList.Data {
		user := response.UserResponse{
			Id:       value.Id,
			Mobile:   value.Mobile,
			NickName: value.NickName,
			Avatar:   value.Avatar,
			Gender:   value.Gender,
			Role:     value.Role,
			Birthday: response.JsonDateFormat(value.Birthday),
		}
		result = append(result, &user)
	}
	ctx.JSON(http.StatusOK, result)
}
