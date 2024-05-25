#!/bin/bash
#
## mv -v [--verbose]: explain what is being done.
#
### I also use it with a combination of additional flags to know what they do.
### For instance:
#### mv -T -v SOURCE DEST
#### mv -t -v SOURCE DEST
#### mv --no-copy -v SOURCE DEST
#
# Full documentation <https://www.gnu.org/software/coreutils/mv> or
# available locally via: info '(coreutils) mv invocation'
#
mv -v TEST_DIR_1 TEST_DIR_2
