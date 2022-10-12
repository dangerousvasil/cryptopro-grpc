package rwCloser

import "io"

// BufferReader читаем соединения до конца пакета
func BufferReader(r io.Reader) (buffer []byte, len int, err error) {

	buf := make([]byte, DEFAULT_BUFFER_LENGTH)
	for {
		n, err := r.Read(buf[0:])

		buffer = append(buffer, buf[0:n]...)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		len = n + len

		// прервать чтение из соединения, если:
		// - прочитано 0 байтов, ИЛИ
		// - прочитано больше 0 байтов, но последний байт - символ перехода на новую строку
		if n == 0 || (n > 0 && buf[n-1] == '\n') {
			break
		}
	}
	return buffer, len, err
}
