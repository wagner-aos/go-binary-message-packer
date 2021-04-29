package utils

/*
var isoJSON parse.IsoJSON

//CreateISOmessageReadMeFile - It creates a READ for ISO Message
func CreateISOmessageReadMeFile(jsonString string) {

	error := json.Unmarshal([]byte(jsonString), &isoJSON)
	if error != nil {
		print(error)
	}

	file, err := os.Create("ISO-README.md")
	if err != nil {
		print("Error creating file...")
	}

	for _, v := range isoJSON.Elements {

		line := []string{"| "}

		strings.Join(line, string(v.Field))
		strings.Join(line, string(" | "))
		strings.Join(line, string(v.Label))
		strings.Join(line, string(" | "))
		strings.Join(line, string(v.Value))
		strings.Join(line, string(" | "))
		strings.Join(line, string(v.ContentType))
		strings.Join(line, string(" |"))

		file.WriteString(strings.Join(line, ""))

	}

}

*/
