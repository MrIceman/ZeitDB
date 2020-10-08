package data

type Engine interface {
	Update(model *ZeitModel)
	Insert(model *ZeitModel)
	Delete(model *ZeitModel)
	GetAll() []ZeitModel
	GetLatest(n int) []ZeitModel
	GetByRange(from int64, to int64) []ZeitModel
}

type ZeitEngine struct {
	
}

func (z ZeitEngine) Update(model *ZeitModel) {
	panic("implement me")
}

func (z ZeitEngine) Insert(model *ZeitModel) {
	panic("implement me")
}

func (z ZeitEngine) Delete(model *ZeitModel) {
	panic("implement me")
}

func (z ZeitEngine) GetAll() []ZeitModel {
	panic("implement me")
}

func (z ZeitEngine) GetLatest(n int) []ZeitModel {
	panic("implement me")
}

func (z ZeitEngine) GetByRange(from int64, to int64) []ZeitModel {
	panic("implement me")
}
