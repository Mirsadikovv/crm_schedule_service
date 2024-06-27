package postgres

import (
	"context"
	"database/sql"
	"fmt"
	tc "go_schedule_service/genproto/student_perfomance_service"
	"go_schedule_service/pkg"
	"go_schedule_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type perfomanceRepo struct {
	db *pgxpool.Pool
}

func NewPerfomanceRepo(db *pgxpool.Pool) storage.PerfomanceRepoI {
	return &perfomanceRepo{
		db: db,
	}
}

func (c *perfomanceRepo) Create(ctx context.Context, req *tc.CreatePerfomance) (*tc.GetPerfomance, error) {

	id := uuid.NewString()

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO perfomances (
			id,
			student_id,
			schedule_id,
			attended,
			task_score
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
		)`,
		id,
		req.BranchI,
		userLogin,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		pasword,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.StartWorking,
		end_working)
	if err != nil {
		log.Println("error while creating perfomance", comtag)
		return nil, err
	}

	perfomance, err := c.GetById(ctx, &tc.PerfomancePrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting perfomance by id")
		return nil, err
	}
	return perfomance, nil
}

func (c *perfomanceRepo) Update(ctx context.Context, req *tc.UpdatePerfomance) (*tc.GetPerfomance, error) {
	var end_working sql.NullString
	if req.EndWorking == "" {
		end_working = sql.NullString{Valid: false}
	} else {
		end_working = sql.NullString{String: req.EndWorking, Valid: true}
	}
	_, err := c.db.Exec(ctx, `
		UPDATE perfomances SET
		branch_id = $1,
		birthday = $2,
		gender = $3,
		fullname = $4,
		email = $5,
		phone = $6,
		salary = $7,
		ielts_score = $8,
		ielts_attempts_count = $9,
		start_working = $10,
		end_working = $11,
		updated_at = NOW()
		WHERE id = $12
		`,
		req.BranchId,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.StartWorking,
		end_working,
		req.Id)
	if err != nil {
		log.Println("error while updating perfomance")
		return nil, err
	}

	perfomance, err := c.GetById(ctx, &tc.PerfomancePrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting perfomance by id")
		return nil, err
	}
	return perfomance, nil
}

func (c *perfomanceRepo) GetAll(ctx context.Context, req *tc.GetListPerfomanceRequest) (*tc.GetListPerfomanceResponse, error) {
	perfomances := tc.GetListPerfomanceResponse{}
	var (
		created_at    sql.NullString
		updated_at    sql.NullString
		start_working sql.NullString
		end_working   sql.NullString
	)
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND fullname ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				branch_id,
				user_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				salary,
				ielts_score,
				ielts_attempts_count,
				start_working,
				end_working,
				created_at,
				updated_at
			FROM perfomances
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all perfomances")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			perfomance tc.GetPerfomance
		)
		if err = rows.Scan(
			&perfomance.Id,
			&perfomance.BranchId,
			&perfomance.UserLogin,
			&perfomance.Birthday,
			&perfomance.Gender,
			&perfomance.Fullname,
			&perfomance.Email,
			&perfomance.Phone,
			&perfomance.Salary,
			&perfomance.IeltsScore,
			&perfomance.IeltsAttemptsCount,
			&start_working,
			&end_working,
			&created_at,
			&updated_at,
		); err != nil {
			return &perfomances, err
		}
		perfomance.StartWorking = pkg.NullStringToString(start_working)
		perfomance.EndWorking = pkg.NullStringToString(end_working)
		perfomance.CreatedAt = pkg.NullStringToString(created_at)
		perfomance.UpdatedAt = pkg.NullStringToString(updated_at)

		perfomances.Perfomances = append(perfomances.Perfomances, &perfomance)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from perfomances WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&perfomances.Count)
	if err != nil {
		return &perfomances, err
	}

	return &perfomances, nil
}

func (c *perfomanceRepo) GetById(ctx context.Context, id *tc.PerfomancePrimaryKey) (*tc.GetPerfomance, error) {
	var (
		perfomance    tc.GetPerfomance
		created_at    sql.NullString
		updated_at    sql.NullString
		start_working sql.NullString
		end_working   sql.NullString
	)

	query := `SELECT
				id,
				branch_id,
				user_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				salary,
				ielts_score,
				ielts_attempts_count,
				start_working,
				end_working,
				created_at,
				updated_at
			FROM perfomances
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&perfomance.Id,
		&perfomance.BranchId,
		&perfomance.UserLogin,
		&perfomance.Birthday,
		&perfomance.Gender,
		&perfomance.Fullname,
		&perfomance.Email,
		&perfomance.Phone,
		&perfomance.Salary,
		&perfomance.IeltsScore,
		&perfomance.IeltsAttemptsCount,
		&start_working,
		&end_working,
		&created_at,
		&updated_at); err != nil {
		return &perfomance, err
	}
	perfomance.StartWorking = pkg.NullStringToString(start_working)
	perfomance.EndWorking = pkg.NullStringToString(end_working)
	perfomance.CreatedAt = pkg.NullStringToString(created_at)
	perfomance.UpdatedAt = pkg.NullStringToString(updated_at)

	return &perfomance, nil
}

func (c *perfomanceRepo) Delete(ctx context.Context, id *tc.PerfomancePrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE perfomances SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
