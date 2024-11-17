package v1

import (
	"github.com/otto8-ai/nah/pkg/fields"
	"github.com/otto8-ai/otto8/apiclient/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	_ fields.Fields = (*WorkflowExecution)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WorkflowExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowExecutionSpec   `json:"spec,omitempty"`
	Status WorkflowExecutionStatus `json:"status,omitempty"`
}

func (in *WorkflowExecution) Has(field string) bool {
	return in.Get(field) != ""
}

func (in *WorkflowExecution) Get(field string) string {
	if in != nil {
		switch field {
		case "spec.webhookName":
			return in.Spec.WebhookName
		case "spec.cronJobName":
			return in.Spec.CronJobName
		case "spec.workflowName":
			return in.Spec.WorkflowName
		case "spec.parentRunName":
			return in.Spec.ParentRunName
		}
	}

	return ""
}

func (in *WorkflowExecution) FieldNames() []string {
	return []string{"spec.webhookName", "spec.cronJobName", "spec.workflowName", "spec.parentRunName"}
}

func (in *WorkflowExecution) GetColumns() [][]string {
	return [][]string{
		{"Name", "Name"},
		{"State", "Status.State"},
		{"Thread", "Status.ThreadName"},
		{"Workflow", "Spec.WorkflowName"},
		{"After", "Spec.AfterWorkflowStepName"},
		{"Created", "{{ago .CreationTimestamp}}"},
	}
}

type WorkflowExecutionSpec struct {
	Input                 string `json:"input,omitempty"`
	WorkflowName          string `json:"workflowName,omitempty"`
	WebhookName           string `json:"webhookName,omitempty"`
	EmailReceiverName     string `json:"emailReceiverName,omitempty"`
	CronJobName           string `json:"cronJobName,omitempty"`
	ParentThreadName      string `json:"parentThreadName,omitempty"`
	ParentRunName         string `json:"parentRunName,omitempty"`
	AfterWorkflowStepName string `json:"afterWorkflowStepName,omitempty"`
	WorkspaceName         string `json:"workspaceName,omitempty"`
	WorkflowGeneration    int64  `json:"workflowGeneration,omitempty"`
	RunUntilStep          string `json:"runUntilStep,omitempty"`
}

func (in *WorkflowExecution) DeleteRefs() []Ref {
	return []Ref{
		{ObjType: &Workflow{}, Name: in.Spec.WorkflowName},
		{ObjType: &Thread{}, Name: in.Status.ThreadName},
	}
}

type WorkflowExecutionStatus struct {
	State              types.WorkflowState     `json:"state,omitempty"`
	Output             string                  `json:"output,omitempty"`
	Error              string                  `json:"error,omitempty"`
	ThreadName         string                  `json:"threadName,omitempty"`
	WorkflowManifest   *types.WorkflowManifest `json:"workflowManifest,omitempty"`
	EndTime            *metav1.Time            `json:"endTime,omitempty"`
	WorkflowGeneration int64                   `json:"workflowGeneration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WorkflowExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []WorkflowExecution `json:"items"`
}
