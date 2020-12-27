#!/usr/bin/env bash

function print() {
    line=$(caller)
    echo $line ":" $@
}

go build
jun=$(realpath ./jun)
function jun(){
  $jun $@
}

testdir="git-test"
if [ -e "$testdir" ]; then
    rm -rf $testdir
fi

mkdir $testdir
cd $testdir

print "test init"
jun init junDir
git init gitDir
echo -e

print "test hash-object"
echo "version1" > junDir/file.txt
echo "version1" > gitDir/file.txt
jun hash-object -w junDir/file.txt > hash1
git hash-object -w gitDir/file.txt > hash2
cmp hash1 hash2
echo -e

print "test cat-file"
h1=$(<hash1)
h2=$(<hash2)
jun cat-file -p $h1 > fileContent1
git cat-file -p $h2 > fileContent2
cmp fileContent1 fileContent2

jun cat-file -t $h1 > fileType1
git cat-file -t $h2 > fileType2
cmp fileType1 fileType2
jun cat-file -s $h1 > fileSize1
git cat-file -s $h2 > fileSize2
cmp fileSize1 fileSize2
echo -e

print "test update-index --add, git ls-files --stage"
cd junDir
jun update-index --add file.txt
jun ls-files --stage > ../stage1
cd ../gitDir
git update-index --add file.txt
git ls-files --stage > ../stage2
cd ..
cmp stage1 stage2
echo -e

print "test write-tree"
cd junDir
mkdir dir
echo 123 > dir/file1
echo abc > dir/file2
jun update-index --add dir/file1
jun update-index --add dir/file2
jun write-tree > treeobj1sha1
#mkdir ../../test-data
#jun cat-file -p $(<treeobj1) > ../../test-data/test-write-tree-data.txt
jun cat-file -p $(<treeobj1sha1) > treeobjContent
cmp treeobjContent ../../test-data/test-write-tree-data.txt
echo -e

print "test commit-tree, log"
jun commit-tree $(<treeobj1sha1) -m "first commit" > commit1Sha1
echo version2 > file.txt
jun update-index --add file.txt
jun write-tree > treeobj2sha1
jun commit-tree $(<treeobj2sha1) -p $(<commit1Sha1) -m "second commit" > commit2Sha1
jun log $(<commit2Sha1) > log
echo -e show log :
cat log
echo -e

print "test update-ref"
jun update-ref refs/heads/master $(<commit2Sha1)
jun log master

print "test symbolic-ref"
jun update-ref refs/heads/test $(<commit1Sha1)
jun symbolic-ref HEAD refs/heads/test
jun log

print "test commit"
jun commit -m 'test commit'
jun log