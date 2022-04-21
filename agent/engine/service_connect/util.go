// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package serviceconnect

import (
	"fmt"
	"math/rand"
	"time"

	apitask "github.com/aws/amazon-ecs-agent/agent/api/task"
)

const (
	// From https://www.kernel.org/doc/html/latest//networking/ip-sysctl.html#ip-variables
	ephemeralPortMin         = 32768
	ephemeralPortMax         = 60999
	maxPortSelectionAttempts = 100
)

// DNSConfigToDockerExtraHostsFormat converts a DNSConfig to a list of ExtraHost entries that Docker will
// recognize.
func DNSConfigToDockerExtraHostsFormat(dnsConf apitask.DNSConfig) []string {
	var hosts []string
	for _, vipConf := range dnsConf {
		if len(vipConf.IPV4Address) > 0 {
			hosts = append(hosts,
				fmt.Sprintf("%s:%s", vipConf.HostName, vipConf.IPV4Address))
		}
		if len(vipConf.IPV6Address) > 0 {
			hosts = append(hosts,
				fmt.Sprintf("%s:%s", vipConf.HostName, vipConf.IPV6Address))
		}
	}
	return hosts
}

// GenerateEphemeralPortNumbers generates a list of n unique port numbers in the 32768-60999 range. The resulting port
// number list is guaranteed to not include any port number present in "exclude" parameter.
func GenerateEphemeralPortNumbers(n int, exclude []int) ([]int, error) {
	toExcludeSet := map[int]struct{}{}
	for _, e := range exclude {
		toExcludeSet[e] = struct{}{}
	}
	rand.Seed(time.Now().UnixNano())

	var result []int
	var portSelectionAttempts int
	for len(result) < n {
		// The intention of maxPortSelectionAttempts is to avoid a super highly unlikely case where we
		// keep getting ports that collide, thus creating an infinite loop.
		if portSelectionAttempts > maxPortSelectionAttempts {
			return nil, fmt.Errorf("maximum number of attempts to generate unique ports reached")
		}
		port := rand.Intn(ephemeralPortMax-ephemeralPortMin+1) + ephemeralPortMin
		if _, ok := toExcludeSet[port]; ok {
			portSelectionAttempts++
			continue
		}
		toExcludeSet[port] = struct{}{}
		result = append(result, port)
	}
	return result, nil
}
