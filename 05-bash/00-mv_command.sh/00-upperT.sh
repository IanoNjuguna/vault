#!/bin/bash

#-----------------------------------------------------------------------------------------
#                                                                                         |
# mv -T SOURCE/ DEST/                                                                     |
#                                                                                         |
# -> mv -T [--no-target-directory]: treat DEST as a normal file.                          |
#                                                                                         |
# -> Overwrite 'SOURCE/' onto an empty directory 'DEST/'.                                 |
#                                                                                         |
# --> What this means is that the execution of the mv command with this flag,             |
# --> will change the name 'SOURCE/' to 'DEST/' and                                       |
# --> move the contents of the 'SOURCE/' directory to the empty 'DEST/' directory         |
#                                                                                         |
# ----> rename 'SOURCE/' -> 'DEST'                                                        |
#                                                                                         |
# --> SOURCE and DEST should be directories.                                              |
# ---> If not, console prints:                                                            |
# ----> mv: cannot overwrite non-directory 'SOURCE' with directory 'DEST/' or,            |
# ----> mv: cannot overwrite directory 'SOURCE/' with non-directory.                      |
#                                                                                         |
# --> The 'SOURCE/' directory can either empty or not empty.                              |
# ---> After all, our need is to move it's contents to the 'DEST/' directory.             |
# ---> What difference will it make if either is the case?                                |
#                                                                                         |
# --> However, the 'DEST/' directory should be an empty directory.                        |
# ---> [--no-target-directory]                                                            |
# ---> If not, console prints:                                                            |
# ----> mv: cannot overwrite 'DEST': Directory not empty.                                 |
#                                                                                         |
# -> I would use the flag:                                                                |
# --> in instances where I need to ensure or check that 'DEST/' is an empty directory.    |
#                                                                                         |
# -> Full documentation <https://www.gnu.org/software/coreutils/mv> or,                   |
# --> available locally via: info '(coreutils) mv invocation'                             |
#                                                                                         |
#-----------------------------------------------------------------------------------------
mv -v -T SOURCE DEST
