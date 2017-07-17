package gcache

type gdb struct {
	size   int
	gdatas []*gdata
}

/*
 *增加/修改一个元素
 */
func (g *gdb) setNode(key string, valueType uint8, value interface{}) bool {
	hash := hashValue(&key, &db.size)
	node := db.gdatas[hash]

	if node == nil {
		newNode := new(gdata)
		newNode.setValue(key, valueType, value)
		db.gdatas[hash] = newNode
		return true
	}

	for {
		if node.key == key {
			node.setValue(key, valueType, value)
			return true
		}
		if node.next == nil {
			break
		}
		node = node.next
	}
	newNode := new(gdata)
	newNode.setValue(key, valueType, value)
	node.next = newNode
	return true
}

func (g *gdb) getNode(key string) (string, bool) {
	hash := hashValue(&key, &db.size)
	node := g.gdatas[hash]

	for {
		if node == nil {
			return "", true
		}
		if node.key == key {
			return interfaceToString(node.getValue()), true
		}
		node = node.next
	}
}

/*
 *删除一个元素
 */
func (g *gdb) deleteNode(key string) bool {
	hash := hashValue(&key, &db.size)
	node := g.gdatas[hash]

	if node == nil {
		return false
	}
	if node.key == key {
		g.gdatas[hash] = node.next
		return true
	}

	dNode := node.next
	for {
		if dNode == nil {
			return false
		}
		if dNode.key == key {
			node.next = dNode.next
			return true
		}
		node = node.next
		dNode = dNode.next
	}
}
