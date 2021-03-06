package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// EvaluateResults is a nested struct in drds response
type EvaluateResults struct {
	TaskId            int               `json:"TaskId" xml:"TaskId"`
	TaskName          string            `json:"TaskName" xml:"TaskName"`
	InstId            int               `json:"InstId" xml:"InstId"`
	DbName            string            `json:"DbName" xml:"DbName"`
	TaskStatus        string            `json:"TaskStatus" xml:"TaskStatus"`
	GmtCreate         string            `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified       string            `json:"GmtModified" xml:"GmtModified"`
	RdsType           string            `json:"RdsType" xml:"RdsType"`
	RdsCount          int               `json:"RdsCount" xml:"RdsCount"`
	RdsTotalSpace     string            `json:"RdsTotalSpace" xml:"RdsTotalSpace"`
	DrdsType          string            `json:"DrdsType" xml:"DrdsType"`
	DrdsCount         int               `json:"DrdsCount" xml:"DrdsCount"`
	AllSqlCount       int               `json:"AllSqlCount" xml:"AllSqlCount"`
	SlowSqlCount      int               `json:"SlowSqlCount" xml:"SlowSqlCount"`
	TableShardResults TableShardResults `json:"TableShardResults" xml:"TableShardResults"`
	RdsInstInfos      RdsInstInfos      `json:"RdsInstInfos" xml:"RdsInstInfos"`
}
