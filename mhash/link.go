package mhash

type node struct {
	key   string
	value int
	next  *node
}

type link struct {
	head *node
}

// 设置结点key
// key 存在 覆盖
// 不存在 添加
func (l *link) set(key string, value int) {
	if l == nil {
		panic("not exit link")
		return
	}
	curr := l.head
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}

	nd := new(node)
	nd.key = key
	nd.value = value
	nd.next = l.head
	l.head = nd
}

// 查看key 信息
func (l *link) get(key string) (int, bool) {
	if l == nil || l.head == nil {
		return 0, false
	}

	curr := l.head
	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
		curr = curr.next
	}
	return 0, false
}

// 删除对应结点
// 不存在key 忽略
func (l *link) del(key string) {
	if l == nil || l.head == nil {
		return
	}
	if l.head.key == key {
		l.head = l.head.next
		return
	}
	last := l.head
	for last.next != nil {
		if last.next.key == key {
			last.next = last.next.next
			return
		}
		last =last.next
	}
	return
}

// 链表长度
func (l *link)len() int  {
	if l == nil || l.head == nil {
		return 0
	}
	len := 0
	curr := l.head
	for curr != nil {
		len ++
		curr = curr.next
	}
	return  len
}