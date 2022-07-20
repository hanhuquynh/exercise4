package main

import "strconv"

func transaction(birth int64, id string) error {
	session := engine.NewSession()

	session.Engine().ShowSQL()

	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	_, err := session.Where("id = ?", id).Update(&User{
		Birth: birth,
	})

	if err != nil {
		session.Rollback()
		return err
	}

	var point Point

	_, err = session.Where("user_id = ?", id).Get(&point)

	if err != nil {
		session.Rollback()
		return err
	}

	newPoint := point.Points + 10
	newMaxPoint := point.Max_points + 10

	_, err = session.Where("user_id = ?", id).Update(&Point{
		Points:     newPoint,
		Max_points: newMaxPoint,
	})

	if err != nil {
		session.Rollback()
		return err
	}

	var user User
	_, err = session.Where("id = ?", id).Cols("name", "updated_at").Get(&user)

	if err != nil {
		session.Rollback()
		return err
	}

	newName := user.Name + " / " + strconv.Itoa(int(user.Updated_at))

	_, err = session.Where("id = ?", id).Update(&User{
		Name: newName,
	})

	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}
