package studioService

import (
	"FinalProject/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type StudioImpl struct {
	DB *sql.DB
}

func NewStudioService(DB *sql.DB) StudioService {
	return &StudioImpl {
		DB: DB,
	}
}


func (s *StudioImpl) GetStudio() (*[]models.Studio, error) {
	var studios = []models.Studio{}

	sql 		:= `SELECT * FROM studios`
	rows, err 	:= s.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var studio = models.Studio{}

		err = rows.Scan(&studio.Id, &studio.Name, &studio.FilmId, &studio.CreatedAt, &studio.UpdatedAt,)
		if err != nil {
			return nil, err
		}

		studios = append(studios, studio)
	}

	return &studios, nil
}

func (s *StudioImpl) AddStudio(studio *models.AddStudio) (*models.Studio, error) {
	var newStudio = models.Studio{}

	if studio.Name == "" {
		return nil, errors.New("Name Must Be Filled")
	}else if strconv.Itoa(studio.FilmId) == "" {
		return nil, errors.New("Film Id Must Be Filled")
	}

	sql := `INSERT INTO studios (name, film_id) VALUES ($1, $2) Returning *`
	err := s.DB.QueryRow(sql, studio.Name, studio.FilmId,).Scan(
		&newStudio.Id, &newStudio.Name, &newStudio.FilmId, &newStudio.CreatedAt, &newStudio.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newStudio, nil
}

func (s *StudioImpl) GetStudioById(id int) (*models.Studio, error) {
	var studio = models.Studio{}
	sql := `SELECT * FROM studios WHERE id=($1)`
	err := s.DB.QueryRow(sql, id).Scan(&studio.Id, &studio.Name, &studio.FilmId, &studio.CreatedAt, &studio.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &studio, err
}

func (s *StudioImpl) UpdateStudio(id int, studio *models.Studio) (int, error) {
	if studio.Name == "" {
		return 0, errors.New("Name Must Be Filled")
	}else if strconv.Itoa(studio.FilmId) == "" {
		return 0, errors.New("Film Id Must Be Filled")
	}

	sqlStatement := `UPDATE studios SET name=$2, film_id=$3 WHERE id=$1;`
	
	result, err := s.DB.Exec(sqlStatement, id, studio.Name,  studio.FilmId,)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating studio record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the studio, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (s *StudioImpl) DeleteStudio(id int) (int, error) {
	sql := `DELETE FROM studios WHERE id=$1;`
	res, err := s.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete studio record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the studio, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


