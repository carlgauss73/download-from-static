package main


func parse(h string) ([500]string, int){
	sz := len(h)
	var ans [500]string
	ind := 0
	var cmp string = "<a href="
	for i := 0; i<sz-8; i++ {
		if string(h[i:i+8]) == cmp {
			for ;int(h[i]) != 34; i++ {}
			aux := ""
			i++ 
			for ;int(h[i]) != 34; i++ {
				if h[i] == '%' {
					aux = aux + " ";
					i = i + 2
				} else {
					aux = aux + string(h[i])
				}
			}
			ans[ind] = aux
			ind++
		}
	}
	return ans, ind;
}

func dfs(curd , root string) ([500]string, int, error){
	var ans [500]string
	nz := 0
	h, err := html_dump(curd)
	if err != nil {
		return ans, nz, err
	}
	tab, sz := parse(h)
	for i := 0; i<sz-1; i++ {
		if tab[i][len(tab[i])-3:len(tab[i])] != "../" {
			temp := tab[i];
			if temp[len(temp)-1]!='/' {
				ans[nz] = root + temp[1:len(temp)]
				nz++
			} else {
				tmp2, sz2, err := dfs(root + temp[1:len(temp)], root)
				if err != nil {
					return ans, nz, err
				}
				for j := 0; j < sz2; j++ {
					ans[nz] = root + tmp2[j][1:len(tmp2[j])]
					nz++
				}
			}
		}
	}
	return ans, nz, nil
}
