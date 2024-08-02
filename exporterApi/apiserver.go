package exporterApi

import (
	"CollectLet/cache"
	"CollectLet/collector"
	"CollectLet/constants"
	"CollectLet/logger"
	"encoding/json"
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

var cacheFactory *cache.LakeFactory

var computeLake cache.Lake[cache.ComputeCache]

type HttpServer struct {
	Server *http.Server
	mux    *http.ServeMux
}

func init() {
	cacheFactory = cache.NewLakeFactory()
	computeLakeInterface, err := cacheFactory.GetObject(constants.Compute)
	if err != nil {
		logger.GetLogger().Error("%s Error getting cache Lake: %s", logTag, err.Error())
	}
	computeLake = computeLakeInterface.(cache.Lake[cache.ComputeCache])
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
	_, err := w.Write([]byte("sample"))
	if err != nil {
		logger.GetLogger().Error("%s %s", logTag, err.Error())
	}
}

func CPUInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	data, err := computeLake.Get()
	if err != nil {
		logger.GetLogger().Error("%s error getting data from computeLake:%s", logTag, err.Error())
	}
	jsonData, err := json.Marshal(data)
	_, err = w.Write(jsonData)
	if err != nil {
		logger.GetLogger().Error("%s %s", logTag, err.Error())
	}
}

func memoryInfo(w http.ResponseWriter, r *http.Request) {
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Decrease(w http.ResponseWriter, r *http.Request) {

}

func (s *HttpServer) Start() {
	s.mux.HandleFunc("/", Index)
	s.mux.HandleFunc("/monitor/cpu", CPUInfo)
	s.mux.HandleFunc("/monitor/mem", memoryInfo)
	s.mux.HandleFunc("/limit/add", Add)
	s.mux.HandleFunc("/limit/decrease", Decrease)
	s.Server.Handler = s.mux
	collector.Start()
	go func() {
		log.Printf("Starting server on %s\n", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogger().Fatal("%s ListenAndServe: %s", logTag, err.Error())
		}
	}()
}
