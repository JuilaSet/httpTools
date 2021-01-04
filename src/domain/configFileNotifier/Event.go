package configFileNotifier

import (
	"github.com/fsnotify/fsnotify"
	"httpTools/src/domain/configFileNotifier/model"
	"httpTools/src/domain/configFileNotifier/model/fileWatcher"
	"httpTools/src/domain/httpServer"
	"httpTools/src/infrastructure/config"
	"httpTools/src/infrastructure/event"
	"httpTools/src/infrastructure/event/vo"
	"log"
)

const EventName = "StartWatcher"

type Event struct {
	filename string
	server  *model.Server
	emitter *event.Emitter
}

func NewEvent(emitter *event.Emitter, filename string) *Event {
	if filename == "" {
		filename = "config.yml"
	}

	e := &Event{
		filename: filename,
		emitter: emitter,
	}

	// 注册事件
	e.Register()
	return e
}

func (s *Event) Emit() {
	s.emitter.Emit(EventName, true)
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
		s.server = model.NewFileWatcherServer(fileWatcher.NewFileStatusWatcher(
			fileWatcher.WithFilename(s.filename),
			fileWatcher.WithWriteHandlers(func(event fsnotify.Event) {
				log.Println("config file reload:", event.Name)
				// 修改文件时，构建配置文件
				appConfig := config.NewAppConfig(s.filename)
				s.emitter.Emit(httpServer.EventName, appConfig)
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
