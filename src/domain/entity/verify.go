package entity

type Verify struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (e *Verify) GetKey() string {
	return e.Key
}

func (e *Verify) GetData() string {
	return e.Data
}

func (e *Verify) SetKey(data string) {
	e.Key = data
}

func (e *Verify) SetData(data string) {
	e.Data = data
}
