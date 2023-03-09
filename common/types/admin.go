// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package types

// AddSearchAttributeRequest is an internal type (TBD...)
type AddSearchAttributeRequest struct {
	SearchAttribute map[string]IndexedValueType `json:"searchAttribute,omitempty"`
	SecurityToken   string                      `json:"securityToken,omitempty"`
}

// GetSearchAttribute is an internal getter (TBD...)
func (v *AddSearchAttributeRequest) GetSearchAttribute() (o map[string]IndexedValueType) {
	if v != nil && v.SearchAttribute != nil {
		return v.SearchAttribute
	}
	return
}

// DescribeClusterResponse is an internal type (TBD...)
type DescribeClusterResponse struct {
	SupportedClientVersions *SupportedClientVersions    `json:"supportedClientVersions,omitempty"`
	MembershipInfo          *MembershipInfo             `json:"membershipInfo,omitempty"`
	PersistenceInfo         map[string]*PersistenceInfo `json:"persistenceInfo,omitempty"`
}

// AdminDescribeWorkflowExecutionRequest is an internal type (TBD...)
type AdminDescribeWorkflowExecutionRequest struct {
	Domain    string             `json:"domain,omitempty"`
	Execution *WorkflowExecution `json:"execution,omitempty"`
}

// GetDomain is an internal getter (TBD...)
func (v *AdminDescribeWorkflowExecutionRequest) GetDomain() (o string) {
	if v != nil {
		return v.Domain
	}
	return
}

// AdminDescribeWorkflowExecutionResponse is an internal type (TBD...)
type AdminDescribeWorkflowExecutionResponse struct {
	ShardID                string `json:"shardId,omitempty"`
	HistoryAddr            string `json:"historyAddr,omitempty"`
	MutableStateInCache    string `json:"mutableStateInCache,omitempty"`
	MutableStateInDatabase string `json:"mutableStateInDatabase,omitempty"`
}

// GetShardID is an internal getter (TBD...)
func (v *AdminDescribeWorkflowExecutionResponse) GetShardID() (o string) {
	if v != nil {
		return v.ShardID
	}
	return
}

// GetMutableStateInDatabase is an internal getter (TBD...)
func (v *AdminDescribeWorkflowExecutionResponse) GetMutableStateInDatabase() (o string) {
	if v != nil {
		return v.MutableStateInDatabase
	}
	return
}

// GetWorkflowExecutionRawHistoryV2Request is an internal type (TBD...)
type GetWorkflowExecutionRawHistoryV2Request struct {
	Domain            string             `json:"domain,omitempty"`
	Execution         *WorkflowExecution `json:"execution,omitempty"`
	StartEventID      *int64             `json:"startEventId,omitempty"`
	StartEventVersion *int64             `json:"startEventVersion,omitempty"`
	EndEventID        *int64             `json:"endEventId,omitempty"`
	EndEventVersion   *int64             `json:"endEventVersion,omitempty"`
	MaximumPageSize   int32              `json:"maximumPageSize,omitempty"`
	NextPageToken     []byte             `json:"nextPageToken,omitempty"`
}

// GetDomain is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetDomain() (o string) {
	if v != nil {
		return v.Domain
	}
	return
}

// GetStartEventID is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetStartEventID() (o int64) {
	if v != nil && v.StartEventID != nil {
		return *v.StartEventID
	}
	return
}

// GetStartEventVersion is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetStartEventVersion() (o int64) {
	if v != nil && v.StartEventVersion != nil {
		return *v.StartEventVersion
	}
	return
}

// GetEndEventID is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetEndEventID() (o int64) {
	if v != nil && v.EndEventID != nil {
		return *v.EndEventID
	}
	return
}

// GetEndEventVersion is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetEndEventVersion() (o int64) {
	if v != nil && v.EndEventVersion != nil {
		return *v.EndEventVersion
	}
	return
}

// GetMaximumPageSize is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Request) GetMaximumPageSize() (o int32) {
	if v != nil {
		return v.MaximumPageSize
	}
	return
}

// GetWorkflowExecutionRawHistoryV2Response is an internal type (TBD...)
type GetWorkflowExecutionRawHistoryV2Response struct {
	NextPageToken  []byte          `json:"nextPageToken,omitempty"`
	HistoryBatches []*DataBlob     `json:"historyBatches,omitempty"`
	VersionHistory *VersionHistory `json:"versionHistory,omitempty"`
}

