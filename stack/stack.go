package stack

var keyMap map[string]string

func keyMapFunc() {
	keyMap = make(map[string]string)
	keyMap["("] = ")"
	keyMap["["] = "]"
	keyMap["{"] = "}"
}
func isValid() {
	keyMapFunc()

}
