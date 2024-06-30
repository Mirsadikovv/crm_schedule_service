package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sc "go_schedule_service/genproto/schedule_service"
	"go_schedule_service/pkg"
	"go_schedule_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type scheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) storage.ScheduleRepoI {
	return &scheduleRepo{
		db: db,
	}
}

func (c *scheduleRepo) Create(ctx context.Context, req *sc.CreateSchedule) (*sc.GetSchedule, error) {

	id := uuid.NewString()

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO schedules (
			id,
			group_id,
			lesson_id,
			classroom,
			type_of_group,
			task,
			deadline,
			score,
			start_time,
			end_time
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10
		)`,
		id,
		req.GroupId,
		req.LessonId,
		req.Classroom,
		req.TypeOfGroup,
		req.Task,
		req.Deadline,
		req.Score,
		req.StartTime,
		req.EndTime,
	)
	if err != nil {
		log.Println("error while creating schedules", comtag)
		return nil, err
	}

	schedules, err := c.GetById(ctx, &sc.SchedulePrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting schedules by id")
		return nil, err
	}
	return schedules, nil
}

func (c *scheduleRepo) Update(ctx context.Context, req *sc.UpdateSchedule) (*sc.GetSchedule, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE schedules SET
		group_id = $1,
		lesson_id = $2,
		classroom = $3,
		type_of_group = $4,
		task = $5,
		deadline = $6,
		score = $7,
		start_time = $8,
		end_time = $9,
		updated_at = NOW()
		WHERE id = $10
		`,
		req.GroupId,
		req.LessonId,
		req.Classroom,
		req.TypeOfGroup,
		req.Task,
		req.Deadline,
		req.Score,
		req.StartTime,
		req.EndTime,
		req.Id)
	if err != nil {
		log.Println("error while updating schedules")
		return nil, err
	}

	schedules, err := c.GetById(ctx, &sc.SchedulePrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting schedules by id")
		return nil, err
	}
	return schedules, nil
}

func (c *scheduleRepo) GetAll(ctx context.Context, req *sc.GetListScheduleRequest) (*sc.GetListScheduleResponse, error) {
	schedules := sc.GetListScheduleResponse{}
	var (
		deadline   sql.NullString
		start_time sql.NullString
		end_time   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)
	filter_by_name := ""
	offset := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND theme ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				group_id,
				lesson_id,
				classroom,
				type_of_group,
				task,
				deadline,
				score,
				start_time,
				end_time,
				created_at,
				updated_at
			FROM schedules
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offset, req.Limit)

	if err != nil {
		log.Println("error while getting all schedules")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			schedule sc.GetSchedule
		)
		if err = rows.Scan(
			&schedule.Id,
			&schedule.GroupId,
			&schedule.LessonId,
			&schedule.Classroom,
			&schedule.TypeOfGroup,
			&schedule.Task,
			&deadline,
			&schedule.Score,
			&start_time,
			&end_time,
			&created_at,
			&updated_at,
		); err != nil {
			return &schedules, err
		}
		schedule.Deadline = pkg.NullStringToString(deadline)
		schedule.StartTime = pkg.NullStringToString(start_time)
		schedule.EndTime = pkg.NullStringToString(end_time)
		schedule.CreatedAt = pkg.NullStringToString(created_at)
		schedule.UpdatedAt = pkg.NullStringToString(updated_at)

		schedules.Schedules = append(schedules.Schedules, &schedule)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from schedules WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&schedules.Count)
	if err != nil {
		return &schedules, err
	}

	return &schedules, nil
}

func (c *scheduleRepo) GetById(ctx context.Context, id *sc.SchedulePrimaryKey) (*sc.GetSchedule, error) {
	var (
		deadline   sql.NullString
		schedule   sc.GetSchedule
		start_time sql.NullString
		end_time   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT
				id,
				group_id,
				lesson_id,
				classroom,
				type_of_group,
				task,
				deadline,
				score,
				start_time,
				end_time,
				created_at,
				updated_at
			FROM schedules
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&schedule.Id,
		&schedule.GroupId,
		&schedule.LessonId,
		&schedule.Classroom,
		&schedule.TypeOfGroup,
		&schedule.Task,
		&deadline,
		&schedule.Score,
		&start_time,
		&end_time,
		&created_at,
		&updated_at); err != nil {
		return &schedule, err
	}
	schedule.Deadline = pkg.NullStringToString(deadline)
	schedule.StartTime = pkg.NullStringToString(start_time)
	schedule.EndTime = pkg.NullStringToString(end_time)
	schedule.CreatedAt = pkg.NullStringToString(created_at)
	schedule.UpdatedAt = pkg.NullStringToString(updated_at)

	return &schedule, nil
}

func (c *scheduleRepo) Delete(ctx context.Context, id *sc.SchedulePrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE schedules SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
