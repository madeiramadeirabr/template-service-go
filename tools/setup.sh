#!/bin/sh
YELLOW=$'\e[0;33m'
GREEN=$'\e[0;32m'
RED=$'\e[0;31m'
NC=$'\e[0m'
ERROR=0
DEFAULT_PROJECT_NAME='template-service-go'

echo "${GREEN} Setting up project for local development...${NC}"

echo "  - Installing Git hooks"
{
  ln -s ./tools/hooks/* .git/hooks && echo "    ${GREEN}DONE ${NC}"
} || {
 echo "    ${RED}ERROR ${NC}" && ERROR=1
}

echo "  - Creating .env file"
{
  cp .env.example .env && echo "    ${GREEN}DONE ${NC}"
} || {
 echo "    ${RED}ERROR ${NC}" && ERROR=1
}

{
  while true; do
    read -e -p "  - Rename project imports [Y/n]? " replaceImports
    replaceImports=${replaceImports:-Y}

    case ${replaceImports} in
      [Yy]* )
        read -e -p "    - Project name: " projectName

        if [[ ! -z "$projectName" ]]; then
          grep -RiIl "${DEFAULT_PROJECT_NAME}" --exclude-dir={.git,tools} | xargs -I@ sed -i '' "s/${DEFAULT_PROJECT_NAME}/${projectName}/g" @
          echo "    - Replacing with:${YELLOW} ${projectName} ${NC}"
          echo "    ${GREEN}DONE ${NC}"
        else
          echo "    - Invalid project name, skipping rename"
        fi
        break
      ;;
      [Nn]* )
        echo "    - Decided not rename"
        break
      ;;
      * )
        echo "    - Please use y or n"
        continue
      ;;
    esac
  done
} || {
 echo "    ${RED}ERROR ${NC}" && ERROR=1
}

if [ $ERROR -eq 1 ]; then echo "${RED} Setup failed ${NC}"; fi
if [ $ERROR -eq 0 ]; then echo "${GREEN} Setup successful 🧜🏻‍${NC}"; fi

