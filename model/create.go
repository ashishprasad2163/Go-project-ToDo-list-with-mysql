package model

//CreateToDo created
func CreateToDo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}

//DeleteToDo to delete from database
func DeleteToDo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
