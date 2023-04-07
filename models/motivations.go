package models

type motivation struct {
	body string
	policies []string
}

type Motivations struct {
	motivations map[int]motivation
}

func (motivations *Motivations) loadMotivations(config_location string) error {
	file, _ := ioutil.ReadFile(config_location)
	
	var data []string
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		return err
	}

	fmt.Println("Loaded Motivations")
	for i := 0; i < len(data); i++ {
		fmt.Printf("Motivation %s: %s \n",strconv.Itoa(i), data[i])
	}

	return nil
}