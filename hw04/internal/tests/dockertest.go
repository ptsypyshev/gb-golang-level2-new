package tests

import (
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	postgresImg = "postgres"
	postgresTag = "14.2-alpine"
	dbName      = "friends"
	dbUser      = "postgres"
	dbPassword  = "postgres"
	dbPort      = 5432
)

func Start() (*dockertest.Pool, *dockertest.Resource) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: postgresImg,
			Tag:        postgresTag,
			Env: []string{
				fmt.Sprintf("POSTGRES_DB=%s", dbName),
				fmt.Sprintf("POSTGRES_USER=%s", dbUser),
				fmt.Sprintf("POSTGRES_PASSWORD=%s", dbPassword),
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"5432/tcp": {
					{HostIP: "localhost", HostPort: fmt.Sprintf("%d/tcp", dbPort)},
				},
			},
		}, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return pool, resource
}

func Stop(pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		fmt.Printf("Could not purge resource: %s\n", err)
	}

	fmt.Printf("Purge resource: %s\n", "OK")
}
