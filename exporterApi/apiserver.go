package exporterApi

import (
	"CollectLet/logger"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

var logTag = "[apiServer]"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

type HttpServer struct {
	Server *http.Server
	mux    *http.ServeMux
}

func NewHttpServer() *HttpServer {
	byteValue, err := os.ReadFile("./config/server.yaml")
	if err != nil {
		logger.GetLogger().Error("%s %s", logTag, err.Error())
	}
	var config Config
	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		logger.GetLogger().Error("%s %s", logTag, err.Error())
	}
	return &HttpServer{
		mux: http.NewServeMux(),
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		},
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		logger.GetLogger().Error("%s %s", logTag, err.Error())
	}
}

func (s *HttpServer) Start() {
	s.mux.HandleFunc("/", Index)
	s.Server.Handler = s.mux
	go func() {
		log.Printf("Starting server on %s\n", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogger().Fatal("%s ListenAndServe: %s", logTag, err.Error())
		}
	}()
}
