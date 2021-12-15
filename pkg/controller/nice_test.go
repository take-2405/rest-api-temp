package controller
//
//import (
//	"net/http"
//	"reflect"
//	"testing"
//	"backend-record/pkg/view"
//	"github.com/gin-gonic/gin"
//
//	"backend-record/pkg/model/dto"
//	)
//
//
//func TestUpdateAddLikeHandler(t *testing.T) {
//	tests := []struct {
//		name string
//		want *gin.HandlerFunc
//	}{
//		{
//			name: "TestAddLikeHandker",
//			want: gin.Context.JSON(http.StatusOK, view.ReturnNiceResopnse(dto.Nice{1})),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := UpdateAddLikeHandler(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UpdateAddLikeHandler() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}