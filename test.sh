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

testdir="../git-test"
if [ -e "$testdir" ]; then
    rm -rf $testdir
fi

mkdir $testdir
cd $testdir

print "test git init"
jun init junDir
git init gitDir
echo -e

print "test git hash-object"
echo "version1" > junDir/file.txt
echo "version1" > gitDir/file.txt
jun hash-object -w junDir/file.txt > hash1
git hash-object -w gitDir/file.txt > hash2
cmp hash1 hash2
echo -e

print "test git cat-file"
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