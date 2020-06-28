package concurrency

type opinstance struct {
	name string
}

func (this opinstance) Name() string {
	return this.name
}

func (this opinstance) MyPrint1() {

}

func (this opinstance) MyPrint2() {

}

func (this opinstance) MyPrint3() {

}

func NewOperator(name string) opinstance {
	return opinstance{
		name: name,
	}
}
