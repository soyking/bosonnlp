package bosonnlp

import (
	"fmt"
	"log"
)

func Example() {
	bosonClt := NewBosonNLPClient("your token")
	resp, err := bosonClt.ClassifyAnalysis([]string{"俄否决安理会谴责叙军战机空袭阿勒颇平民", "邓紫棋谈男友林宥嘉：我觉得我比他唱得好"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("第一条分类: %s\n", NewsType[resp[0]])
	fmt.Printf("第二条分类: %s\n", NewsType[resp[1]])

	//Output:
	//第一条分类: 军事
	//第二条分类: 娱乐
}
