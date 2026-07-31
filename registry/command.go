package registry 


type Command struct {
	Description		string
	Usage			string
	Run				func([]string)
}
