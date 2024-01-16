// Parts of this file were reproduced verbatim from: https://github.com/segmentio/kafka-go/blob/main/protocol/protocol.go
// kafka-go is licensed under the MIT license:
/*
MIT License

Copyright (c) 2017 Segment

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package protocol

import "strconv"

type ApiKey int16

func (k ApiKey) String() string {
	if v, ok := apiNames[k]; ok {
		return v
	}
	return strconv.Itoa(int(k))
}

const (
	Produce                      ApiKey = 0
	Fetch                        ApiKey = 1
	ListOffsets                  ApiKey = 2
	Metadata                     ApiKey = 3
	LeaderAndIsr                 ApiKey = 4
	StopReplica                  ApiKey = 5
	UpdateMetadata               ApiKey = 6
	ControlledShutdown           ApiKey = 7
	OffsetCommit                 ApiKey = 8
	OffsetFetch                  ApiKey = 9
	FindCoordinator              ApiKey = 10
	JoinGroup                    ApiKey = 11
	Heartbeat                    ApiKey = 12
	LeaveGroup                   ApiKey = 13
	SyncGroup                    ApiKey = 14
	DescribeGroups               ApiKey = 15
	ListGroups                   ApiKey = 16
	SaslHandshake                ApiKey = 17
	ApiVersions                  ApiKey = 18
	CreateTopics                 ApiKey = 19
	DeleteTopics                 ApiKey = 20
	DeleteRecords                ApiKey = 21
	InitProducerId               ApiKey = 22
	OffsetForLeaderEpoch         ApiKey = 23
	AddPartitionsToTxn           ApiKey = 24
	AddOffsetsToTxn              ApiKey = 25
	EndTxn                       ApiKey = 26
	WriteTxnMarkers              ApiKey = 27
	TxnOffsetCommit              ApiKey = 28
	DescribeAcls                 ApiKey = 29
	CreateAcls                   ApiKey = 30
	DeleteAcls                   ApiKey = 31
	DescribeConfigs              ApiKey = 32
	AlterConfigs                 ApiKey = 33
	AlterReplicaLogDirs          ApiKey = 34
	DescribeLogDirs              ApiKey = 35
	SaslAuthenticate             ApiKey = 36
	CreatePartitions             ApiKey = 37
	CreateDelegationToken        ApiKey = 38
	RenewDelegationToken         ApiKey = 39
	ExpireDelegationToken        ApiKey = 40
	DescribeDelegationToken      ApiKey = 41
	DeleteGroups                 ApiKey = 42
	ElectLeaders                 ApiKey = 43
	IncrementalAlterConfigs      ApiKey = 44
	AlterPartitionReassignments  ApiKey = 45
	ListPartitionReassignments   ApiKey = 46
	OffsetDelete                 ApiKey = 47
	DescribeClientQuotas         ApiKey = 48
	AlterClientQuotas            ApiKey = 49
	DescribeUserScramCredentials ApiKey = 50
	AlterUserScramCredentials    ApiKey = 51
	DescribeQuorum               ApiKey = 55
	AlterPartition               ApiKey = 56
	UpdateFeatures               ApiKey = 57
	Envelope                     ApiKey = 58
	DescribeCluster              ApiKey = 60
	DescribeProducers            ApiKey = 61
	UnregisterBroker             ApiKey = 64
	DescribeTransactions         ApiKey = 65
	ListTransactions             ApiKey = 66
	AllocateProduceIds           ApiKey = 67
	ConsumerGroupHeartbeat       ApiKey = 68
)

var apiNames = map[ApiKey]string{
	Produce:                      "Produce",
	Fetch:                        "Fetch",
	ListOffsets:                  "ListOffsets",
	Metadata:                     "Metadata",
	LeaderAndIsr:                 "LeaderAndIsr",
	StopReplica:                  "StopReplica",
	UpdateMetadata:               "UpdateMetadata",
	ControlledShutdown:           "ControlledShutdown",
	OffsetCommit:                 "OffsetCommit",
	OffsetFetch:                  "OffsetFetch",
	FindCoordinator:              "FindCoordinator",
	JoinGroup:                    "JoinGroup",
	Heartbeat:                    "Heartbeat",
	LeaveGroup:                   "LeaveGroup",
	SyncGroup:                    "SyncGroup",
	DescribeGroups:               "DescribeGroups",
	ListGroups:                   "ListGroups",
	SaslHandshake:                "SaslHandshake",
	ApiVersions:                  "ApiVersions",
	CreateTopics:                 "CreateTopics",
	DeleteTopics:                 "DeleteTopics",
	DeleteRecords:                "DeleteRecords",
	InitProducerId:               "InitProducerId",
	OffsetForLeaderEpoch:         "OffsetForLeaderEpoch",
	AddPartitionsToTxn:           "AddPartitionsToTxn",
	AddOffsetsToTxn:              "AddOffsetsToTxn",
	EndTxn:                       "EndTxn",
	WriteTxnMarkers:              "WriteTxnMarkers",
	TxnOffsetCommit:              "TxnOffsetCommit",
	DescribeAcls:                 "DescribeAcls",
	CreateAcls:                   "CreateAcls",
	DeleteAcls:                   "DeleteAcls",
	DescribeConfigs:              "DescribeConfigs",
	AlterConfigs:                 "AlterConfigs",
	AlterReplicaLogDirs:          "AlterReplicaLogDirs",
	DescribeLogDirs:              "DescribeLogDirs",
	SaslAuthenticate:             "SaslAuthenticate",
	CreatePartitions:             "CreatePartitions",
	CreateDelegationToken:        "CreateDelegationToken",
	RenewDelegationToken:         "RenewDelegationToken",
	ExpireDelegationToken:        "ExpireDelegationToken",
	DescribeDelegationToken:      "DescribeDelegationToken",
	DeleteGroups:                 "DeleteGroups",
	ElectLeaders:                 "ElectLeaders",
	IncrementalAlterConfigs:      "IncrementalAlterConfigs",
	AlterPartitionReassignments:  "AlterPartitionReassignments",
	ListPartitionReassignments:   "ListPartitionReassignments",
	OffsetDelete:                 "OffsetDelete",
	DescribeClientQuotas:         "DescribeClientQuotas",
	AlterClientQuotas:            "AlterClientQuotas",
	DescribeUserScramCredentials: "DescribeUserScramCredentials",
	AlterUserScramCredentials:    "AlterUserScramCredentials",
	DescribeQuorum:               "DescribeQuorum",
	AlterPartition:               "AlterPartition",
	UpdateFeatures:               "UpdateFeatures",
	Envelope:                     "Envelope",
	DescribeCluster:              "DescribeCluster",
	DescribeProducers:            "DescribeProducers",
	UnregisterBroker:             "UnregisterBroker",
	DescribeTransactions:         "DescribeTransactions",
	ListTransactions:             "ListTransactions",
	AllocateProduceIds:           "AllocateProduceIds",
	ConsumerGroupHeartbeat:       "ConsumerGroupHeartbeat",
}
