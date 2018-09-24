#!/bin/bash

export reset="\e[0m"                                                            
export blue="\e[1;34m"                                                          
export green="\e[1;42m"                                                           
export red="\e[1;31m"

#export ASCII_ART_HEADER="                                                              
#         __v_                                                                   
#        (.___\/{                                                                
#~^~^~^~^~^~^~^~^~^~^~^~^~"

export ASCII_ART_HEADER_BASE='                                                      
                    ##        .                                                 
              ## ## ##       ==                                                 
           ## ## ## ##      ===                                                 
       /""""""""""""""""\___/ ===                                               
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~                                        
       \______ o          __/                                                   
         \    \        __/                                                      
          \____\______/                                                         
                                                                                
          |          |                                                          
       __ |  __   __ | _  __   _                                                
      /  \| /  \ /   |/  / _\ |                                                 
      \__/| \__/ \__ |\_ \__  |                                            
'

export ASCII_ART_UP1=${ASCII_ART_HEADER_BASE}'      Bootsraping Docker Cluster...\n'
export ASCII_ART_UP2=${ASCII_ART_HEADER_BASE}'        Docker Cluster Ready...\n'
export ASCII_ART_DOWN=${ASCII_ART_HEADER_BASE}'     Tearing down Docker Cluster...\n'
export ASCII_ART_LOGIN=${ASCII_ART_HEADER_BASE}'   Loging in to the Docker Cluster...\n'
export ASCII_ART_EXEC=${ASCII_ART_HEADER_BASE}' Executing command on Docker Cluster...\n'
export ASCII_ART_SCALE=${ASCII_ART_HEADER_BASE}'     Scaling Docker Cluster...\n'
