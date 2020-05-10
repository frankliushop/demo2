package dataservice

import (
	datacommon "myproject1/datacommon"
	datamodel "myproject1/datamodel"
	"log"
)

type PersonService struct {
}

func GetPersonServiceInstance() *PersonService {
	inst := &PersonService{}
	return inst
}

func (personService *PersonService) GetAll() []datamodel.PersonResponse {
	db := datacommon.OpenDB()
	defer db.Close()
	rows, err := db.Query("select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var records []datamodel.PersonResponse

	for rows.Next() {
		var rec datamodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, rec)
	}
	return records
}

func (personService *PersonService) GetPerson(model *datamodel.GetPersonRequest) *datamodel.PersonResponse {
	db := datacommon.OpenDB()
	defer db.Close()
	rows, err := db.Query(`select ID,Name,Phone,MobilePhone,Address,Birthday from personInfo
						   where ID = ?`, model.ID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var rec datamodel.PersonResponse
		err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone, &rec.MobilePhone, &rec.Address, &rec.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		return &rec
	}
	return nil
}

func (personService *PersonService) AddPerson(model datamodel.AddPersonRequest) bool {
	db := datacommon.OpenDB()
	defer db.Close()
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

func (personService *PersonService) UpdatePerson(model datamodel.UpdatePersonRequest) bool {
	db := datacommon.OpenDB()
	defer db.Close()
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
	return rowAffected > 0
}

func (personService *PersonService) DeletePerson(model datamodel.DeletePersonRequest) bool {
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
