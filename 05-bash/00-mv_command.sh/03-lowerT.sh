#!/bin/bash

#----------------------------------------------------------------------------
#                                                                            |
# mv -t [--target-directory=DIRECTORY]: move all SOURCE(s) into DIRECTORY.   |
#                                                                            |
# -> mv -v -t SOURCE... DIRECTORY/                                           |
# --> rename 'SOURCE(s)' -> 'DIRECTORY/SOURCE(s)'                            |
# --> 'SOURCE(s)' can be a multiple number of directories/files.             |
#                                                                            |
# -> If 'DIRECTORY/' is not a directory, console prints:                     |
# --> mv: target directory 'DIRECTORY': Not a directory                      |
#                                                                            |
# Full documentation <https://www.gnu.org/software/coreutils/mv> or          |
# available locally via: info '(coreutils) mv invocation'                    |
#-----------------------------------------------------------------------------

mv -v -t SOURCE DIRECTORY