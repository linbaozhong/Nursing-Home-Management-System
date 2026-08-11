package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

type user struct{}

var User = &user{}

// Register 用户注册（Go 端用户即员工 staff）
func (u *user) Register(ctx context.Context, in *dto.OperateStaffQuery, out *dto.OperateStaffVO) error {
	has, e := dao.Staff(db).Exists(ctx,
		ace.Where(tblstaff.Phone.Eq(types.String(*in.Phone))).
			Or(tblstaff.Email.Eq(types.String(*in.Email))),
	)
	if e != nil {
		return e
	}
	if has {
		return constant.ErrPhoneOrEmailRepeat
	}
	bean := new(do.Staff)
	bean.RoleId = types.BigInt(*in.RoleID)
	bean.Name = types.String(*in.Name)
	bean.IdNum = types.String(*in.IDNum)
	bean.Age = types.Int32(*in.Age)
	bean.Sex = types.String(*in.Sex)
	bean.Phone = types.String(*in.Phone)
	bean.Email = types.String(*in.Email)
	bean.Address = types.String(*in.Address)
	bean.Avator = types.String(*in.Avator)
	bean.LeaveFlag = types.Int8(constant.YesNoNo)
	if ph := *in.Phone; len(ph) >= 11 {
		bean.Pass = types.String(ph[5:11])
	}
	ok, e := dao.Staff(db).InsertOne(ctx, bean,
		tblstaff.RoleId, tblstaff.Name, tblstaff.IdNum, tblstaff.Age, tblstaff.Sex,
		tblstaff.Phone, tblstaff.Email, tblstaff.Address, tblstaff.Avator,
		tblstaff.LeaveFlag, tblstaff.Pass, tblstaff.CreateId, tblstaff.CreateTime)
	if e != nil {
		return e
	}
	if !ok {
		return errors.New("register failed")
	}
	out.ID = int64(bean.Id)
	out.RoleID = int64(bean.RoleId)
	out.Name = bean.Name.String()
	return nil
}

// Get 根据用户编号查询（Go 端用户即员工 staff）
func (u *user) Get(ctx context.Context, in *dto.IDReq, out *dto.OperateStaffVO) error {
	bean, has, e := dao.Staff(db).GetByID(ctx, types.BigInt(*in.ID),
		tblstaff.Id, tblstaff.RoleId, tblstaff.Name, tblstaff.IdNum, tblstaff.Age,
		tblstaff.Sex, tblstaff.Phone, tblstaff.Email, tblstaff.Address, tblstaff.Avator)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("user not found")
	}
	out.ID = int64(bean.Id)
	out.RoleID = int64(bean.RoleId)
	out.Name = bean.Name.String()
	out.Email = bean.Email.String()
	return nil
}
