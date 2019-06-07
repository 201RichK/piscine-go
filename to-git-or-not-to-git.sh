#!/bin/bash

curl https://api.github.com/users/201RichK | grep '"id":' | cut -d ':' -f2 | cut -d "," -f1


