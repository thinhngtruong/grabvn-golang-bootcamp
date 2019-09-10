package datadog

import (
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

// Client -> Stuct for client connection
type Client struct {
	Monitor *statsd.Client
}

func prepEvent(category, priority, title, text, source string, tags []string) statsd.Event {
	host, err := os.Hostname()
	if err != nil {
		log.Println("No hostname set for OS. Setting it to localhost.")
		host = "localhost"
	}

	pri := statsd.Success
	switch priority {
	case "error":
		pri = statsd.Error
		break
	case "warning":
		pri = statsd.Warning
		break
	case "info":
		pri = statsd.Info
		break
	default:
		break
	}
	evt := statsd.Event{
		Title:          title,
		Text:           text,
		Timestamp:      time.Now(),
		Hostname:       host,
		AggregationKey: category,
		Priority:       statsd.Normal,
		SourceTypeName: source,
		AlertType:      pri,
		Tags:           tags,
	}

	return evt
}

// Connect -> Connect to datadog agent
func Connect(namespace, host, port string) *Client {
	addr := host + ":" + port
	c, err := statsd.NewBuffered(addr, 2)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	c.Namespace = namespace
	cl := Client{
		Monitor: c,
	}
	return &cl
}

// SimpleEvent -> Sends out simple datadog event to statsd
func (cl *Client) SimpleEvent(title, text string) bool {
	err := cl.Monitor.SimpleEvent(title, text)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

// Event -> Sends out a complete datadog event to statsd
func (cl *Client) Event(category, priority, title, text, source string, tags []string) bool {
	evt := prepEvent(category, priority, title, text, source, tags)
	err := cl.Monitor.Event(&evt)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

// Gauge -> Sends out metrics to datadog
func (cl *Client) Gauge(name string, value float64, tags []string, rate float64) bool {
	err := cl.Monitor.Gauge(name, value, tags, rate)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

/*
// This example uses a custom wrapper over the datadog-go/statsd library
// by datadog for easier event building and error handling. Have a look at
// the wrapper datadog.go in datadog folder.
// The main function here will send a SimpleEvent, a complete event and a metric
// to Datadog.
func testDatadogWrapper() {
	// Creating connection to datadog statsD on localhost:8125
	// Host and port can be different.
	conn := datadog.Connect("example-app", "127.0.0.1", "8125")

	// Firing a simple event. This event has no other information except title and text.
	// Function SimpleEvent(title, text string)
	ok := conn.SimpleEvent("Hello World!", "This is the text that appears in the event.")
	if !ok {
		log.Println("Cannot fire simple datadog event.")
	}

	// Firing a proper event with tags and alert type.
	// Function Event(category, priority, title, text, source string, tags []string)
	tags := []string{"datadog-demo"}
	ok = conn.Event("Error", "error", "Hello World!", "a chaos monkey broke your server", "Demo App", tags)
	if !ok {
		log.Println("Cannot fire datadog event.")
	}

	// Sending a metric to Datadog for graphs and monitoring.
	// Function Gauge(name string, value float64, tags []string, rate float64) bool
	mTags := []string{"hello-world", "example", "metric"}
	for i := 0; i <= 100; i++ {
		ok = conn.Gauge("hello.world.metric", float64(i), mTags, 1)
		time.Sleep(1 * time.Second)
		if !ok {
			log.Println("Cannot send datadog metric.")
		}
	}

	// API Initialize
	connAPI := datadog.Init("11f3172a3d2b5de7c36580c6c7798b7a", "9e244aa64e94817b02db463dceb8254d3aa3d66c")

	// Fire event to datadog using API
	// Function PostEvent(title, text string, tags []string, alertType string)
	ok = connAPI.PostEvent("Hello World!", "This is the text that appears in the event.", tags, "info")

	// Send metric to datadog using API
	// Function SendMetric(metric string, points float64, ty string, tags []string)
	ok = connAPI.SendMetric("test", 20, "gauge", mTags)
	if !ok {
		log.Println("Cannot send datadog metric")
	}
}
*/
