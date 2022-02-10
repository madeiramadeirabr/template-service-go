#!/bin/sh
GREEN=$'\e[0;32m'
RED=$'\e[0;31m'
NC=$'\e[0m'
ERROR=0

echo "${GREEN} Setting up project for local development...${NC}"

echo "  - Installing Git hooks"
{
  ln -s ./tools/hooks/* .git/hooks && echo "${GREEN}  DONE ${NC}"
} || {
 echo "${RED} ERROR ${NC}" && ERROR=1
}

echo "  - Creating .env file"
{
  cp .env.example .env && echo "${GREEN}  DONE ${NC}"
} || {
 echo "${RED} ERROR ${NC}" && ERROR=1
}

if [ $ERROR -eq 1 ]; then echo "${RED} Setup failed ${NC}"; fi
if [ $ERROR -eq 0 ]; then echo "${GREEN} Setup successful 🧜🏻‍${NC}"; fi

