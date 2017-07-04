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
func (g *gconnnect) read() (string, error) {
	reader := bufio.NewReader(g.Conn)
	readBytes, err := reader.ReadBytes(DELIMITER)
	if err != nil {
		return "", err
	}
	return string(readBytes[:len(readBytes)-1]), nil
}

/**
 * 写缓冲器
 */
func (g *gconnnect) write(content string) {
	writer := bufio.NewWriter(g.Conn)
	writer.WriteString(content)
	writer.Flush()
}

/**
 * 关闭
 */
func (g *gconnnect) close() {
	g.Conn.Close()
}
