package fileNotifier

import (
	"github.com/fsnotify/fsnotify"
	"httpTools/src/domain/fileNotifier/model"
	"httpTools/src/domain/fileNotifier/model/fileWatcher"
	"httpTools/src/domain/httpServer"
	"httpTools/src/infrastructure/event"
	"httpTools/src/infrastructure/event/vo"
	"log"
)

const EventName = "StartWatcher"

type Event struct {
	server  *model.Service
	emitter *event.Emitter
}

func NewEvent(emitter *event.Emitter) *Event {
	e := &Event{
		emitter: emitter,
	}

	// 注册事件
	e.Register()
	return e
}

func (s *Event) Register() {
	s.emitter.On(EventName, s.CreateHandler())
}

func (s *Event) CreateHandler() vo.HFunc {
	return func(data vo.VData) {
		if s.server != nil {
			s.server.Quit()
		}

		// 监视器服务
		s.server = model.NewFileWatcherService(fileWatcher.NewFileStatusWatcher(
			fileWatcher.WithFilename("config.yml"),
			fileWatcher.WithWriteHandlers(func(event fsnotify.Event) {
				log.Println("写入文件 : ", event.Name)
				s.emitter.Emit(httpServer.EventName, event)
			}),
		))

		// 启动监视器
		go func() {
			if err := s.server.Run(); err != nil {
				log.Fatal(err)
			}
		}()
	}
}
