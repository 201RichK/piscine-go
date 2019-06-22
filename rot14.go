package piscine

func Rot14(str string) string {

	newStr := []byte(str)

	for i := 0; i < len(str); i++ {
		if (newStr[i] >= 'A' && newStr[i] <= 'L') || (newStr[i] >= 'a' && newStr[i] <= 'l') {
			newStr[i] += 14
		} else if (str[i] >= 'M' && str[i] <= 'Z') || (str[i] >= 'o' && str[i] <= 'z') {
			newStr[i] -= 12
		}
	}
	return string(newStr)
}


package piscine

func Rot14(str string) string {
	bstr:=[]byte(str)
	for i:= range bstr{
		if bstr[i]>=65 && bstr[i]<=90{
			if (bstr[i]+14)>90{
				bstr[i]=bstr[i]-12			
			}else{
				bstr[i]+=14
			}
		}
		if bstr[i]>=97 && bstr[i]<=122{
			if (bstr[i]+14)>122{
				bstr[i]=bstr[i]-12			
			}else{
				bstr[i]+=14
			}
		}
	}
	return string(bstr)
}
