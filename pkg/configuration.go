package main

import "github.com/ydataai/backend-interview-sample/pkg/common"

type Configuration struct {
	// logLevel             logrus.Level
	enableLeaderElection bool
	metricsPort          string
}

func (c *Configuration) LoadEnvVars() error {
	logLevel, err := common.VariableFromEnvironment("LOG_LEVEL")
	if err != nil {
		return err
	}
	print(logLevel)

	// level, err := logrus.ParseLevel(logLevel)
	// if err != nil {
	// 	return err
	// }

	// c.logLevel = level

	enableLeaderElection, err := common.BooleanVariableFromEnvironment("ENABLE_LEADER_ELECTION")
	if err != nil {
		return err
	}
	c.enableLeaderElection = enableLeaderElection

	c.metricsPort, err = common.VariableFromEnvironment("METRICS_PORT")
	if err != nil {
		return err
	}

	return nil
}
