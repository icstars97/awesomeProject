package main

import (
	"fmt"
	"bufio"

	"os"
	"strconv"
)
func pull() string {
	in := bufio.NewScanner(os.Stdin);
	in.Scan();
	return in.Text()
}


func push(stringval string)  (float64,string) {
	val,_:=strconv.ParseFloat(stringval,64);

	stringval="";
	return val,stringval;

}
func add(s string, stringval string,flag bool, inperr bool) (string,bool,bool) {
	_,err:=strconv.ParseInt(s,0,64);
	if err==nil||s=="." {
		stringval+=s;
		flag=true;
	} else{
		inperr=true;
	}
	return stringval,flag,inperr;

}

func main(){


	var input, stringval,s string;
	flag,inperr:=false,false;
	fmt.Println("Введите выражение(числа и знаки необходимо вводить через пробел)\n");
		input = pull();
		l := len(input);
			if input[l-1:]!="="{
				inperr=true;
			}
			sc,nc:=0,0;
			stack := make([]float64, l);
			p, i := 0, 0;
			space:=true;
			for i = 0; i < l; i++ {
				s = input[i:i+1];
					switch s {

					case " ":
						space=true;
						if flag {
							stack[p],stringval= push(stringval);
							nc++;
							p++;
							flag = false;
						}
					case "+":
						if sc<nc&&space {
							stack[p-2] = stack[p-2] + stack[p-1];
							p--;sc++;space=false;

						}else{
							inperr=true;
						}
					case "-":
						if sc<nc&&space {
							stack[p-2] = stack[p-2] - stack[p-1];
							p--;sc++;space=false;
						}else{
							inperr=true;
						}
					case "*":
						if sc<nc&&space {
							stack[p-2] = stack[p-2] * stack[p-1];
							p--;sc++;space=false;
						}else{
							inperr=true;
						}
					case "/":
						if sc<nc&&space {
							stack[p-2] = stack[p-2] / stack[p-1];
							p--;sc++;space=false;
						}else{
							inperr=true;
						}
					case "=":
						if sc<nc&&space&&i==l-1 {
							fmt.Println("результат:", stack[p-1]);
							//ls:=len(stack);
							//fmt.Println(stack[0:]);
						}else{
							inperr=true;
							}

					default:
						{
							stringval, flag, inperr = add(s, stringval, flag, inperr);
							space=false;
						};

					}
					if inperr {
						fmt.Println("\nошибка ввода");
						inperr=false;
						break;
					}

				}
			}




