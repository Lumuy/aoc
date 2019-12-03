package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput(input string) (res []int) {
	arr := strings.Split(input, "\n")

	for _, v := range arr {
		add, _ := strconv.Atoi(v)
		res = append(res, add)
	}

	return res
}

func countRequiredModule(mass int) int {
	return mass / 3 - 2
}

func countAddedRequiredModule(mass int) (res int) {
  for {
    mass = countRequiredModule(mass)
    if mass > 0 {
      res += mass
    } else {
      break
    }
  }

  return res
}

func main() {
  var res int
  arr := getInput(input)

  // Part 1
  for _, v := range arr {
    res += countRequiredModule(v)
  }
	fmt.Println(res)

  // Part 2
  res = 0
  for _, v := range arr {
    res += countAddedRequiredModule(v)
  }
	fmt.Println(res)
}

const input string = `120150
105328
70481
86674
112434
94883
147500
146007
103982
65758
132357
60885
97516
96977
129085
80330
124081
102501
102505
70029
54155
69253
60120
53192
89470
137125
136296
104546
92859
74937
135044
66238
126678
86364
138559
82393
96947
107749
115332
117563
95431
99640
107667
120427
108389
51567
57493
68518
114565
107248
50627
122517
129687
118989
52459
83726
106765
75872
147111
78822
65058
142460
122496
148942
72753
141599
105711
141860
134066
116716
107455
56673
148238
92318
86652
69312
91352
94528
73441
137814
80247
101115
61773
100951
77189
119083
93841
109090
83370
70230
144812
67647
105358
135780
85225
100697
100998
131151
52826
114084`
