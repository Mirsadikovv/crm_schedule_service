package storage

import (
	"context"
	"go_schedule_service/genproto/journal_service"
	"go_schedule_service/genproto/lesson_service"
	"go_schedule_service/genproto/schedule_service"
	"go_schedule_service/genproto/student_perfomance_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Lesson() LessonRepoI
	Journal() JournalRepoI
	Schedule() ScheduleRepoI
	Perfomance() PerfomanceRepoI
}

type LessonRepoI interface {
	Create(context.Context, *lesson_service.CreateLesson) (*lesson_service.GetLesson, error)
	Update(context.Context, *lesson_service.UpdateLesson) (*lesson_service.GetLesson, error)
	GetAll(context.Context, *lesson_service.GetListLessonRequest) (*lesson_service.GetListLessonResponse, error)
	GetById(context.Context, *lesson_service.LessonPrimaryKey) (*lesson_service.GetLesson, error)
	Delete(context.Context, *lesson_service.LessonPrimaryKey) (emptypb.Empty, error)
}

type JournalRepoI interface {
	Create(context.Context, *journal_service.CreateJournal) (*journal_service.GetJournal, error)
	Update(context.Context, *journal_service.UpdateJournal) (*journal_service.GetJournal, error)
	GetAll(context.Context, *journal_service.GetListJournalRequest) (*journal_service.GetListJournalResponse, error)
	GetById(context.Context, *journal_service.JournalPrimaryKey) (*journal_service.GetJournal, error)
	Delete(context.Context, *journal_service.JournalPrimaryKey) (emptypb.Empty, error)
}

type ScheduleRepoI interface {
	Create(context.Context, *schedule_service.CreateSchedule) (*schedule_service.GetSchedule, error)
	Update(context.Context, *schedule_service.UpdateSchedule) (*schedule_service.GetSchedule, error)
	GetAll(context.Context, *schedule_service.GetListScheduleRequest) (*schedule_service.GetListScheduleResponse, error)
	GetById(context.Context, *schedule_service.SchedulePrimaryKey) (*schedule_service.GetSchedule, error)
	Delete(context.Context, *schedule_service.SchedulePrimaryKey) (emptypb.Empty, error)
}

type PerfomanceRepoI interface {
	Create(context.Context, *student_perfomance_service.CreateStudentPerfomance) (*student_perfomance_service.GetStudentPerfomance, error)
	Update(context.Context, *student_perfomance_service.UpdateStudentPerfomance) (*student_perfomance_service.GetStudentPerfomance, error)
	GetAll(context.Context, *student_perfomance_service.GetListStudentPerfomanceRequest) (*student_perfomance_service.GetListStudentPerfomanceResponse, error)
	GetById(context.Context, *student_perfomance_service.StudentPerfomancePrimaryKey) (*student_perfomance_service.GetStudentPerfomance, error)
	Delete(context.Context, *student_perfomance_service.StudentPerfomancePrimaryKey) (emptypb.Empty, error)
}
