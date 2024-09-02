/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/task/types"
)

func getStepRecord(t *types.Task) []*StepRecord {
	records := make([]*StepRecord, 0, len(t.Steps))
	for _, step := range t.Steps {
		record := &StepRecord{
			TaskID:              t.TaskID,
			Name:                step.Name,
			Alias:               step.Alias,
			Executor:            step.Executor,
			Payload:             step.Payload,
			Status:              step.Status,
			Message:             step.Message,
			SkipOnFailed:        step.SkipOnFailed,
			ETA:                 step.ETA,
			RetryCount:          step.RetryCount,
			MaxRetries:          step.MaxRetries,
			Params:              step.Params,
			Start:               step.Start,
			End:                 step.End,
			ExecutionTime:       step.ExecutionTime,
			MaxExecutionSeconds: step.MaxExecutionSeconds,
		}
		records = append(records, record)
	}

	return records
}
func getTaskRecord(t *types.Task) *TaskRecord {
	stepSequence := make([]string, 0, len(t.Steps))
	for i := range t.Steps {
		stepSequence = append(stepSequence, t.Steps[i].Name)
	}

	record := &TaskRecord{
		TaskID:              t.TaskID,
		TaskType:            t.TaskType,
		TaskIndex:           t.TaskIndex,
		TaskIndexType:       t.TaskIndexType,
		TaskName:            t.TaskName,
		CurrentStep:         t.CurrentStep,
		StepSequence:        stepSequence,
		CallbackName:        t.CallbackName,
		CommonParams:        t.CommonParams,
		CommonPayload:       t.CommonPayload,
		Status:              t.Status,
		Message:             t.Message,
		Start:               t.Start,
		End:                 t.End,
		ExecutionTime:       t.ExecutionTime,
		MaxExecutionSeconds: t.MaxExecutionSeconds,
		Creator:             t.Creator,
		Updater:             t.Updater,
	}
	return record
}

func toTask(task *TaskRecord, steps []*StepRecord) *types.Task {
	t := &types.Task{
		TaskID:              task.TaskID,
		TaskType:            task.TaskType,
		TaskIndex:           task.TaskIndex,
		TaskIndexType:       task.TaskIndexType,
		TaskName:            task.TaskName,
		CurrentStep:         task.CurrentStep,
		CallbackName:        task.CallbackName,
		CommonParams:        task.CommonParams,
		CommonPayload:       task.CommonPayload,
		Status:              task.Status,
		Message:             task.Message,
		Start:               task.Start,
		End:                 task.End,
		ExecutionTime:       task.ExecutionTime,
		MaxExecutionSeconds: task.MaxExecutionSeconds,
		Creator:             task.Creator,
		LastUpdate:          task.UpdatedAt,
		Updater:             task.Updater,
	}

	stepMap := map[string]*StepRecord{}
	for _, step := range steps {
		stepMap[step.Name] = step
	}

	t.Steps = make([]*types.Step, 0, len(task.StepSequence))
	for _, step := range task.StepSequence {
		step, ok := stepMap[step]
		if !ok {
			continue
		}
		t.Steps = append(t.Steps, step.ToStep())
	}
	return t
}

var (
	// updateTaskField task 支持更新的字段
	updateTaskField = []string{
		"CurrentStep",
		"CommonParams",
		"CommonPayload",
		"Status",
		"Message",
		"Start",
		"End",
		"ExecutionTime",
		"Updater",
	}

	// updateStepField step 支持更新的字段
	updateStepField = []string{
		"Params",
		"Payload",
		"Status",
		"Message",
		"Start",
		"End",
		"ExecutionTime",
		"RetryCount",
	}
)

func getUpdateTaskRecord(t *types.Task) *TaskRecord {
	record := &TaskRecord{
		CurrentStep:   t.CurrentStep,
		CommonParams:  t.CommonParams,
		CommonPayload: t.CommonPayload,
		Status:        t.Status,
		Message:       t.Message,
		Start:         t.Start,
		End:           t.End,
		ExecutionTime: t.ExecutionTime,
		Updater:       t.Updater,
	}
	return record
}

func getUpdateStepRecord(t *types.Step) *StepRecord {
	record := &StepRecord{
		Params:        t.Params,
		Payload:       t.Payload,
		Status:        t.Status,
		Message:       t.Message,
		Start:         t.Start,
		End:           t.End,
		ExecutionTime: t.ExecutionTime,
		RetryCount:    t.RetryCount,
	}
	return record
}
