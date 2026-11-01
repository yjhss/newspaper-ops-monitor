package workflow

import (
	"fmt"
	"strings"
)

type Stage string

const (
	StageInput    Stage = "INPUT"
	StagePDF      Stage = "PDF"
	StageRIPInput Stage = "RIP_INPUT"
	StageTIFF     Stage = "TIFF"
	StageJPG      Stage = "JPG"
	StageCTP      Stage = "CTP"
	StageBPQueue  Stage = "BP_QUEUE"
	StageYNQueue  Stage = "YN_QUEUE"
)

func ParseStage(value string) (Stage, error) {
	stage := Stage(strings.ToUpper(strings.TrimSpace(value)))

	switch stage {
	case StageInput,
		StagePDF,
		StageRIPInput,
		StageTIFF,
		StageJPG,
		StageCTP,
		StageBPQueue,
		StageYNQueue:
		return stage, nil
	default:
		return "", fmt.Errorf("unsupported workflow stage: %s", value)
	}
}
