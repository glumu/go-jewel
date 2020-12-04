package gl_machinery

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"

	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1/log"
	tasks2 "github.com/RichardKnop/machinery/v1/tasks"
)

func (mery *Machinery) Worker(tasks map[string]interface{}) (*machinery.Worker, error) {
	cleanup, err := tracers.SetupTracer(mery.Config.ConsumerTag)
	if err != nil {
		log.FATAL.Fatalln("Unable to instantiate a tracer:", err)
	}
	defer cleanup()

	err = mery.Server.RegisterTasks(tasks)
	if err != nil {
		return nil, err
	}

	worker := mery.Server.NewWorker(mery.Config.ConsumerTag, 10)

	errorHandler := func(err error) {
		if mery.ErrorLog != nil {
			mery.ErrorLog(Logger{
				Level: "ERROR",
				Type:  "[Worker Error]",
				Msg:   fmt.Sprintf("I am an error handler: %s", err.Error()),
			})
		} else {
			log.ERROR.Println("I am an error handler:", err)
		}
	}

	preTaskHandler := func(signature *tasks2.Signature) {
		if mery.InfoLog != nil {
			mery.InfoLog(Logger{
				Level: "INFO",
				Type:  "[Worker Start]",
				Msg:   fmt.Sprintf("start task handler:", signature.Name),
			})
		}
	}

	postTaskHandler := func(signature *tasks2.Signature) {
		if mery.InfoLog != nil {
			mery.InfoLog(Logger{
				Level: "INFO",
				Type:  "[Worker End]",
				Msg:   fmt.Sprintf("end task handler:", signature.Name),
			})
		}
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)

	err = worker.Launch()
	return worker, err
}
