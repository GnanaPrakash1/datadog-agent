// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build !docker

package common

// isDockerRunning check if docker is running
func isDockerRunning() bool {
	return false
}
