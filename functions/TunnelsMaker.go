package functions

import "slices"

func TunnelsMaker(Tunnels map[string][]string, a string, b string) {
	check_1 := slices.Contains(Tunnels[a], b)
	if !check_1 {
		Tunnels[a] = append(Tunnels[a], b)
	}
}
