package entity

type Verify struct {
	Key    string `json:"key"`
	Data   string `json:"data"`
	Value1 string `json:"value_1"`
}

func (e *Verify) GetKey() string {
	return e.Key
}

func (e *Verify) GetData() string {
	return e.Data
}

func (e *Verify) GetValue1() string {
	return e.Value1
}

func (e *Verify) GetValue2() string {
	return e.Value2
}

func (e *Verify) SetKey(data string) {
	e.Key = data
}

func (e *Verify) SetData(data string) {
	e.Data = data
}
