package innchild

import "io"

// childReader читаем приглашение потомка
func childReader(r io.Reader) (buffer []byte, len int, err error) {

	temp := make([]byte, 20)
	for {
		n, err := r.Read(temp[0:])
		buffer = append(buffer, temp[0:n]...)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		len = n + len

		if n == 0 || (n > 0 && temp[n-1] == '\n') {
			break
		}
	}
	return buffer, len, err
}
