package data

type ZeitModel struct {
	id        int64
	Timestamp int64
	Value     interface{}
	Label     string
}

func (z *ZeitModel) Id() int64 {
	return z.Id()
}
