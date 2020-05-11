package personservice

import (
	"log"
	datacommon "myproject1/datacommon"
	personmodel "myproject1/datamodel/personmodel"
)

func GetAll() []personmodel.PersonResponse {
	db := datacommon.OpenDB()
	defer db.Close()
	rows, err := db.Query("select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var records []personmodel.PersonResponse

	for rows.Next() {
		var rec personmodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, rec)
	}
	return records
}

func GetPerson(model *personmodel.GetPersonRequest) *personmodel.PersonResponse {
	db := datacommon.OpenDB()
	defer db.Close()
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

func AddPerson(model personmodel.AddPersonRequest) bool {
	db := datacommon.OpenDB()
	defer db.Close()
	statement, err := db.Prepare(`insert into personInfo(Name,Phone,MobilePhone,Address,Birthday) 
								  values (?,?,?,?,?)`)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(model.Name.Data, model.Phone.Data, model.MobilePhone.Data, model.Address.Data, model.Birthday)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0
}

func UpdatePerson(model personmodel.UpdatePersonRequest) bool {
	db := datacommon.OpenDB()
	defer db.Close()
	statement, err := db.Prepare(`update personInfo Set Name = ?,
								  Phone = ?,MobilePhone = ?,Address = ?,Birthday = ? 
								  where ID = ?`)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(model.Name.Data, model.Phone.Data, model.MobilePhone.Data,
		model.Address.Data, model.Birthday, model.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowAffected, _ := res.RowsAffected()
	return rowAffected > 0
}

func DeletePerson(model personmodel.DeletePersonRequest) bool {
	db := datacommon.OpenDB()
	defer db.Close()
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
