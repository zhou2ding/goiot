package main

func main() {
	//sc := bufio.NewScanner(os.Stdin)
	//sc.Scan()
	//arr := strings.Split(sc.Text(), " ")
	//fmt.Println(getMaxCode(arr))
}

func getMaxCode(arr []string) string {
	if len(arr) == 1 {
		return arr[0]
	}

	res := make([]string, 0)
	// 	sorted := make([]string, len(arr))
	codeMap := make(map[string]bool)
	for _, s := range arr {
		codeMap[s] = true
	}
	// 	copy(sorted, arr)
	// 	sort.Strings(sorted)

	for i := range arr {
		tmp := arr[i]
		exist := false
		for j := len(tmp) - 1; j > 0; j-- {
			if codeMap[tmp[:j]] {
				exist = true
				break
			}
		}
		if exist {
			res = append(res, tmp)
		}
	}
	if len(res) == 0 {
		return ""
	}
	maxLenRes := getMaxLen(res)
	return maxLenRes[len(maxLenRes)-1]
}

func getMaxLen(arr []string) []string {
	maxLen := 0
	for _, s := range arr {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	res := make([]string, 0)
	for _, s := range arr {
		if len(s) == maxLen {
			res = append(res, s)
		}
	}
	return res
}
