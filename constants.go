package main

import "time"

const (
	pollingInterval time.Duration = 10 * time.Second

	envVarRepoFullName     string = "GITHUB_REPOSITORY"
	envVarRunID            string = "GITHUB_RUN_ID"
	envVarRepoOwner        string = "GITHUB_REPOSITORY_OWNER"
	envVarToken            string = "INPUT_SECRET"
	envVarApprovers        string = "INPUT_APPROVERS"
	envVarApprovalWait     string = "INPUT_APPROVAL-WAIT"
	envVarMinimumApprovals string = "INPUT_MINIMUM-APPROVALS"
	envIssueTitle          string = "INPUT_ISSUE-TITLE"
	envIssueBody           string = "INPUT_ISSUE-BODY"
)

var (
	approvedWords = []string{"approved", "approve", "lgtm", "yes"}
	deniedWords   = []string{"denied", "deny", "no"}
)
