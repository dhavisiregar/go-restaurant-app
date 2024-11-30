package resto

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dhavisiregar/go-restaurant-app/internal/model"
	"github.com/dhavisiregar/go-restaurant-app/internal/model/constant"
	"github.com/dhavisiregar/go-restaurant-app/internal/repository/menu"
	"github.com/dhavisiregar/go-restaurant-app/internal/repository/order"
	"github.com/dhavisiregar/go-restaurant-app/internal/repository/user"
	"github.com/golang/mock/gomock"
)

func Test_restoUsecase_Order(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	type args struct {
		ctx     context.Context
		menuType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		{
		name: "success get menu list",
		fields: fields{
			menuRepo: func () menu.Repository {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockMenuRepository(ctrl)
				
				mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
					Times(1).
					Return([]model.MenuItem{
							{
                            Name:      "Spaghetti Carbonara",
                            OrderCode: "spaghetti_carbonara",
                            Price:      50000,
                            Type:      constant.MenuTypeFood,
                        	},
						}, nil)

					return mock
				}(),
			},
			args: args{
				ctx:     context.Background(),
				menuType: string(constant.MenuTypeFood),		
			},
			want: []model.MenuItem{
				{
                    Name:      "Spaghetti Carbonara",
                    OrderCode: "spaghetti_carbonara",
                    Price:      50000,
                    Type:      constant.MenuTypeFood,
                },
			},
			wantErr: false,
		},
		{
			name: "fail get menu list",
			fields: fields{
				menuRepo: func () menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)
					
					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return(nil, errors.New("mock error"))
	
					return mock
				}(),
			},
			args: args{
				ctx:     context.Background(),
				menuType: string(constant.MenuTypeFood),		
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restoUsecase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}
			got, err := r.Order(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.Order() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
