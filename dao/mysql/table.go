package mysql

func CheckTable(inter interface{}) error {
	err := db.AutoMigrate(inter)
	return err
}
