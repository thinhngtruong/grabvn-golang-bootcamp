package main

import (
	"math/rand"
	"net/http"
	"os"
	"log"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/nhaancs/grabvn-golang-bootcamp/week6/cslogger"
)

// Server represents our server.
type Server struct {
	logger        cslogger.CSLogger
	datadogClient *statsd.Client
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
		s.datadogClient.Count("fail", 1, []string{"localhost"}, 1)
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		s.logger.Message("a chaos monkey broke your server")

		return
	}

	// Happy path
	s.datadogClient.Count("success", 1, []string{"localhost"}, 1)
	writer.WriteHeader(200)
	request.Write(writer)
}

func initLogger(logFile *os.File) cslogger.CSLogger {
	logger := cslogger.NewCSLogger()
	logger.SetLogLevel(cslogger.ErrorLevel)

	if logFile != nil {
		logger.SetOutput(logFile)
	}

	return logger
}

func initDatadogClient() *statsd.Client {
	client, err := statsd.New("127.0.0.1:8125",
		statsd.WithNamespace("cs-grabvn-bootcamp."),                  // prefix every metric with the app name
		statsd.WithTags([]string{"region:us-east-1a"}), // send the EC2 availability zone as a tag with every metric
	)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	logFile, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err == nil {
		defer logFile.Close()
	}

	server := Server{
		logger:        initLogger(logFile),
		datadogClient: initDatadogClient(),
	}

	// Start the server
	server.ListenAndServe()
}