// GetHistoryBatches is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Response) GetHistoryBatches() (o []*DataBlob) {
	if v != nil && v.HistoryBatches != nil {
		return v.HistoryBatches
	}
	return
}

// GetVersionHistory is an internal getter (TBD...)
func (v *GetWorkflowExecutionRawHistoryV2Response) GetVersionHistory() (o *VersionHistory) {
	if v != nil && v.VersionHistory != nil {
		return v.VersionHistory
	}
	return
}

// PersistenceSetting is used to expose persistence engine settings
type PersistenceSetting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PersistenceFeature is used to expose store specific feature.
// Feature can be cadence or store specific.
type PersistenceFeature struct {
	Key     string `json:"key"`
	Enabled bool   `json:"enabled"`
}

// PersistenceInfo is used to expose store configuration
type PersistenceInfo struct {
	Backend  string                `json:"backend"`
	Settings []*PersistenceSetting `json:"settings,omitempty"`
	Features []*PersistenceFeature `json:"features,omitempty"`
}

// ResendReplicationTasksRequest is an internal type (TBD...)
type ResendReplicationTasksRequest struct {
	DomainID      string `json:"domainID,omitempty"`
	WorkflowID    string `json:"workflowID,omitempty"`
	RunID         string `json:"runID,omitempty"`
	RemoteCluster string `json:"remoteCluster,omitempty"`
	StartEventID  *int64 `json:"startEventID,omitempty"`
	StartVersion  *int64 `json:"startVersion,omitempty"`
	EndEventID    *int64 `json:"endEventID,omitempty"`
	EndVersion    *int64 `json:"endVersion,omitempty"`
}

// GetWorkflowID is an internal getter (TBD...)
func (v *ResendReplicationTasksRequest) GetWorkflowID() (o string) {
	if v != nil {
		return v.WorkflowID
	}
	return
}

// GetRunID is an internal getter (TBD...)
func (v *ResendReplicationTasksRequest) GetRunID() (o string) {
	if v != nil {
		return v.RunID
	}
	return
}

// GetRemoteCluster is an internal getter (TBD...)
func (v *ResendReplicationTasksRequest) GetRemoteCluster() (o string) {
	if v != nil {
		return v.RemoteCluster
	}
	return
}

type GetDynamicConfigRequest struct {
	ConfigName string                 `json:"configName,omitempty"`
	Filters    []*DynamicConfigFilter `json:"filters,omitempty"`
}

type GetDynamicConfigResponse struct {
	Value *DataBlob `json:"value,omitempty"`
}

type UpdateDynamicConfigRequest struct {
	ConfigName   string                `json:"configName,omitempty"`
	ConfigValues []*DynamicConfigValue `json:"configValues,omitempty"`
}

type RestoreDynamicConfigRequest struct {
	ConfigName string                 `json:"configName,omitempty"`
	Filters    []*DynamicConfigFilter `json:"filters,omitempty"`
}

// AdminDeleteWorkflowRequest is an internal type (TBD...)
type AdminDeleteWorkflowRequest struct {
	Domain     string             `json:"domain,omitempty"`
	Execution  *WorkflowExecution `json:"execution,omitempty"`
	SkipErrors bool               `json:"skipErrors,omitempty"`
}

func (v *AdminDeleteWorkflowRequest) GetDomain() (o string) {
	if v != nil {
		return v.Domain
	}
	return
}

// GetExecution is an internal getter (TBD...)
func (v *AdminDeleteWorkflowRequest) GetExecution() (o *WorkflowExecution) {
	if v != nil && v.Execution != nil {
		return v.Execution
	}
	return
}

func (v *AdminDeleteWorkflowRequest) GetSkipErrors() (o bool) {
	if v != nil {
		return v.SkipErrors
	}
	return
}

type AdminDeleteWorkflowResponse struct {
	HistoryDeleted    bool `json:"historyDeleted,omitempty"`
	ExecutionsDeleted bool `json:"executionsDeleted,omitempty"`
	VisibilityDeleted bool `json:"visibilityDeleted,omitempty"`
}

type AdminMaintainWorkflowRequest = AdminDeleteWorkflowRequest
type AdminMaintainWorkflowResponse = AdminDeleteWorkflowResponse

type ListDynamicConfigRequest struct {
	ConfigName string `json:"configName,omitempty"`
}

type ListDynamicConfigResponse struct {
	Entries []*DynamicConfigEntry `json:"entries,omitempty"`
}
