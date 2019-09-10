package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/nhaancs/grabvn-golang-bootcamp/week6/cslogger"
	"github.com/nhaancs/grabvn-golang-bootcamp/week6/datadog"
)

// Server represents our server.
type Server struct {
	logger  cslogger.CSLogger
	datadog *datadog.Client
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() {
	s.logger.Info("echo server is starting on port 8080...")
	http.HandleFunc("/", s.echo)
	http.ListenAndServe(":8080", nil)
}

// Echo echos back the request as a response
func (s *Server) echo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// 30% chance of failure
	if rand.Intn(100) < 30 {
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		s.logger.Message("a chaos monkey broke your server")

		tags := []string{"datadog-demo", "error"}
		ok := s.datadog.Event("Error", "error", "server down", "a chaos monkey broke your server", "Demo App", tags)
		if !ok {
			log.Println("Cannot fire datadog event.")
		}

		return
	}

	// Happy path
	writer.WriteHeader(200)
	request.Write(writer)

	tags := []string{"datadog-demo", "success"}
	ok := s.datadog.Event("Error", "error", "server down", "a chaos monkey broke your server", "Demo App", tags)
	if !ok {
		log.Println("Cannot fire datadog event.")
	}
}

func initLogger(logFile *os.File) cslogger.CSLogger {
	logger := cslogger.NewCSLogger()
	logger.SetLogLevel(cslogger.ErrorLevel)

	if logFile != nil {
		logger.SetOutput(logFile)
	}

	return logger
}

func main() {
	logFile, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err == nil {
		defer logFile.Close()
	}

	server := Server{
		logger:  initLogger(logFile),
		datadog: datadog.Connect("demo-app", "127.0.0.1", "8125"),
	}

	// Start the server
	server.ListenAndServe()
}
