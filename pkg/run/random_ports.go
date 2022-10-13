package run

import "net"

// GetRandomPorts returns a slice of random ports.
func GetRandomPorts(count int) (ports []int, retErr error) {
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}

		defer func(l *net.TCPListener) {
			err := l.Close()
			if err != nil {
				retErr = err
			}
		}(l)

		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}

	return ports, nil
}
