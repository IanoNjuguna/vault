#!/bin/bash
#
#----------------------------------------------------------------------
# mv SOURCE(s) DEST                                                    |
#                                                                      |
# -> Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.            |
#                                                                      |
# Usage:                                                               |
#                                                                      |
# -> mv SOURCE... DIRECTORY                                            |
# --> move SOURCE(s) to DIRECTORY                                      |
# ---> SOURCE... entails the SOURCE/ can be multiple directories/files.|
# ---> rename 'SOURCE...' -> 'DIRECTORY/SOURCE...'                     |
# ---> DIRECTORY can be empty or non-empty.                            |
#                                                                      |
# -> mv [-T] SOURCE DEST                                               |
# --> Documentation: `cat 00-upperT.sh`                                |
#                                                                      |
# -> mv -t DIRECTORY SOURCE...                                         |
# --> Documentation: `cat 03-lowerT.sh`                                |
#                                                                      |
# Full documentation <https://www.gnu.org/software/coreutils/mv> or    |
# available locally via: info '(coreutils) mv invocation'              |
#----------------------------------------------------------------------
mv SOURCE DEST
