package cmd

import (
	"github.com/fatih/color"
)

func Greet() {
	nameASCIIFirst := `
 ####                                       ###        
 ######                                   #######      
 ########                       ######   #########     
 ##########                    ########  ###   ##      
 ####  #####                   ###  ###  ###           
 ####   #####                  ###       ###           
 ####    #####                 ###       ###           
 ####     #####                ###       ###           
 ####     #####                ###       ###           
 ####      #####               ###       ###          
 ####      #####               ###       ###       ### 
 ####      #####  ###   ####   ###       ###    ###### 
 ####      #####  ###   ####   ###       ### #######  
 ####      #####  ###   ####   ###       ########      
 ####       ####  ###   ####   ###      ######         
 ####       ####  ###   ####   ###   ########          
 ####       ####  #### #####   ##############          
 ####      #####  #########    #######   ####          
 ####      #####   ########  ######       ###          
 ####      ####      ###  ########        ###          
 ####     #####        ####### ###        ###          
 ####    #####       ######    ###   ###  ###          
 ####   #####      #####    #   ##  #### ####          
 #### ######      ###     ####  ##  ########           
 #########                ########   ######            
 #######                   #######                     
 #####                      ####                       
	`
	color.New(color.FgBlue).Print(nameASCIIFirst)

}