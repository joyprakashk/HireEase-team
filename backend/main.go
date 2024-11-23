package main

import (
	"gofr.dev/pkg/gofr"

	jobhandler "hire-ease/handler/job"
	rechandler "hire-ease/handler/recruiter"
	stdhandler "hire-ease/handler/student"
	"hire-ease/migrations"
	jobsvc "hire-ease/service/job"
	recsvc "hire-ease/service/recruiter"
	stdsvc "hire-ease/service/student"
	jobstr "hire-ease/store/job"
	recstr "hire-ease/store/recruiter"
	stdstr "hire-ease/store/student"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())

	jobStr := jobstr.New()
	jobSvc := jobsvc.New(jobStr)
	jobHndlr := jobhandler.New(jobSvc)

	app.POST("/jobs", jobHndlr.Create)
	app.GET("/jobs", jobHndlr.GetAll)
	app.GET("/jobs/{id}", jobHndlr.GetByID)
	app.PUT("/jobs/{id}", jobHndlr.Update)
	app.DELETE("/jobs/{id}", jobHndlr.Delete)

	recStr := recstr.New()
	recSvc := recsvc.New(recStr)
	recHndlr := rechandler.New(recSvc)

	app.POST("/recruiters", recHndlr.Create)
	app.GET("/recruiters", recHndlr.GetAll)
	app.GET("/recruiters/{id}", recHndlr.GetByID)
	app.PUT("/recruiters/{id}", recHndlr.Update)
	app.DELETE("/recruiters/{id}", recHndlr.Delete)
	app.POST("/recruiters/login/{id}", recHndlr.Login)

	stdStr := stdstr.New()
	stdSvc := stdsvc.New(stdStr)
	stdHndlr := stdhandler.New(stdSvc)

	app.POST("/students", stdHndlr.Create)
	app.GET("/students", stdHndlr.GetAll)
	app.GET("/students/{id}", stdHndlr.GetByID)
	app.PUT("/students/{id}", stdHndlr.Update)
	app.DELETE("/students/{id}", stdHndlr.Delete)
	app.POST("/students/login/{id}", stdHndlr.Login)

	app.Run()
}
