package gcache

import (
	"bufio"
	"net"
)

/**
 * 包装conn
 */
type gconnnect struct {
	Conn net.Conn
}

/**
 * 读缓冲器
 */
func (g *gconnnect) read() ([]byte, error) {
	reader := bufio.NewReader(g.Conn)
	readBytes, err := reader.ReadBytes(IO_BUFF_END)
	if err != nil {
		return []byte(""), err
	}
	return readBytes[:len(readBytes)-1], nil
}

/**
 * 写缓冲器
 */
func (g *gconnnect) write(content []byte) {
	writer := bufio.NewWriter(g.Conn)
	writer.Write(content)
	writer.WriteByte(IO_BUFF_END)
	writer.Flush()
}

/**
 * 关闭
 */
func (g *gconnnect) close() {
	g.Conn.Close()
}
