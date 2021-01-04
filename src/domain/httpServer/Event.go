package httpServer

import (
	"encoding/json"
	"httpTools/src/domain/domEvent/ConfigFileReload"
	"httpTools/src/domain/httpServer/model"
	"httpTools/src/infrastructure/config"
	"httpTools/src/infrastructure/event"
	"httpTools/src/infrastructure/event/vo"
	"log"
)

const EventName = "RestartServer"

type Event struct {
	server  *model.Server
	emitter *event.Emitter
	serviceConfig *config.Config
}

func NewEvent(emitter *event.Emitter, serviceConfig *config.Config) *Event {
	e := &Event{
		emitter: emitter,
		serviceConfig: serviceConfig,
	}

	// 注册事件
	e.Register()
	return e
}

func (s *Event) Emit() {
	s.emitter.Emit(EventName, s.serviceConfig)
}

func (s *Event) Register() {
	s.emitter.On(EventName, s.CreateHandler())
}

func (s *Event) CreateHandler() vo.HFunc {
	return func(data vo.VData) {
		if s.server != nil {
			s.server.Quit()
		}

		// 读取配置
		app := model.NewServiceConfig(
			model.WithConfig(ConfigFileReload.FromData(data).Data()),
		)

		// 构建应用
		s.server = model.NewHttpServer(app)

		msg, _ := json.MarshalIndent(app, "", "\t")
		log.Println("app info", string(msg))

		// 启动服务器
		go func(server *model.Server) {
			if err := server.Run(); err != nil {
				panic(err)
			}
		}(s.server)
	}
}
