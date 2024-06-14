This is just a learning repo. I will add some commands that can be used

vc init -> initialize a new repo 
vc hash-object [file name] -> creates a hash object of a given file
vc cat-file [oid/tag name] [expected item type] -> gets a hash file and dehashifyis it
vc write-tree -> creates a new tree representation hashed in the .vc/object directory
vc read-tree [oid/tag name] -> reads the tree into file system and overwrites everything    to the vc state. This is used by the checkout command
vc commit -m [commit message] -> commits changes to the working tree
vc log [oid/tag name] -> shows the commits under the oid or tag. default HEAD.
vc checkout [oid] -> checks out the commit by oid or tag name
vc tag -oid [oid] -name [tag name] -> creates a new tag at oid or default HEAD

