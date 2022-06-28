#!/usr/bin/python3
import sys
import yaml

print("JEG Open source code")

argumentlist = sys.argv
arg0, arg1 = argumentlist
# print ('Number of arguments:', len(sys.argv), 'arguments.')
# print ('Argument List:', str(argumentlist))

#
# =================================
# Open the file describe in arg1

fo = open(arg1, "r")

fname = fo.name
ID = fname[:8]
TRANSFERDATE = fname[9:17]

infileseq = int(fo.read(6))
infiledate = fo.read(8)

with open('pilote01.yaml', 'r') as file:
    yamlfile = yaml.safe_load(file)

yamlfileseq = yamlfile['lastSequence']

status ={}
status = dict([
    ('File',str(arg1)),
    ('ID',ID),
    ('transferdate',TRANSFERDATE),
    ('filesequence',infileseq),
    ('csvsequence',yamlfileseq),
    ('filedate',infiledate)
    ])

print (status)


# # # DEBUG # # # 
# print ("yamlfile :", yamlfile)
# print (yamlfile['prefix'][ID])
# with open('pilote03.yaml', 'w') as file:
#     yaml.dump(status, file)
# # # DEBUG # # # 

print()

if infileseq == yamlfileseq + 1 :
    print ("filesequence is + 1")
    yamlfile['lastSequence'] = infileseq
    with open('pilote01.yaml', 'w') as file:
        yaml.dump(yamlfile, file)
    print('check directory target')
    print('if busy, wait 10s')
    print('if no same filetype, deplace file on it')

# YAML
lastSequence: 56648
prefix:
  bbb: wrkBBB
  pilote01: wrkpilote01

