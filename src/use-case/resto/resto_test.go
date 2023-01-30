package resto

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fakriardian/Go-kelas.work/src/mocks"
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/fakriardian/Go-kelas.work/src/repository/menu"
	"github.com/fakriardian/Go-kelas.work/src/repository/order"
	"github.com/fakriardian/Go-kelas.work/src/repository/user"
	"github.com/golang/mock/gomock"
)

func Test_restoUseCase_GetMenuList(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}

	type args struct {
		ctx      context.Context
		menuType string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), constant.MenuTypeFood).
						Times(1).
						Return([]model.MenuItem{
							{
								OrderCode: "nasi_uduk",
								Name:      "Nasi Uduk",
								Price:     38000,
								Type:      constant.MenuTypeFood,
							},
						}, nil)
					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: constant.MenuTypeFood,
			},
			want: []model.MenuItem{
				{

					OrderCode: "nasi_uduk",
					Name:      "Nasi Uduk",
					Price:     38000,
					Type:      constant.MenuTypeFood,
				},
			},
			wantErr: false,
		},
		{
			name: "fail get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), constant.MenuTypeFood).
						Times(1).
						Return(nil, errors.New("mock error"))
					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: constant.MenuTypeFood,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restoUseCase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}

			got, err := r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUseCase.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUseCase.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}
