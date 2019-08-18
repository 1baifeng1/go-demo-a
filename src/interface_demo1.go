package main

type BaseQuestion struct {
	questionId int
	questionContent string
}

type ChoiceQuestion struct {
	baseQuestion BaseQuestion
	options []string
}

type BlankQuestion struct {
	baseQuestion BaseQuestion
	blank string
}

func fetchFromChoiceTable(id int) (data *ChoiceQuestion) {
	if 1001 == id {
		return &ChoiceQuestion{
			baseQuestion: BaseQuestion{
				questionId: 1001,
				questionContent: "hahhha",
			},
			options: []string{"A", "B"},
		}
	}
	return nil
}


func main() {

}
