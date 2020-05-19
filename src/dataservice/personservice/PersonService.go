package personservice

import (
	datacommon "myproject1/datacommon"
	datamodel "myproject1/datamodel"
	personmodel "myproject1/datamodel/personmodel"
)

//從資料庫取得所有資料
func GetAll(model *personmodel.GetAllRequest) (*datamodel.PagingResponse, error) {
	count,err := getPersonCount()
	if err != nil {
		return nil,err
	}
	if count == 0 {
		return &datamodel.PagingResponse{
			PageIndex: model.PageIndex.Int64,
			PageSize:  model.PageSize.Int64,
			PageCount: 0,
			ItemCount: 0,
			DataList:  nil,
		},nil
	}
	pagingResponse,err := getPersonList(model)
	if err != nil {
		return nil,err
	}
	pageCount := datacommon.GetPageCount(model.PageSize.Int64, count)
	pagingResponse.PageCount = pageCount
	pagingResponse.ItemCount = count
	return pagingResponse, nil
}

func getPersonList(model *personmodel.GetAllRequest) (*datamodel.PagingResponse, error) {
	db := datacommon.DbCon
	skip := (model.PageIndex.Int64 - 1) * model.PageSize.Int64
	rows, err := db.Query(`select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo limit ?,?`,
		skip, model.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []personmodel.PersonResponse = make([]personmodel.PersonResponse, 0,
		int(model.PageSize.Int64))

	for rows.Next() {
		var rec personmodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			return nil, err
		}
		records = append(records, rec)
	}

	pagingResponse := &datamodel.PagingResponse{
		PageIndex: model.PageIndex.Int64,
		PageSize:  model.PageSize.Int64,
		DataList:  records,
	}

	return pagingResponse, nil
}

func getPersonCount() (int64, error) {
	db := datacommon.DbCon
	rows, err := db.Query(`select count(1) from personInfo`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int64 = 0
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

//從資料庫取得單筆資料
func GetPerson(model *personmodel.GetPersonRequest) (*personmodel.PersonResponse, error) {
	db := datacommon.DbCon
	rows, err := db.Query(`select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo
						   where ID = ?`, model.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec personmodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			return nil, err
		}
		return &rec, nil
	}
	return nil, nil
}

//新增資料庫資料
func AddPerson(model personmodel.AddPersonRequest) (bool, error) {
	db := datacommon.DbCon
	statement, err := db.Prepare(`insert into personInfo(Name,Phone,MobilePhone,Address,Birthday) 
								  values (?,?,?,?,?)`)
	defer statement.Close()
	if err != nil {
		return false, err
	}
	res, err := statement.Exec(model.Name, model.Phone, model.MobilePhone, model.Address, model.Birthday)
	if err != nil {
		return false, err
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0, nil
}

//更新資料庫資料
func UpdatePerson(model personmodel.UpdatePersonRequest) (bool, error) {
	db := datacommon.DbCon
	statement, err := db.Prepare(`update personInfo Set Name = ?,
								  Phone = ?,MobilePhone = ?,Address = ?,Birthday = ? 
								  where ID = ?`)
	defer statement.Close()
	if err != nil {
		return false, err
	}
	res, err := statement.Exec(model.Name, model.Phone, model.MobilePhone,
		model.Address, model.Birthday, model.ID)
	if err != nil {
		return false, err
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected >= 0, nil
}

//刪除資料庫資料
func DeletePerson(model personmodel.DeletePersonRequest) (bool, error) {
	db := datacommon.DbCon
	statement, err := db.Prepare(`delete from personInfo where ID = ?`)
	defer statement.Close()
	if err != nil {
		return false, nil
	}
	res, err := statement.Exec(model.ID)
	if err != nil {
		return false, nil
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0, nil
}
