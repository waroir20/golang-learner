package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

type PostgresContainer struct {
	container  *testcontainers.Container
	dbURL      string
	port       int
	host       string
	dbName     string
	schemaName string
	username   string
	password   string
}

// NewTestContainer - Used to spin up Postgres using a specific set of credentials with a default database
func NewTestContainer(ctx context.Context, dbName string, schemaName string, username string, password string) (PostgresContainer, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": password,
		"POSTGRES_USER":     username,
		"POSTGRES_DB":       dbName,
	}
	var port = "5432/tcp"
	var filesToCopy []testcontainers.ContainerFile
	if schemaName != "" {
		err := os.WriteFile("init-schema.sql", []byte(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)), 777)
		if err != nil {
			return PostgresContainer{}, err
		}
		filesToCopy = append(filesToCopy, testcontainers.ContainerFile{
			HostFilePath:      "./init-schema.sql",
			ContainerFilePath: "/docker-entrypoint-initdb.d/init.sql",
			FileMode:          700,
		})
	}
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForListeningPort("5432/tcp"),
			Hostname:     "postgres",
			Files:        filesToCopy,
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return PostgresContainer{
			container: &container,
		}, fmt.Errorf("failed to start container: %s", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return PostgresContainer{
			container: &container,
		}, fmt.Errorf("failed to get container external port: %s", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return PostgresContainer{
			container: &container,
		}, fmt.Errorf("failed to get container host: %s", err)
	}

	log.Info().Msgf("postgres container ready and running at port: ", mappedPort)

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, hostIP, mappedPort.Port(), dbName)
	database, err := sql.Open("postgres", url)
	if err != nil {
		return PostgresContainer{container: &container}, fmt.Errorf("failed to establish database connection: %s", err)
	}
	defer func() {
		_ = database.Close()
	}()

	return PostgresContainer{
		container:  &container,
		dbURL:      url,
		host:       hostIP,
		port:       mappedPort.Int(),
		dbName:     dbName,
		schemaName: schemaName,
		username:   username,
		password:   password,
	}, nil
}

func (c PostgresContainer) Config() Config {
	return Config{
		Host:          c.host,
		Port:          c.port,
		EnableSSL:     false,
		DatabaseName:  c.dbName,
		DefaultSchema: c.schemaName,
		Username:      c.username,
		Password:      c.password,
	}
}

func (c PostgresContainer) Shutdown(ctx context.Context) error {
	if err := os.RemoveAll("./init-schema.sql"); err != nil {
		log.Error().Stack().Err(err).Msg("Failed to remove init-schema file")
	}
	return (*c.container).Terminate(ctx)
}
