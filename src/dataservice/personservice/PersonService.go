package personservice

import (
	"log"
	datacommon "myproject1/datacommon"
	datamodel "myproject1/datamodel"
	personmodel "myproject1/datamodel/personmodel"
)

//從資料庫取得所有資料
func GetAll(model *personmodel.GetAllRequest) *datamodel.PagingResponse {
	count := getPersonCount()
	if count == 0 {
		return &datamodel.PagingResponse{
			PageIndex: model.PageIndex.Int64,
			PageSize:  model.PageSize.Int64,
			PageCount: 0,
			ItemCount: 0,
			DataList:  nil,
		}
	}
	pagingResponse := getPersonList(model)
	pageCount := datacommon.GetPageCount(model.PageSize.Int64, count)
	pagingResponse.PageCount = pageCount
	pagingResponse.ItemCount = count
	return pagingResponse
}

func getPersonList(model *personmodel.GetAllRequest) *datamodel.PagingResponse {
	db := datacommon.DbCon
	skip := (model.PageIndex.Int64 - 1) * model.PageSize.Int64
	rows, err := db.Query(`select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo limit ?,?`,
		skip, model.PageSize)

	var records []personmodel.PersonResponse = make([]personmodel.PersonResponse, 0,
		int(model.PageSize.Int64))

	for rows.Next() {
		var rec personmodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, rec)
	}

	pagingResponse := &datamodel.PagingResponse{
		PageIndex: model.PageIndex.Int64,
		PageSize:  model.PageSize.Int64,
		DataList:  records,
	}

	return pagingResponse
}

func getPersonCount() int64 {
	db := datacommon.DbCon
	rows, err := db.Query(`select count(1) from personInfo`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var count int64 = 0
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}
	return count
}

//從資料庫取得單筆資料
func GetPerson(model *personmodel.GetPersonRequest) *personmodel.PersonResponse {
	db := datacommon.DbCon
	rows, err := db.Query(`select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo
						   where ID = ?`, model.ID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var rec personmodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		return &rec
	}
	return nil
}

//新增資料庫資料
func AddPerson(model personmodel.AddPersonRequest) bool {
	db := datacommon.DbCon
	statement, err := db.Prepare(`insert into personInfo(Name,Phone,MobilePhone,Address,Birthday) 
								  values (?,?,?,?,?)`)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(model.Name, model.Phone, model.MobilePhone, model.Address, model.Birthday)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0
}

//更新資料庫資料
func UpdatePerson(model personmodel.UpdatePersonRequest) bool {
	db := datacommon.DbCon
	statement, err := db.Prepare(`update personInfo Set Name = ?,
								  Phone = ?,MobilePhone = ?,Address = ?,Birthday = ? 
								  where ID = ?`)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(model.Name, model.Phone, model.MobilePhone,
		model.Address, model.Birthday, model.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected >= 0
}

//刪除資料庫資料
func DeletePerson(model personmodel.DeletePersonRequest) bool {
	db := datacommon.DbCon
	statement, err := db.Prepare(`delete from personInfo where ID = ?`)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(model.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0
}
