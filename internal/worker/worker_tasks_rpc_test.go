package worker_test

import (
	"context"
	"github.com/cirruslabs/cirrus-ci-agent/api"
	"io"
	"runtime"
)

type TasksRPC struct {
	SucceededCommands []string

	api.UnimplementedCirrusCIServiceServer
}

func (tasksRPC *TasksRPC) InitialCommands(
	ctx context.Context,
	request *api.InitialCommandsRequest,
) (*api.CommandsResponse, error) {
	var checkScript string
	if runtime.GOOS == "windows" {
		checkScript = "type go.mod"
	} else {
		checkScript = "test -e go.mod"
	}

	return &api.CommandsResponse{
		Environment: map[string]string{
			"CIRRUS_REPO_CLONE_URL": "http://github.com/cirruslabs/cirrus-cli.git",
			"CIRRUS_BRANCH":         "master",
		},
		Commands: []*api.Command{
			{
				Name: "clone",
				Instruction: &api.Command_CloneInstruction{
					CloneInstruction: &api.CloneInstruction{},
				},
			},
			{
				Name: "check",
				Instruction: &api.Command_ScriptInstruction{
					ScriptInstruction: &api.ScriptInstruction{
						Scripts: []string{checkScript},
					},
				},
			},
		},
		ServerToken:      serverSecret,
		TimeoutInSeconds: 3600,
	}, nil
}

func (tasksRPC *TasksRPC) ReportSingleCommand(
	ctx context.Context,
	request *api.ReportSingleCommandRequest,
) (*api.ReportSingleCommandResponse, error) {
	if request.Succeded {
		tasksRPC.SucceededCommands = append(tasksRPC.SucceededCommands, request.CommandName)
	}

	return &api.ReportSingleCommandResponse{}, nil
}

func (tasksRPC *TasksRPC) StreamLogs(stream api.CirrusCIService_StreamLogsServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return stream.SendAndClose(&api.UploadLogsResponse{})
}

func (tasksRPC *TasksRPC) Heartbeat(
	ctx context.Context,
	request *api.HeartbeatRequest,
) (*api.HeartbeatResponse, error) {
	return &api.HeartbeatResponse{}, nil
}
