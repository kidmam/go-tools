// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net2

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// NormalizeMac normalizes the mac.
//
// If fill is true, pad with leading zero, such as 01:02:03:04:05:06.
//
// If upper is true, output the upper character, such as AA:BB:11:22:33:44.
// Or, output the lower character, such as aa:bb:cc:11:22:33.
//
// Return "" if the mac is a invalid mac.
func NormalizeMac(mac string, fill, upper bool) string {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return ""
	}

	width := ""
	_upper := "x"
	if upper {
		_upper = "X"
	}
	if fill {
		width = "2"
	}
	formatter := fmt.Sprintf("%%0%s%s", width, _upper)

	for i, m := range macs {
		v, err := strconv.ParseUint(m, 16, 8)
		if err != nil {
			return ""
		}
		macs[i] = fmt.Sprintf(formatter, v)
	}

	return strings.Join(macs, ":")
}

// NormalizeMacFU is equal to NormalizeMac(mac, true, true).
func NormalizeMacFU(mac string) string {
	return NormalizeMac(mac, true, true)
}

// NormalizeMacFu is equal to NormalizeMac(mac, true, false).
func NormalizeMacFu(mac string) string {
	return NormalizeMac(mac, true, false)
}

// NormalizeMacfU is equal to NormalizeMac(mac, false, true).
func NormalizeMacfU(mac string) string {
	return NormalizeMac(mac, false, true)
}

// NormalizeMacfu is equal to NormalizeMac(mac, false, false).
func NormalizeMacfu(mac string) string {
	return NormalizeMac(mac, false, false)
}

// GetMacByInterface returns the MAC of the interface iface.
func GetMacByInterface(iface string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		if i.Name == iface {
			return i.HardwareAddr.String(), nil
		}
	}

	return "", fmt.Errorf("no interface '%s'", iface)
}

// GetMacByIP returns the MAC of the interface to which the ip is bound.
func GetMacByIP(ip string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	ip = strings.ToLower(ip)
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			_ip := strings.Split(addr.String(), "/")[0]
			if _ip == ip {
				return iface.HardwareAddr.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no mac about '%s'", ip)
}
