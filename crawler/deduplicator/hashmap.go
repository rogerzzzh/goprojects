package deduplicator

type MapDeduplicator struct {
	data map[interface{}]bool
}

func (d *MapDeduplicator) IsExist(key interface{}) bool {
	return d.data[key]
}

func (d *MapDeduplicator) Add(key interface{}) {
	if d.data == nil {
		d.data = make(map[interface{}]bool)
	}
	d.data[key] = true
}
