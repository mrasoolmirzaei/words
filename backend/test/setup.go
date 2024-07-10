package test

import (
	"github.com/mrasoolmirzaei/words/backend/internal/db"
	"github.com/mrasoolmirzaei/words/backend/pkg/server"
	"github.com/mrasoolmirzaei/words/backend/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"time"
)

const (
	apiListen = "localhost:8090"
)

type testSuite struct {
	suite.Suite
	server *server.Server
}

func (suite *testSuite) SetupSuite() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stderr)

	dbMock := &db.PQMock{}
	api := api.NewAPI(dbMock, logger)
	config := &server.Config{
		Logger: logger,
		Api:    api,
	}
	srv, err := server.NewServer(config)
	if err != nil {
		suite.FailNow(err.Error())
	}
	suite.server = srv

	go func() {
		suite.NoError(srv.Serve(apiListen))
	}()
	time.Sleep(1 * time.Second)
}

func (suite *testSuite) SetupTest() {
}

func (suite *testSuite) TearDownSuite() {
	suite.NoError(suite.server.Stop())
}
