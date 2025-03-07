#!/usr/bin/env sh

# FROM: https://stackoverflow.com/questions/60165440/how-do-i-refactor-module-name-in-go (user: Matthew Trent)

export CUR="minecart.compfest.id" # example: github.com/user/old-lame-name
export NEW="github.com/crlnravel/go-fiber-template" # example: github.com/user/new-super-cool-name
go mod edit -module ${NEW}
find . -type f -name '*.go' -exec perl -pi -e 's/$ENV{CUR}/$ENV{NEW}/g' {} \;
