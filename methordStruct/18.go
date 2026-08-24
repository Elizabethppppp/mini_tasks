package methordStruct

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }
