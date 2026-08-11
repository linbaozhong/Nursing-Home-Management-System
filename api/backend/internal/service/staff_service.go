package service

import (
	"context"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblrole"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*staffService)(nil)

type staffService struct{}

// GetRole 获取角色下拉列表
func (s *staffService) GetRole(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	list, _, e := dao.Role(db).List(ctx, ace.Where(tblrole.Id.Gt(types.BigInt(0))).Cols(tblrole.Id, tblrole.Name))
	if e != nil {
		return e
	}
	res := make([]dto.DropDown, 0, len(list))
	for _, r := range list {
		res = append(res, dto.DropDown{ID: int64(r.Id), Name: r.Name.String()})
	}
	*out = res
	return nil
}

// PageStaffByKey 分页查询员工（联角色表）
func (s *staffService) PageStaffByKey(ctx context.Context, in *dto.PageStaffByKeyQuery, out *[]dto.PageStaffByKeyVO) error {
	q := db.Table(tblstaff.TableName).
		LeftJoin(tblstaff.RoleId, tblrole.Id)
	if in.Name != nil {
		q = q.Where(ace.Where(tblstaff.Name.Like(*in.Name)))
	}
	if in.Phone != nil {
		q = q.And(tblstaff.Phone.Like(*in.Phone))
	}
	if in.RoleID != nil {
		q = q.And(tblstaff.RoleId.Eq(types.BigInt(*in.RoleID)))
	}
	return q.Page(in.PageNum, in.PageSize).
		Cols(
			tblstaff.Id.As("id"),
			tblstaff.Name.As("name"),
			tblrole.Name.As("role_name"),
			tblstaff.IdNum.As("id_num"),
			tblstaff.Sex.As("sex"),
			tblstaff.Phone.As("phone"),
			tblstaff.Email.As("email"),
			tblstaff.LeaveFlag.As("leave_flag"),
		).
		OrderBy(tblstaff.CreateTime, false).
		Select().Gets(ctx, out)
}

// PageSearchStaffByKey 分页搜索员工
func (s *staffService) PageSearchStaffByKey(ctx context.Context, in *dto.PageSearchStaffByKeyQuery, out *[]dto.PageSearchStaffByKeyVO) error {
	q := db.Table(tblstaff.TableName)
	if in.Name != nil {
		q = q.Where(ace.Where(tblstaff.Name.Like(*in.Name)))
	}
	if in.Phone != nil {
		q = q.And(tblstaff.Phone.Like(*in.Phone))
	}
	return q.Page(in.PageNum, in.PageSize).
		Cols(
			tblstaff.Id.As("id"),
			tblstaff.Name.As("name"),
			tblstaff.Phone.As("phone"),
		).
		OrderBy(tblstaff.CreateTime, false).
		Select().Gets(ctx, out)
}

// AddStaff 新增员工
func (s *staffService) AddStaff(ctx context.Context, in *dto.OperateStaffQuery) (*dto.EmptyResp, error) {
	has, e := dao.Staff(db).Exists(ctx,
		ace.Where(tblstaff.Phone.Eq(types.String(*in.Phone))).
			Or(tblstaff.Email.Eq(types.String(*in.Email))),
	)
	if e != nil {
		return nil, e
	}
	if has {
		return nil, constant.ErrPhoneOrEmailRepeat
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
	// 初始密码：手机号后 6 位（Java 端为 aesEncode(phone[5:11])，Go 端暂明文存储）
	if ph := *in.Phone; len(ph) >= 11 {
		bean.Pass = types.String(ph[5:11])
	}
	_, e = dao.Staff(db).InsertOne(ctx, bean,
		tblstaff.RoleId, tblstaff.Name, tblstaff.IdNum, tblstaff.Age, tblstaff.Sex,
		tblstaff.Phone, tblstaff.Email, tblstaff.Address, tblstaff.Avator,
		tblstaff.LeaveFlag, tblstaff.Pass, tblstaff.CreateId, tblstaff.CreateTime)
	if e != nil {
		return nil, e
	}
	return new(dto.EmptyResp), nil
}

// GetStaffById 查询员工详情
func (s *staffService) GetStaffById(ctx context.Context, in *dto.IDReq, out *dto.OperateStaffVO) error {
	bean, has, e := dao.Staff(db).GetByID(ctx, types.BigInt(*in.ID),
		tblstaff.Id, tblstaff.RoleId, tblstaff.Name, tblstaff.IdNum, tblstaff.Age,
		tblstaff.Sex, tblstaff.Phone, tblstaff.Email, tblstaff.Address, tblstaff.Avator)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(bean.Id)
	out.RoleID = int64(bean.RoleId)
	out.Name = bean.Name.String()
	out.IDNum = bean.IdNum.String()
	out.Age = int(bean.Age)
	out.Sex = bean.Sex.String()
	out.Phone = bean.Phone.String()
	out.Email = bean.Email.String()
	out.Address = bean.Address.String()
	out.Avator = bean.Avator.String()
	return nil
}

// EditStaff 编辑员工
func (s *staffService) EditStaff(ctx context.Context, in *dto.OperateStaffQuery) (*dto.EmptyResp, error) {
	has, e := dao.Staff(db).Exists(ctx,
		ace.Where(tblstaff.Phone.Eq(types.String(*in.Phone))).
			Or(tblstaff.Email.Eq(types.String(*in.Email))).
			And(tblstaff.Id.NotEq(types.BigInt(*in.ID))),
	)
	if e != nil {
		return nil, e
	}
	if has {
		return nil, constant.ErrPhoneOrEmailRepeat
	}
	_, e = dao.Staff(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblstaff.RoleId.Set(types.BigInt(*in.RoleID)),
		tblstaff.Name.Set(types.String(*in.Name)),
		tblstaff.IdNum.Set(types.String(*in.IDNum)),
		tblstaff.Age.Set(types.Int32(*in.Age)),
		tblstaff.Sex.Set(types.String(*in.Sex)),
		tblstaff.Phone.Set(types.String(*in.Phone)),
		tblstaff.Email.Set(types.String(*in.Email)),
		tblstaff.Address.Set(types.String(*in.Address)),
		tblstaff.Avator.Set(types.String(*in.Avator)),
	)
	if e != nil {
		return nil, e
	}
	return new(dto.EmptyResp), nil
}

// LeaveStaff 员工离职
func (s *staffService) LeaveStaff(ctx context.Context, in *dto.IDReq) (*dto.EmptyResp, error) {
	_, e := dao.Staff(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblstaff.LeaveFlag.Set(types.Int8(constant.YesNoYes)),
	)
	if e != nil {
		return nil, e
	}
	return new(dto.EmptyResp), nil
}
