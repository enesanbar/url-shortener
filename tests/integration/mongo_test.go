package integration

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MongoContainer struct {
	container testcontainers.Container
	hostname  string
	port      string
	username  string
	password  string
}

func NewMongoContainer(username string, password string) *MongoContainer {
	return &MongoContainer{username: username, password: password}
}

func (m *MongoContainer) Start(ctx context.Context) error {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.4",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": m.username,
			"MONGO_INITDB_ROOT_PASSWORD": m.password,
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	mappedPort, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return err
	}

	m.container = container
	m.SetHostname(hostIP)
	m.SetPort(mappedPort.Port())

	return nil
}

func (m MongoContainer) Stop(ctx context.Context) error {
	return m.container.Terminate(ctx)
}

func (m *MongoContainer) Hostname() string {
	return m.hostname
}

func (m *MongoContainer) SetHostname(hostname string) {
	m.hostname = hostname
}

func (m *MongoContainer) Port() string {
	return m.port
}

func (m *MongoContainer) SetPort(port string) {
	m.port = port
}

func (m *MongoContainer) Username() string {
	return m.username
}

func (m *MongoContainer) SetUsername(username string) {
	m.username = username
}

func (m *MongoContainer) Password() string {
	return m.password
}

func (m *MongoContainer) SetPassword(password string) {
	m.password = password
}
