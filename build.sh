#!/bin/bash
IMAGE_NAME="atena"
clear

hello(){
  # echo "*** $IMAGE_NAME"
  print "$IMAGE_NAME"
  print "Hello"
}

help(){
    RED=`tput setaf 1`
    GREEN=`tput setaf 2`
    YELLOW=`tput setaf 3`
    WHITE=`tput setaf 7`
    print "Help"
    echo "${GREEN}build: ${YELLOW}Build the Dockerfile to a docker image"
    echo "${GREEN}run: ${YELLOW}Run the docker image and enter on it. Build image if does not exist"
    echo "${GREEN}clean: ${YELLOW}Delete the image"
    echo "${GREEN}destroy: ${YELLOW}Delete all image and containers"
}

print(){
  #Format filling with empty spaces
    OUTPUT=$(printf "%-50s" $1"&")


  #Replace empty spaces
  OUTPUT="${OUTPUT// /*}"

  #Workaroung to keep one empty space very after text
  OUTPUT="${OUTPUT//&/ }"
  echo "*** $OUTPUT"
}

build(){
  echo "*** Building image $IMAGE_NAME"
  docker build -t $IMAGE_NAME .
}

build_if_not_exists(){
  if [ "$(docker images -q $IMAGE_NAME 2> /dev/null)" = "" ]; then
    echo "*** Image does not exists"
    build
  fi
}

run(){
  build_if_not_exists
  echo "*** Running"
  docker run -ti -v $(pwd):/go/src -e USER=$USER -e USERID=$UID $IMAGE_NAME bash
}

clean(){
  echo "*** Removing image $IMAGE_NAME"
  docker rmi -f $IMAGE_NAME
}

destroy(){
  echo "*** This will remove all containers"
  ask_for_confirmation
  OPTION=$?
  if [ "$OPTION" = 1 ]; then
    docker rm -f $(docker ps -aq)
    docker rmi -f $(docker images -aq)
  fi
}

ask_for_confirmation(){
  read -p "Continue (y/n)?" CHOICE
  case "$CHOICE" in
    y|Y ) OPTION=1;;
    n|N ) OPTION=0;;
    * ) echo "Invalid option!";;
  esac

  return $OPTION
}

# Execute the tasks
for TASK in "$@"
do
 $TASK
done

# If nothing, build default
if [ "$1" = "" ]; then
  build
fi
