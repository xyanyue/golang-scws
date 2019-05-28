package scws

/*
#cgo CFLAGS: -I/usr/local/scws/include/scws -fno-stack-protector
#cgo LDFLAGS: -L/usr/local/scws/lib/ -lscws

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include </usr/local/scws/include/scws/scws.h>
#define TAG_NUM 50

typedef struct {
    char r[500];
    int  n;
}RESULT[TAG_NUM];

char *scws_get_top(char *text){

    RESULT *R = (RESULT *)malloc(sizeof(RESULT)*TAG_NUM);
    scws_t s;
    scws_top_t res, cur;

    if (!(s = scws_new())) {
      printf("ERROR: cann't init the scws!\n");
    }
    scws_set_charset(s, "utf8");
    //scws_set_dict(s, "/usr/local/scws/etc/dict.utf8.xdb", SCWS_XDICT_XDB);
    scws_set_dict(s, "/home/miliguy/golang/user_profile/src/dict.txt", SCWS_XDICT_TXT);
    // scws_set_rule(s, "/usr/local/scws/etc/rules.utf8.ini");

    scws_send_text(s, text, strlen(text));
    // cur = scws_get_result(s);
    int i = 0;
    int j;
    int l = 0;
    //printf("%s",text);
    res = cur = scws_get_tops(s,TAG_NUM,"@");
    while (cur != NULL)
    {
        char * word = cur->word;
        strcpy(R[i]->r,word);
        R[i]->n = cur->times;
        cur = cur->next;
        l = l + (strlen(word))*sizeof(char) + 3*5*sizeof(char);
        //printf("WORLD:%s:%d",R[i]->r,l);
        i++;
    }

    scws_free_tops(res);
    scws_free(s);
    char *result = malloc(l+7);
    strcpy(result,"{");
    for(j=0;j<i;j++){
        //sprintf(result,"'%s':%d",R[j]->r,R[j]->n);
        strcat(result,"\"");
        strcat(result,R[j]->r);
        strcat(result,"\"");
        strcat(result,":");
        char numstr[10];
        sprintf(numstr,"%d",R[j]->n);
        strcat(result,numstr);
        if(j< i-1){
            strcat(result,",");
        }
    }
    strcat(result,"}");
    free(R);
    R = NULL;
    return result;
}
*/
import "C"

import (
	"encoding/json"
	// "fmt"
	"unsafe"
)

func RunTop(text string) map[string]int {
	// var pDetectInfo unsafe.Pointer
	// fmt.Println("--------1------", text)
	str := C.CString(text)
	// fmt.Println("--------2------", str)
	var y map[string]int
	a := C.scws_get_top(str)
	// fmt.Println("-------3-------", C.GoString(a))
	json.Unmarshal([]byte(C.GoString(a)), &y)
	// fmt.Println(y)

	C.free(unsafe.Pointer(str))

	return y
}
